package middleware

import (
	"Supply_and_Demand/config"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
Logger Gin日志中间件

参数:

	cfg - 应用配置

返回值:

	gin.HandlerFunc - 日志记录中间件函数

功能:
1. 记录请求方法、路径、客户端IP
2. 记录响应状态码
3. 记录请求处理耗时
4. 根据运行模式调整日志详细程度
*/
func Logger(cfg *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		startTime := time.Now()

		// 请求信息
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		query := c.Request.URL.RawQuery

		// 记录完整路径（包含查询参数）
		fullPath := path
		if query != "" {
			fullPath = path + "?" + query
		}

		// 处理请求
		c.Next()

		// 计算处理耗时
		latency := time.Since(startTime)

		// 获取响应状态码
		statusCode := c.Writer.Status()

		// 获取请求ID（如果存在）
		requestID := c.GetString("request_id")
		if requestID == "" {
			requestID = "-"
		}

		// 根据状态码获取日志级别
		level := getLogLevel(statusCode)

		// 格式化日志输出
		logMessage := fmt.Sprintf("[%s] %s | %3d | %13v | %15s | %-7s %s",
			requestID,
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			fullPath,
		)

		// 根据日志级别输出日志
		switch level {
		case "ERROR":
			log.Printf("[ERROR] %s", logMessage)
			// 在调试模式下输出更多错误详情
			if gin.Mode() == gin.DebugMode {
				log.Printf("[ERROR] Request Body: %s", c.GetString("request_body"))
			}
		case "WARN":
			log.Printf("[WARN] %s", logMessage)
		default:
			// Release模式下只记录重要信息
			if gin.Mode() == gin.DebugMode {
				log.Printf("[INFO] %s", logMessage)
			}
		}
	}
}

// getLogLevel 根据HTTP状态码返回日志级别
func getLogLevel(statusCode int) string {
	switch {
	case statusCode >= http.StatusInternalServerError:
		return "ERROR"
	case statusCode >= http.StatusBadRequest:
		return "WARN"
	default:
		return "INFO"
	}
}
