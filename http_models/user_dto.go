package http_models

// UserRegister 用户注册请求
type UserRegister struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required,len=11"`
}

// UserLoginByPhone 手机号登录
type UserLoginByPhone struct {
	Phone    string `json:"phone" binding:"required,len=11"`
	Password string `json:"password" binding:"required,min=6"` //密码，私密，不序列化到json
}

// UserLoginByEmail 邮箱登录
type UserLoginByEmail struct {
	Email    string `json:"email" binding:"required,email"` //邮箱，唯一索引
	Password string `json:"password" binding:"required,min=6"`
}

// LoginSuccessData 登录成功返回数据---包括token
type LoginSuccessData struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Expire   int64  `json:"expire"` //token过期时间戳
}
