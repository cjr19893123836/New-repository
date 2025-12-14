package middleware

import (
	"Supply/Supply_and_Demand/config"
	"Supply/Supply_and_Demand/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthHook 中间件的认证
func AuthHook() gin.HandlerFunc {
	return func(ctx *gin.Context) { //返回的是中间件的处理函数
		//从http请求头中获取Authorization字段
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(config.Unauthorized, gin.H{ //给前端返回错误响应(401)
				"error": "Authorization header is empty",
			})
			ctx.Abort() //终止流程
			return
		}
		//将Authorization进行分割,Bearer <token>
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(config.Unauthorized, gin.H{
				"error": "Authorization header is invalid,应为:Bearer <token>",
			})
			ctx.Abort()
			return
		}
		tokenString := parts[1] //提取jwt令牌中的字符串
		//解析令牌
		_, err := utils.ParseToken(tokenString)
		if err != nil {
			ctx.JSON(config.Unauthorized, gin.H{
				"error": "令牌无效或已过期:" + err.Error(),
			})
			ctx.Abort()
			return
		}
		////验证用户是否存在
		//user, err := dao.GetUserByID(claims.UserID)
		//if err != nil {
		//	ctx.JSON(config.Unauthorized, gin.H{
		//		"error": "用户不存在:" + err.Error(),
		//	})
		//	ctx.Abort()
		//	return
		//}
		//// 手动构建用户信息，因为User结构体没有JSON标签了
		//ctx.Set("user", gin.H{
		//	"user_id":    user.UserID,
		//	"username":   user.Username,
		//	"email":      user.Email,
		//	"phone":      user.Phone,
		//	"created_at": user.CreatedAt,
		//	"last_login": user.LastLogin,
		//})
		ctx.Next()
	}
}

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")                              //允许所有来源跨域
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST,PUT,DELETE, OPTIONS") //允许的请求方法,OPTIONS预检请求方法
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")    //// 允许的请求头（需包含前端传递的自定义头，如 Authorization、Content-Type）
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204) //直接返回 204 状态码并终止流程
			return
		}
		ctx.Next() //放行正常请求
	}
}
