package middleware

import (
	"Supply/Supply_and_Demand/config"
	"golang.org/x/time/rate"

	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// IPRateLimiter IP限流器
type IPRateLimiter struct {
	ips      map[string]*rate.Limiter
	lastSeen map[string]time.Time
	mu       *sync.RWMutex
	r        rate.Limit
	b        int
}

// cleanupInactiveIPs 定期清理不活跃的IP记录
func (i *IPRateLimiter) cleanupInactiveIPs(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		i.mu.Lock()
		for ip, last := range i.lastSeen {
			if time.Since(last) > interval {
				delete(i.ips, ip)
				delete(i.lastSeen, ip)
			}
		}
		i.mu.Unlock()
	}
}

// NewIPRateLimiter 创建IP限流器
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	limiter := &IPRateLimiter{
		ips:      make(map[string]*rate.Limiter),
		lastSeen: make(map[string]time.Time),
		mu:       &sync.RWMutex{},
		r:        r,
		b:        b,
	}
	// 启动后台清理任务，每小时清理一次
	go limiter.cleanupInactiveIPs(time.Hour)
	return limiter
}

// AddIP 添加IP到限流器
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	i.lastSeen[ip] = time.Now()

	return limiter
}

// GetLimiter 获取IP的限流器
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.RLock()
	limiter, exists := i.ips[ip]
	i.mu.RUnlock()

	if !exists {
		return i.AddIP(ip)
	}

	i.mu.Lock()
	i.lastSeen[ip] = time.Now()
	i.mu.Unlock()

	return limiter
}

var limiter *IPRateLimiter

// RateLimit 限流中间件
func RateLimit(cfg *config.AppConfig) gin.HandlerFunc {
	// 初始化限流器
	if limiter == nil {
		limiter = NewIPRateLimiter(
			rate.Limit(cfg.RateLimit.Rate),
			cfg.RateLimit.Capacity,
		)
	}

	return func(c *gin.Context) {
		if !cfg.RateLimit.Enable {
			c.Next()
			return
		}

		// 获取客户端IP
		ip := c.ClientIP()

		limiter := limiter.GetLimiter(ip)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "请求过于频繁，请稍后再试",
			})
			return
		}

		c.Next()
	}
}
