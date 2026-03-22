/*
此文件用于初始化mysql
并且引入logger记录错误
改用viper读取配置
*/

package config

import (
	"Supply_and_Demand/http_models"
	"context"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"runtime"
	"time"
)

//ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

func InitMysql(cfg MySQLConfig) (*gorm.DB, error) {
	// 自定义Logger
	LogMode := logger.Info
	// 没有读取到viper配置则输出错误
	if !viper.GetBool("mode.develop") {
		LogMode = logger.Error
	}

	// 验证数据库密码
	if cfg.Password == "" {
		return nil, fmt.Errorf("数据库密码未设置或使用默认值")
	}

	// 构建MySQL连接字符串(DSN)
	// parseTime=True 确保时间字段正确解析为time.Time
	// loc=Local 设置时区为本地时区
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset)

	// 初始化GORM连接
	// 使用Info级别的日志模式，记录SQL语句
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 单一表名
			SingularTable: true,
			// 如果SingularTable为true的话，表名为sys_ + 表名
			// 如果SingularTable为false的话，表名为蛇形命名
			TablePrefix: "sys_",
		},
		Logger: logger.Default.LogMode(LogMode),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		SkipDefaultTransaction: true, // 关闭默认事务
		PrepareStmt:            true, // 缓存预编译语句
	})
	if err != nil {
		return nil, fmt.Errorf("MySQL连接失败 [host:%s port:%s db:%s user:%s]: %w",
			cfg.Host, cfg.Port, cfg.DBName, cfg.User, err)
	}

	// 获取底层数据库连接实例
	// 用于配置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库实例失败: %w", err)
	}

	// 优化连接池配置
	// 根据CPU核心数设置连接池大小
	numCPU := runtime.NumCPU()
	maxIdle := numCPU * 2
	if maxIdle < 2 {
		maxIdle = 2
	}
	maxOpen := numCPU * 5 // 降低最大连接数
	if maxOpen > 100 {    // 不超过数据库最大连接数限制	sqlDB.SetMaxIdleConns(viper.GetInt("db.max_idle"))
		// 最大开启连接个数
		sqlDB.SetMaxOpenConns(viper.GetInt("db.max_open"))
		// 最长连接时间
		sqlDB.SetConnMaxLifetime(360 * time.Hour)
		maxOpen = 100
	}

	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetConnMaxLifetime(10 * time.Minute) // 更短的生命周期
	sqlDB.SetConnMaxIdleTime(2 * time.Minute)  // 更快的空闲连接释放

	// 测试数据库连接
	// 使用5秒超时上下文，避免长时间阻塞
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("数据库连接测试失败 [host:%s db:%s]: %w",
			cfg.Host, cfg.DBName, err)
	}

	err = autoMigrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// autoMigrate 自动迁移数据库模型
func autoMigrate(db *gorm.DB) error {
	// 定义需要自动迁移的模型列表
	models := []interface{}{
		&http_models.User{},
		&http_models.Product{},
		// 添加其他模型...
	}

	// 执行自动迁移
	return db.AutoMigrate(models...)
}

/*
createMonthlyPartition 创建按月分区表

参数:
- db *gorm.DB: GORM数据库实例
- tableName string: 需要分区的表名

返回值:
- error: 分区过程中出现的错误

功能:
1. 为指定表创建按月分区
2. 适用于数据量大且按时间查询频繁的表
3. 需要数据库支持表分区功能
*/
func createMonthlyPartition(db *gorm.DB, tableName string) error {
	// 实现按月自动分区逻辑
	// 这里需要根据业务需求实现具体分区策略
	// 示例SQL: ALTER TABLE table_name PARTITION BY RANGE (YEAR(created_at)*100 + MONTH(created_at)) (...)
	return nil
}
