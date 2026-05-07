package http_models

type LoginByEmailRep struct {
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Code        string `json:"code" binding:"required,len=6"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
type LoginByPhoneNumberRep struct {
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required,len=6"`
}

type RegisterByPhone struct {
	UserName    string `json:"user_name" binding:"required,min=3,max=20"`
	Password    string `json:"password" binding:"required,min=8,max=20"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required,len=6"`
}
type RegisterByEmail struct {
	UserName    string `json:"user_name" binding:"required,min=3,max=20"`
	Password    string `json:"password" binding:"required,min=8,max=20"`
	Email       string `json:"email" binding:"required"`
	Code        string `json:"code" binding:"required,len=6"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
type UpdateUserNameRep struct {
	UserName string `json:"user_name" binding:"required,min=3,max=20"`
}
type UpdateAvatarRep struct {
	Avatar string `json:"Avatar" binding:"required"`
}
type SendSmsCodeRep struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// VerifySmsCodeRep 校验短信验证码是否正确
type VerifySmsCodeRep struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required,len=6"`
}
type RegisterBySmsReq struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	UserName    string `json:"user_name" binding:"required,min=3,max=20"`
	Password    string `json:"password" binding:"required,min=8,max=20"`
	Code        string `json:"code" binding:"required,len=6"`
}
type LoginBySmsRep struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required,len=6"`
}
type LoginAndRegister struct {
	Token    string `json:"token" `
	ID       uint   `json:"id" `
	UserName string `json:"username" `
	Avatar   string `json:"avatar"`
	Expire   int64  `json:"expire" `
	Code     string `json:"code"`
}
