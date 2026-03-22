package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	// MySQL DSN格式：用户名:密码@tcp(IP:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local
	DSN := "root:123456@tcp(127.0.0.1:3306)/aaa?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接MySQL并初始化GORM
	var err error
	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("连接MySQL失败：" + err.Error())
	}
	var db *gorm.DB // 自动建表（开发用，生产建议手动建表）
}
