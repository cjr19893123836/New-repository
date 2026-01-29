package dao

import (
	"Supply_and_Demand/http_models"
	"math/rand"
	"time"

	"Supply_and_Demand/global"

	"gorm.io/gorm"
)

type ProductDao struct {
	Orm *gorm.DB
}

func NewProductDao() *ProductDao {
	return &ProductDao{Orm: global.DB}
}

func (m *ProductDao) CreateProduct(product http_models.Product) error {
	err := m.Orm.Create(&product).Error
	return err
}

func (m *ProductDao) GetProductByID(id int64) (http_models.Product, error) {
	product := http_models.Product{}
	err := m.Orm.First(&product, id).Error
	return product, err
}

func (m *ProductDao) GetRandomProducts(limit int) ([]http_models.Product, error) {
	var products []http_models.Product
	// 获取商品总数
	var count int64
	if err := m.Orm.Model(&http_models.Product{}).Count(&count).Error; err != nil {
		return nil, err
	}
	if count == 0 {
		return products, gorm.ErrRecordNotFound
	}
	// 确保limit不超过总数
	if limit > int(count) {
		limit = int(count)
	}
	// 生成随机偏移量
	rand.Seed(time.Now().UnixNano())
	maxOffset := count - int64(limit)
	var offset int64
	if maxOffset > 0 {
		offset = rand.Int63n(maxOffset)
	} else {
		offset = 0
	}
	// 使用偏移量查询多条记录
	err := m.Orm.Offset(int(offset)).Limit(limit).Find(&products).Error
	return products, err
}
