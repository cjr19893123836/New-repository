package controller

import (
	"Supply/Supply_and_Demand/config"
	"project1/http_models"
	"project1/service"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	registerReq := http_models.UserRegister{}
	if err := ctx.ShouldBindJSON(&registerReq); err != nil {
		ctx.JSON(config.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	userID, err := service.UserRegister(registerReq)
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
func UserLoginByPhone(ctx *gin.Context) {
	loginReq := http_models.UserLoginByPhone{}
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(config.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	loginData, err := service.UserLoginByPhone(&loginReq)
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
func UserLoginByEmail(ctx *gin.Context) {
	loginReq := http_models.UserLoginByEmail{}
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(config.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}
	loginData, err := service.UserLoginByEmail(&loginReq)
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
func GetUserProfile(ctx *gin.Context) {
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
