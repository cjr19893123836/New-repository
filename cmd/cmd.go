package cmd

import (
	"Supply/Supply_and_Demand/config"
	"Supply/Supply_and_Demand/global"
	"Supply/Supply_and_Demand/router"
	"Supply/Supply_and_Demand/service"
	"Supply/Supply_and_Demand/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	var initErr error
	// 初始化并验证系统配置文件
	fmt.Println()
	fmt.Println("正在加载配置文件...")
	// cfg 包含所有服务配置信息，包括服务器、数据库、JWT等
	cfg, err := config.LoadConfig("config/config.yaml")

	fmt.Println("查看配置文件")

	fmt.Println(cfg)

	if err != nil {
		fmt.Printf("加载配置失败: %v", err)
	}

	fmt.Println("配置文件加载成功，开始验证配置...")
	// 验证必要配置项是否设置
	if cfg.Server.Port == "" {
		fmt.Println("服务器端口配置缺失")
	} // 验证必要配置项是否设置
	if cfg.Server.Port == "" {
		fmt.Println("服务器端口配置缺失")
	}
	//if cfg.JWT.Secret == "your-secret-key" {
	//	global.Logger.Error("请修改默认JWT密钥，使用强密码")
	//}
	if cfg.MySQL.Host == "" || cfg.MySQL.DBName == "" {
		fmt.Println("数据库配置不完整，请检查MySQL主机和数据库名配置")
	}
	fmt.Println("配置验证通过")

	// 初始化数据库连接
	fmt.Println()
	fmt.Println("正在初始化MySQL连接...")
	// db 是MySQL数据库连接实例
	db, err := config.InitMysql(cfg.MySQL)
	if err != nil {
		fmt.Printf("初始化MySQL失败: %v", err)
	}
	// 错误处理
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	if initErr != nil {
		panic(initErr.Error())
	}
	global.DB = db
	fmt.Println("MySQL连接成功")

	// 初始化redis连接
	fmt.Println()
	fmt.Println("正在初始化Redis连接...")
	// redisClient 是Redis连接实例，用于缓存和会话管理
	config.RedisClient, err = config.ToInitRedis(cfg.Redis)
	// 错误处理
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	if initErr != nil {
		panic(initErr.Error())
	}
	fmt.Println("Redis连接成功")

	// ---------- 初始化 Service ----------
	userService := service.NewUserService()
	// 初始化路由
	fmt.Println("正在初始化路由...")
	r := router.SetupRouter(userService, cfg)
	fmt.Println("路由初始化完成...")

	////定时任务执行
	//go AllTask()

	// 创建HTTP服务器实例
	// srv 配置了监听地址和路由处理器
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port, // 监听地址和端口
		Handler: r,                     // 路由处理器
	}

	// 优雅关闭机制
	// done 通道用于通知主线程可以安全退出
	done := make(chan bool)
	// quit 通道用于接收系统中断信号
	quit := make(chan os.Signal, 1)
	// 监听SIGINT(CTRL+C)和SIGTERM(kill)信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 启动goroutine处理优雅关闭
	f := func() {
		// 等待退出信号
		<-quit
		fmt.Println("服务器正在关闭...")

		// 创建30秒超时的上下文
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// 优雅关闭HTTP服务器
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("强制关闭服务器: %v", err)
		}

		// 关闭数据库连接
		if sqlDB, err := db.DB(); err == nil {
			err := sqlDB.Close()
			if err != nil {
				return
			}
		}
		//redisClient.Close()

		// 通知主线程可以退出
		close(done)
	}
	go f()

	// 启动HTTP服务器
	fmt.Println("服务器启动，监听端口:", cfg.Server.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("服务器启动失败: %v", err)
	}

	// 等待优雅关闭完成
	<-done
}

func Clean() {
	fmt.Println("服务器已关闭")
}
