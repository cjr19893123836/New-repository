package http_models

import "gorm.io/gorm"

// Category 商品类别表模型
type Category struct {
	gorm.Model
	Image     string `gorm:"column:image;size:1024"`      //分类图
	SortOrder int    `gorm:"column:sort_order;default:0"` //排序
	TagName   string `gorm:"column:tag_name;"`            //分类名称
}

func (Category) TableName() string {
	return "category"
}
