package service

import (
	"Supply_and_Demand/dto"
	"Supply_and_Demand/http_models"

	"errors"
)

type ProductService struct {
	Dao *dao.ProductDao
}

func NewProductService() *ProductService {
	return &ProductService{
		Dao: dao.NewProductDao(),
	}
}

func (s *ProductService) CreateProduct(product dto.ProductReq) (int64, error) {
	condition, err := product.GetCondition()
	if err != nil {
		return 0, err
	}
	product.Condition = condition

	tradeType, err := product.GetTradeType()
	if err != nil {
		return 0, err
	}
	product.TradeType = tradeType

	transactionMethod, err := product.GetTransactionMethod()
	if err != nil {
		return 0, err
	}
	product.TransactionMethod = transactionMethod

	// 数据转换(DTO->实体)
	productEntity := http_models.Product{
		UserID:        product.UserID,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		IntroText:     product.IntroText,
		OriginPrice:   product.OriginPrice,
		MainImageUrl:  product.MainImageUrl,
		PublishDate:   time.Now(),
		Status:        "待审核",
		ViewCount:     0,
		WantCount:     0,
		Condition:     product.Condition,
		TradeType:     product.TradeType,
		FreeShipping:  product.FreeShipping,
		CollectCount:  0,
		CreateDate:    time.Now(),
	}
	err = s.Dao.CreateProduct(productEntity)

	if err != nil {
		return 0, err
	}

	return productEntity.ID, nil
}