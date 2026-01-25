package dao

import (
	"Supply_and_Demand/http_models"

	"gorm.io/gorm"
	"Supply_and_Demand/global"
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
