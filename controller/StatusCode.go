package controller

// 统一状态码
const (
	UNAUTHORIZEDOPERATION_ERROR_CODE = 5003 // 请求无权限
	TOKENPARSE_ERROR_CODE            = 5008 // Token解析失败
	TOKENCHECK_ERROR_CODE            = 5009 // Token校验失败

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
