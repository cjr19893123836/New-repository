package config

// 统一状态码
const (
	Success       = 200 //通用成功
	BadRequest    = 400 //请求参数错误
	Unauthorized  = 401 //未授权
	Forbidden     = 403 //禁止访问
	NotFound      = 404 //资源未找到
	InternalError = 500 //服务器内部错误
	//用户相关的状态码
	UserLoginSuccess    = 200 //登录成功
	UserLoginFail       = 401 //登录失败
	UserRegisterSuccess = 201 //注册成功
	UserRegisterFail    = 400 //注册失败
)
