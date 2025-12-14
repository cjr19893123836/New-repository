package http_controller

import (
	"Supply/Supply_and_Demand/config"
	"Supply/Supply_and_Demand/controller"
	"Supply/Supply_and_Demand/dao"
	"Supply/Supply_and_Demand/http_models"
	"Supply/Supply_and_Demand/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	controller.BaseApi
	Service *service.UserService
	Dao     *dao.UserDao
}

/*
NewUserController :

	创建用户控制器实例
*/
func NewUserController(svc *service.UserService, dao *dao.UserDao) *UserController {
	return &UserController{
		BaseApi: controller.NewBaseApi(),
		Service: svc,
		Dao:     dao,
	}
}

// UserRegister 用户注册
func (c UserController) UserRegister(ctx *gin.Context) {
	registerReq := http_models.UserRegister{}
	if err := ctx.ShouldBindJSON(&registerReq); err != nil {
		ctx.JSON(config.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	userID, err := c.Service.UserRegister(registerReq)
	if err != nil {
		ctx.JSON(config.UserRegisterFail, gin.H{
			"error":   "注册失败",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(config.UserLoginSuccess, gin.H{
		"message": "注册成功",
		"user_id": userID,
	})
}

// UserLoginByPhone 手机号登录
func (c UserController) UserLoginByPhone(ctx *gin.Context) {
	loginReq := http_models.UserLoginByPhone{}
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(config.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	loginData, err := c.Service.UserLoginByPhone(&loginReq)
	if err != nil {
		ctx.JSON(config.UserLoginFail, gin.H{
			"error":   "登录失败",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(config.UserLoginSuccess, gin.H{
		"message": "登录成功",
		"data":    loginData,
	})
}

// UserLoginByEmail 邮箱登录
func (c UserController) UserLoginByEmail(ctx *gin.Context) {
	loginReq := http_models.UserLoginByEmail{}
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(config.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	loginData, err := c.Service.UserLoginByEmail(&loginReq)
	if err != nil {
		ctx.JSON(config.UserLoginFail, gin.H{
			"error":   "登录失败",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(config.UserLoginSuccess, gin.H{
		"message": "登录成功",
		"data":    loginData,
	})
}

// GetUserProfile 获取用户信息(需要认证)
func (c UserController) GetUserProfile(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(config.Unauthorized, gin.H{
			"error": "用户未认证",
		})
		ctx.Abort() //终止流程
		return
	}
	ctx.JSON(config.Success, gin.H{
		"message": "获取用户信息成功",
		"data":    user,
	})
}
