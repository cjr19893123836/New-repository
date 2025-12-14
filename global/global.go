package global

import "gorm.io/gorm"

var (
	// DB 定义一个接收全局DB的对象
	DB *gorm.DB
)
