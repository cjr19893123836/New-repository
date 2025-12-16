package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseApi struct {
	// 存储反馈给客户端的数据
	Ctx *gin.Context
	// 存储错误信息(多个)
	Errors error
}

func NewBaseApi() BaseApi {
	return BaseApi{}
}

/*
统一处理请求参数绑定、错误收集、参数校验和响应返回
*/

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data"`
}

/*
Ok :

	请求成功
*/
func Ok(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

/*
Fail :

	请求失败
*/
func Fail(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

/*
Fail :

	请求失败返回JSON
*/
func (m *BaseApi) Fail(resp ResponseJson) {
	Fail(m.Ctx, resp)
}

/*
Ok :

	请求成功返回JSON
*/
func (m *BaseApi) Ok(resp ResponseJson) {
	Ok(m.Ctx, resp)
}
