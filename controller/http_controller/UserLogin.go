package http_controller

import (
	"Supply_and_Demand/controller"
	"Supply_and_Demand/http_models"
	"Supply_and_Demand/service"
	"Supply_and_Demand/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserLoginController struct {
	controller.BaseApi
	Service *service.UserLoginService
}

func NewController(svc *service.UserLoginService) *UserLoginController {
	return &UserLoginController{
		BaseApi: controller.NewBaseApi(),
		Service: svc,
	}
}

// RegisterByPhone 注册
func (c *UserLoginController) RegisterByPhone(ctx *gin.Context) {
	Register := http_models.RegisterByPhone{}
	if err := ctx.ShouldBind(&Register); err != nil {
		ctx.JSON(controller.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	userID, err := c.Service.RegisterByPhone(Register)
	if err != nil {
		ctx.JSON(controller.UserRegisterFail, gin.H{
			"error":   "注册失败",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(controller.UserLoginSuccess, gin.H{
		"message": "注册成功",
		"user_id": userID,
	})
}

// RegisterByEmail 注册
func (c *UserLoginController) RegisterByEmail(ctx *gin.Context) {
	Register := http_models.RegisterByEmail{}
	if err := ctx.ShouldBind(&Register); err != nil {
		ctx.JSON(controller.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	userID, err := c.Service.RegisterByEmail(Register)
	if err != nil {
		ctx.JSON(controller.UserRegisterFail, gin.H{
			"error":   "注册失败",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(controller.UserLoginSuccess, gin.H{
		"message": "注册成功",
		"user_id": userID,
	})
}

// LoginUserByPhoneNumber 手机号登录
func (c *UserLoginController) LoginUserByPhoneNumber(ctx *gin.Context) {
	LoginByPhoneNumber := http_models.LoginByPhoneNumberRep{}
	if err := ctx.ShouldBind(&LoginByPhoneNumber); err != nil {
		ctx.JSON(controller.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	data, err := c.Service.LoginUserByPhone(LoginByPhoneNumber)
	if err != nil {
		ctx.JSON(controller.UserRegisterFail, gin.H{
			"error":   "登录失败",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(controller.UserLoginSuccess, gin.H{
		"message": "登录成功",
		"data":    data,
	})
}

// LoginUserByEmail 邮箱登录
func (c *UserLoginController) LoginUserByEmail(ctx *gin.Context) {
	LoginByEmail := http_models.LoginByEmailRep{}
	if err := ctx.ShouldBind(&LoginByEmail); err != nil {
		ctx.JSON(controller.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	data, err := c.Service.LoginUserByEmail(LoginByEmail)
	if err != nil {
		ctx.JSON(controller.UserRegisterFail, gin.H{
			"error":   "登录失败",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(controller.UserLoginSuccess, gin.H{
		"message": "登录成功",
		"data":    data,
	})
}

func (c *UserLoginController) UpdateUserName(ctx *gin.Context) {
	// 手动解析 token
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(401, gin.H{"error": "未登录"})
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(401, gin.H{"error": "认证格式错误"})
		return
	}
	claims, err := utils.ParseToken(parts[1])
	if err != nil {
		ctx.JSON(401, gin.H{"error": "token无效"})
		return
	}
	userID := claims.UserID
	UpdateName := http_models.UpdateUserNameRep{}
	if err := ctx.ShouldBind(&UpdateName); err != nil {
		ctx.JSON(controller.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	err = c.Service.UpdateUserName(userID, UpdateName.UserName)
	if err != nil {
		ctx.JSON(controller.UserRegisterFail, gin.H{
			"error":   "修改失败",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(controller.UserLoginSuccess, gin.H{
		"message": "修改成功",
	})
}

// UpdateAvatar 修改头像
func (c *UserLoginController) UpdateAvatar(ctx *gin.Context) {
	//先获取用户的ID
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(401, gin.H{"error": "未登录"})
		return
	}
	//手动解析token
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(401, gin.H{"error": "认证格式错误"})
		return
	}
	claims, err := utils.ParseToken(parts[1])
	if err != nil {
		ctx.JSON(401, gin.H{"error": "token无效"})
		return
	}
	userID := claims.UserID
	//从http真的获取头像文件
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(controller.BadRequest, gin.H{
			"error": "未获取到头像文件",
		})
		return
	}
	//验证头像的文件格式
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"image/bmp":  true,
	}
	contentType := file.Header.Get("Content-Type")
	if !allowedTypes[contentType] {
		ctx.JSON(controller.BadRequest, gin.H{
			"error": "文件类型不符合要求"})
		return
	}
	//验证文件的大小
	if file.Size > 5*1024*1024 {
		ctx.JSON(controller.BadRequest, gin.H{
			"error": "你的文件大小超出了要求"})
		return
	}
	//调用Service,保存地址，让前端知道去哪里找地址
	avatarUrl, err := c.Service.UpdateUserAvatar(userID, file)
	if err != nil {
		ctx.JSON(controller.UserRegisterFail, gin.H{
			"errors":  "保存失败",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(controller.UserLoginSuccess, gin.H{
		"message": "保存成功",
		"avatar":  avatarUrl,
	})
}

// SendSmsCode 发送验证码
func (c *UserLoginController) SendSmsCode(ctx *gin.Context) {
	//传入前端的参数
	var req http_models.SendSmsCodeRep
	//把它转化为结构体
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(controller.BadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}
	//调用Service
	err = c.Service.SendSmsCode(req.PhoneNumber)
	if err != nil {
		ctx.JSON(controller.UserRegisterFail, gin.H{
			"error":   "发送失败",
			"message": err.Error()})
		return
	}
	//返回成功
	ctx.JSON(controller.UserLoginSuccess, gin.H{
		"message": "验证码已经成功发送"})
}

// LoginBySmsCode 验证码登录
func (c *UserLoginController) LoginBySmsCode(ctx *gin.Context) {
	//传入前端的参数
	var req http_models.LoginBySmsRep
	//把它转化为结构体
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(controller.BadRequest, gin.H{
			"errors": "请求参数错误"})
		return
	}
	//调用Service
	data, err := c.Service.LoginBySmsCode(req)
	if err != nil {
		ctx.JSON(controller.UserRegisterFail, gin.H{
			"error":   "登录失败",
			"message": err.Error()})
		return
	}
	//返回成功+token
	ctx.JSON(controller.UserLoginSuccess, gin.H{
		"message": "登录成功",
		"data":    data})
}

// RegisterBySms 短信验证码注册
func (c *UserLoginController) RegisterBySms(ctx *gin.Context) {
	// 接收前端的参数
	var req http_models.RegisterBySmsReq
	// 转化为结构体
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(controller.BadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}
	// 调用Service
	data, err := c.Service.RegisterBySms(req)
	if err != nil {
		ctx.JSON(controller.UserRegisterFail, gin.H{
			"error":   "注册失败",
			"message": err.Error(),
		})
		return
	}
	// 成功响应
	ctx.JSON(controller.UserLoginSuccess, gin.H{
		"message": "注册成功",
		"data":    data,
	})
}
