/*
此文件用于初始化redis
*/

package config

import (
	"context"
	"fmt"
	gocache "github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// 全局缓存实例
var (
	// RedisClient 重度读数据，跨进程共享
	RedisClient *redis.Client
	// MemoryCache 轻量实时数据，本进程内存缓存
	MemoryCache *gocache.Cache
)

// DefaultDuration redis数据保留时长（30天）
var DefaultDuration = 30 * 24 * 60 * 60 * time.Second

// ToInitRedis InitRedis 初始化Redis连接并返回错误
func ToInitRedis(cfg RedisConfig) (*redis.Client, error) {
	// 验证Redis密码
	if cfg.Password == "" {
		return nil, fmt.Errorf("redis密码未设置或使用默认值")
	}

	// 优化连接池配置
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Host + ":" + cfg.Port,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  5 * time.Second,  // 更短的连接超时
		ReadTimeout:  15 * time.Second, // 更短的读写超时
		WriteTimeout: 15 * time.Second,
		PoolSize:     50, // 更小的连接池
		MinIdleConns: 5,  // 更少的空闲连接
		MaxRetries:   3,
		PoolTimeout:  5 * time.Second, // 连接池获取超时
	})

	// 默认 cleanupInterval: 1s
	MemoryCache = gocache.New(5*time.Second, 1*time.Second)
	fmt.Println("MemoryCache已经初始化:", MemoryCache)

	// 带重试的测试连接
	var lastErr error
	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		_, err := client.Ping(ctx).Result()
		cancel()

		if err == nil {
			// 启动健康检查协程
			go monitorRedisHealth(client, cfg.Host, cfg.Port, cfg.DB)
			return client, nil
		}

		lastErr = err
		time.Sleep(time.Second * time.Duration(i+1))
	}

	return nil, fmt.Errorf("无法连接Redis [host:%s port:%s db:%d]: %w",
		cfg.Host, cfg.Port, cfg.DB, lastErr)
}

// monitorRedisHealth 监控Redis连接健康状态
func monitorRedisHealth(client *redis.Client, host, port string, db int) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		_, err := client.Ping(ctx).Result()
		cancel()

		if err != nil {
			log.Printf("Redis连接健康检查失败 [host:%s port:%s db:%d]: %v",
				host, port, db, err)
		}
	}
}

/*
Get
封装redis get方法：
Redis Get 命令用于获取指定 key 的值(Get只能用于字符串值)

return值：
key存在：返回与 key 相关联的字符串值
key不存在：nil
key的值不是字符串：返回error
*/
func Get(key string) (any, error) {
	return RedisClient.Get(context.Background(), key).Result()
}
