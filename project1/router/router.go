package router

import (
	"project1/controller"
	"project1/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default() //创建路由引擎
	//使用跨域中间件
	router.Use(middleware.CORSMiddleware())
	//公共路由(不需要认证）
	public := router.Group("/api/v1") //URL统一前缀未api/v1
	{
		//注册接口:POST请求+/api/v1/register->对应的USerRegister函数
		public.POST("/register", controller.UserRegister)
		//手机号登录接口:POST请求+/api/v1/login/phone->对应的USerLoginByPhone函数
		public.POST("/login/phone", controller.UserLoginByPhone)
		//邮箱登录接口:POST请求+/api/v1/login/login->对应的USerLoginByLogin函数
		public.POST("/login/email", controller.UserLoginByEmail)
	}
	//受保护的路由(jwt认证)
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthHook())
	{ // 获取用户信息接口：GET 请求 + /api/v1/user/profile → 对应 GetUserProfile 函数
		protected.GET("/user/profile", controller.GetUserProfile)
	}
	return router
}
