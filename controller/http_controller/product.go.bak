package http_controller

import (
	"Supply_and_Demand/controller"
	"Supply_and_Demand/dto"
	"Supply_and_Demand/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	controller.BaseApi
	Service *service.ProductService
}

func NewProductController(svc *service.ProductService) *ProductController {
	return &ProductController{
		BaseApi: controller.NewBaseApi(),
		Service: svc,
	}
}

// CreateProduct 创建商品
func (u *ProductController) CreateProduct(c *gin.Context) {
	var productRep dto.ProductReq
	if err := c.ShouldBindJSON(&productRep); err != nil {
		c.JSON(controller.BadRequest, gin.H{
			"error":   "请求参数错误",
			"message": err.Error(),
		})
		return
	}

	var ok bool
	userIdVal, exist := c.Get("userID")
	if productRep.UserID, ok = userIdVal.(uint); !exist || !ok {
		c.JSON(controller.Unauthorized, gin.H{
			"error":   "未授权",
			"message": "请先登录",
		})
		return
	}

	productID, err := u.Service.CreateProduct(productRep)
	if err != nil {
		c.JSON(controller.InternalError, gin.H{
			"error":   "创建商品失败",
			"message": err.Error(),
		})
		return
	}

	c.JSON(controller.Success, gin.H{
		"message":    "商品创建成功",
		"product_id": productID,
	})
}

// GetRandomProduct 获取随机商品推荐
func (u *ProductController) GetRandomProduct(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	// 可选：限制最大数量
	if limit > 100 {
		limit = 100
	}
	products, err := u.Service.GetRandomProducts(limit)
	if err != nil {
		c.JSON(controller.InternalError, gin.H{
			"error":   "获取推荐商品失败",
			"message": err.Error(),
		})
		return
	}
	// 转换为DTO
	productResps := make([]dto.ProductResp, len(products))
	for i, p := range products {
		productResps[i] = p.ToResp()
	}
	c.JSON(controller.Success, gin.H{
		"message":  "获取推荐商品成功",
		"products": productResps,
	})
}
