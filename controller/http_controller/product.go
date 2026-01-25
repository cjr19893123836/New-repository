package http_controller

import (
	"Supply_and_Demand/controller"
	"Supply_and_Demand/http_models"
	"Supply_and_Demand/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	controller.BaseApi
	Service *service.ProductService
}

// CreateProduct 创建商品
func (u *ProductController) CreateProduct(c *gin.Context) {
	var productRep http_models.Product
	if err := c.ShouldBindJSON(&productRep); err != nil {
		ctx.JSON(controller.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}

	productRep.UserID, err := c.Get("userID")
	if err != nil {
		ctx.JSON(controller.Unauthorized, gin.H{
			"error":   "未授权",
			"message": "请先登录",
		})
		return
	}

	productID, err := u.Service.CreateProduct(productRep)
	if err != nil {
		ctx.JSON(controller.InternalServerError, gin.H{
			"error":   "创建商品失败",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(controller.Success, gin.H{
		"message":    "商品创建成功",
		"product_id": productID,
	})
}
