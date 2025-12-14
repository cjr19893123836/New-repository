package controller

import "github.com/gin-gonic/gin"

type BaseApi struct {
	// 存储反馈给客户端的数据
	Ctx *gin.Context
	// 存储错误信息(多个)
	Errors error
}

func NewBaseApi() BaseApi {
	return BaseApi{}
}
