package router

import (
	"Supply_and_Demand/config"
	"Supply_and_Demand/controller/http_controller"
	"Supply_and_Demand/middleware"
	"Supply_and_Demand/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
SetupRouter 初始化并返回Gin路由引擎

参数:

	userService - 用户服务实例
	wsService - WebSocket服务实例
	redisClient - Redis连接实例
	cfg - 应用配置

返回值:

	*gin.Engine - 配置好的路由引擎

功能:
1. 设置Gin运行模式
2. 配置全局中间件
3. 初始化API路由
4. 设置WebSocket支持
5. 开发模式下启用Swagger文档
*/

func SetupRouter(
	userService *service.UserService,
	productService *service.ProductService,
	cfg *config.AppConfig) *gin.Engine {
	// 设置Gin运行模式
	gin.SetMode(cfg.Server.Mode)

	// 创建Gin引擎实例
	r := gin.New()

	// 配置全局中间件
	setupMiddlewares(r, cfg)
	// 初始化API控制器
	userController := initControllersUser(
		userService)
	productController := initControllersProduct(
		productService)
	// 配置HTTP路由
	setupHTTPRoutesUser(r, userController, cfg)       //用户
	setupHTTPRoutesProduct(r, productController, cfg) //产品

	// 配置WebSocket路由
	//setupWebSocketRoutes(r,wsService, redisClient, cfg)

	return r
}

/*
setupMiddlewares 配置全局中间件

参数:

	r - Gin引擎实例
	cfg - 应用配置

功能:
1. 自定义日志格式
2. 注册全局中间件
*/
func setupMiddlewares(r *gin.Engine, cfg *config.AppConfig) {
	r.Use(
		gin.Recovery(),                 // 恐慌恢复
		middleware.CorsMiddleware(cfg), // CORS支持
		middleware.RateLimit(cfg),      // 限流
		middleware.AuthMiddleware(cfg), // 认证
		middleware.Logger(cfg),					// log日志
	)
}

/*
initControllers 初始化控制器

参数:

	userService - 用户服务实例
	wsService - WebSocket服务实例

返回值:

	*api.UserController - 用户控制器
	*api.WebSocketController - WebSocket控制器

功能:
1. 初始化用户控制器
2. 初始化WebSocket控制器
*/

func initControllersUser(
	userService *service.UserService,
) *http_controller.UserController {
	fmt.Println("正在初始化用户控制器...")
	userController := http_controller.NewUserController(userService)
	fmt.Println("用户控制器初始化完成")

	return userController
}

func initControllersProduct(
	productService *service.ProductService,
) *http_controller.ProductController {
	fmt.Println("正在初始化产品控制器...")
	productController := http_controller.NewProductController(productService)
	fmt.Println("用户控制器初始化完成")

	return productController
}

/*
setupHTTPRoutes 配置HTTP路由

参数:

	r - Gin引擎实例
	userController - 用户控制器实例

功能:
1. 配置/api/v1路由分组
2. 定义用户相关路由
*/
func setupHTTPRoutesUser(r *gin.Engine, userController *http_controller.UserController, cfg *config.AppConfig) {
	// 从配置加载HTTP路由

	for _, group := range cfg.HTTPRoutes {
		//只处理 /api/v1/users 路由组
		if group.Prefix != "/api/v1/users" {
			continue
		}
		routerGroup := r.Group(group.Prefix)
		for _, route := range group.Routes {
			// 根据handler名称映射到控制器方法

			handler := getHTTPHandlerUser(userController, route.Handler)
			if handler == nil {
				fmt.Println("警告: 未找到用户HTTP路由处理函数", route.Handler)
				continue
			}

			// 注册路由
			for _, method := range route.Methods {
				switch method {
				case "GET":
					routerGroup.GET(route.Path, handler)
				case "POST":
					routerGroup.POST(route.Path, handler)
				case "PUT":
					routerGroup.PUT(route.Path, handler)
				case "DELETE":
					routerGroup.DELETE(route.Path, handler)
				default:
					fmt.Println("警告: 不支持的HTTP方法", method)
				}
			}
		}
	}
}

func setupHTTPRoutesProduct(r *gin.Engine, productController *http_controller.ProductController, cfg *config.AppConfig) {
	// 从配置加载HTTP路由

	for _, group := range cfg.HTTPRoutes {
		//只处理 /api/v1/products 路由组
		if group.Prefix != "/api/v1/products" {
			continue
		}
		routerGroup := r.Group(group.Prefix)
		for _, route := range group.Routes {
			// 根据handler名称映射到控制器方法

			handler := getHTTPHandlerProduct(productController, route.Handler)
			if handler == nil {
				fmt.Println("警告: 未找到产品HTTP路由处理函数", route.Handler)
				continue
			}

			// 注册路由
			for _, method := range route.Methods {
				switch method {
				case "GET":
					routerGroup.GET(route.Path, handler)
				case "POST":
					routerGroup.POST(route.Path, handler)
				case "PUT":
					routerGroup.PUT(route.Path, handler)
				case "DELETE":
					routerGroup.DELETE(route.Path, handler)
				default:
					fmt.Println("警告: 不支持的HTTP方法", method)
				}
			}
		}
	}
}

// getHTTPHandlerUser 根据名称获取对应的处理函数
func getHTTPHandlerUser(inController *http_controller.UserController, name string) gin.HandlerFunc {
	switch name {
	case "UserRegister":
		return inController.UserRegister //注册接口:POST请求+/api/v1/register->对应的USerRegister函数
	case "UserLoginByPhone":
		return inController.UserLoginByPhone //手机号登录接口:POST请求+/api/v1/login/phone->对应的USerLoginByPhone函数
	case "UserLoginByEmail":
		return inController.UserLoginByEmail //邮箱登录接口:POST请求+/api/v1/login/login->对应的USerLoginByLogin函数
	default:
		return nil
	}
}

// getHTTPHandlerProduct 根据名称获取对应的处理函数
func getHTTPHandlerProduct(inController *http_controller.ProductController, name string) gin.HandlerFunc {
	switch name {
	case "CreateProduct":
		return inController.CreateProduct
	case "GetRandomProduct":
		return inController.GetRandomProduct
	default:
		return nil
	}
}
