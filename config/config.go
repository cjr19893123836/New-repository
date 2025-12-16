package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server string `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	JWT    JWT    `yaml:"jwt"`
}

type Mysql struct {
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
}
type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}
type JWT struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Host     string `yaml:"host"`     // 数据库地址
	Port     string `yaml:"port"`     // 数据库端口
	User     string `yaml:"user"`     // 用户名
	Password string `yaml:"password"` // 密码
	DBName   string `yaml:"dbname"`   // 数据库名
	Charset  string `yaml:"charset"`  // 字符集
}

// RedisConfig redis配置
type RedisConfig struct {
	Host     string `yaml:"host"`     // Redis地址
	Port     string `yaml:"port"`     // Redis端口
	Password string `yaml:"password"` // 密码
	DB       int    `yaml:"db"`       // 数据库索引
}

// RateLimitConfig 限流配置
type RateLimitConfig struct {
	Enable   bool    `yaml:"enable"`   // 是否启用
	Rate     float64 `yaml:"rate"`     // 令牌生成速率(个/秒)
	Capacity int     `yaml:"capacity"` // 令牌桶容量
}

// RouteConfig 路由配置
type RouteConfig struct {
	Path    string   `yaml:"path"`    // 路由路径
	Methods []string `yaml:"methods"` // 允许的HTTP方法
	Handler string   `yaml:"handler"` // 处理函数名称
}

// RouteGroupConfig 路由组配置
type RouteGroupConfig struct {
	Prefix string        `yaml:"prefix"` // 路由组前缀
	Routes []RouteConfig `yaml:"routes"` // 路由列表
}

// JWTConfig JWT认证配置
type JWTConfig struct {
	Secret    string   `yaml:"secret"`    // 密钥
	Expire    int      `yaml:"expire"`    // 过期时间(小时)
	Issuer    string   `yaml:"issuer"`    // 签发者
	Whitelist []string `yaml:"whitelist"` // 认证豁免路径前缀
}

type ServerConfig struct {
	Port         string `yaml:"port"`          // 服务端口
	Mode         string `yaml:"mode"`          // 运行模式(debug/release)
	Timeout      int    `yaml:"timeout"`       // 请求超时(秒)
	AllowOrigin  string `yaml:"allow_origin"`  // 允许的跨域来源，默认为*或根据模式自动设置
	AllowMethods string `yaml:"allow_methods"` // 允许的HTTP方法，默认为"GET,POST,PUT,DELETE,OPTIONS"
	AllowHeaders string `yaml:"allow_headers"` // 允许的请求头
	CorsMaxAge   int    `yaml:"cors_max_age"`  // 预检请求缓存时间(秒)，默认为86400
}

// AppConfig 应用配置
type AppConfig struct {
	MySQL      MySQLConfig        `yaml:"mysql"`
	Redis      RedisConfig        `yaml:"redis"`
	Server     ServerConfig       `yaml:"server"`
	RateLimit  RateLimitConfig    `yaml:"rate_limit"`
	JWT        JWTConfig          `yaml:"jwt"`
	HTTPRoutes []RouteGroupConfig `yaml:"http_routes"` // HTTP路由配置
	WSRoutes   []RouteGroupConfig `yaml:"ws_routes"`   // WebSocket路由配置
}

// LoadConfig 加载配置文件
func LoadConfig(path string) (*AppConfig, error) {
	config := &AppConfig{}
	file, err := os.ReadFile(path)
	// 配置文件名
	viper.SetConfigName("config")
	// 配置文件类型
	viper.SetConfigType("yaml")
	// 配置文件路径
	viper.AddConfigPath("Supply_and_Demand/config")
	// 报错处理
	_ = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return config, nil
}
