package middleware

import (
	"Supply_and_Demand/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

/*
CorsMiddleware 跨域资源共享中间件

参数:

	cfg - 应用配置

返回值:

	gin.HandlerFunc - CORS中间件函数

功能:
1. 根据配置设置允许的跨域请求来源
2. 定义允许的HTTP方法
3. 配置允许的请求头
4. 设置预检请求缓存时间
5. 处理OPTIONS预检请求
*/
func CorsMiddleware(cfg *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 如果是 WebSocket 协议升级请求，直接放行
		if strings.Contains(c.GetHeader("Connection"), "Upgrade") &&
			strings.Contains(c.GetHeader("Upgrade"), "theWebsocket") {
			c.Next()
			return
		}

		// 设置允许的来源
		if cfg.Server.AllowOrigin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", cfg.Server.AllowOrigin)
			if cfg.Server.AllowOrigin != "*" {
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		} else {
			// 默认根据运行模式设置
			if gin.Mode() == gin.ReleaseMode {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "https://yourdomain.com")
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			} else {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			}
		}

		// 设置允许的方法
		methods := cfg.Server.AllowMethods
		if methods == "" {
			methods = "GET, POST, PUT, DELETE, OPTIONS"
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", methods)

		// 设置允许的请求头
		headers := cfg.Server.AllowHeaders
		if headers == "" {
			headers = "Content-Type, Authorization"
		}
		c.Writer.Header().Set("Access-Control-Allow-Headers", headers)

		// 设置预检请求缓存时间
		maxAge := cfg.Server.CorsMaxAge
		if maxAge == 0 {
			maxAge = 86400
		}
		c.Writer.Header().Set("Access-Control-Max-Age", strconv.Itoa(maxAge))

		// 处理OPTIONS预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
