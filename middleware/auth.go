/*
认证中间件模块

该文件定义了:
1. JWT认证中间件
2. 令牌验证流程
3. 认证豁免路由
4. 用户信息上下文设置
*/

package middleware

import (
	"Supply_and_Demand/config"
	"Supply_and_Demand/controller"
	"Supply_and_Demand/utils"
	"fmt" // 格式化输出
	"strconv"
	"strings" // 字符串处理

	"github.com/gin-gonic/gin" // Gin Web框架
)

/*
AuthMiddleware JWT认证中间件

参数:
- cfg *config.AppConfig: 应用配置，包含JWT密钥等信息

返回值:
- gin.HandlerFunc: Gin中间件函数

功能:
1. 检查请求头中的Authorization令牌
2. 验证JWT令牌的有效性
3. 解析令牌中的用户信息
4. 将用户信息存入请求上下文
5. 处理认证失败情况
*/

const (
	TOKEN_NAME   = "Authorization"
	TOKEN_PREFIX = "Bearer "
)

func AuthMiddleware(cfg *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 如果是 WebSocket 协议升级请求，直接放行
		if strings.Contains(c.GetHeader("Connection"), "Upgrade") &&
			strings.Contains(c.GetHeader("Upgrade"), "websocket") {
			c.Next()
			return
		}
		// 检查白名单路由
		for _, prefix := range cfg.JWT.Whitelist {
			if strings.HasPrefix(c.Request.URL.Path, prefix) {
				c.Next()
				return
			}
		}

		// 从请求头获取Authorization字段
		authHeader := c.GetHeader(TOKEN_NAME)
		if authHeader == "" {
			controller.Fail(c, controller.ResponseJson{
				Code: controller.UNAUTHORIZEDOPERATION_ERROR_CODE,
				Msg:  "缺失相应权限",
			})
			return
		}

		// 解析Bearer令牌
		tokenString := strings.TrimPrefix(authHeader, TOKEN_PREFIX)
		c.Set("inToken", tokenString)
		token, err := utils.ParseToken(tokenString)
		userExit := token.UserID

		// 检查令牌是否成功解析
		if err != nil || userExit == 0 {
			controller.Fail(c, controller.ResponseJson{
				Code: controller.TOKENPARSE_ERROR_CODE,
				Msg:  "token解析失败",
			})
			return
		}

		// 验证令牌是否有效
		tokenValid := strings.Replace("LOGIN_USER_{id}", "{id}", strconv.Itoa(int(userExit)), -1)
		stUserToken, err := config.Get(tokenValid)

		if tokenString != stUserToken || err != nil {
			controller.Fail(c, controller.ResponseJson{
				Code: controller.TOKENCHECK_ERROR_CODE,
				Msg:  "登录已失效,请重新登录",
			})
			return
		}
		fmt.Println(userExit)

		// 将用户ID存入上下文
		c.Set("userID", uint(userExit))

		// 继续处理请求
		c.Next()
	}
}
