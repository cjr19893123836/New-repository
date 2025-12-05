package main

import (
	"fmt"
	"log"
	"project1/config"
	"project1/dao"
	"project1/router"
)

func main() {
	//加载配置文件
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("加载配置文件失败:", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.Account,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.DbName,
	)
	//初始化数据库
	err = dao.InitDB(dsn)
	if err != nil {
		log.Fatal("数据库初始化失败:", err)
	}
	//设置路由
	r := router.SetupRouter()
	//启动服务器
	log.Println("服务端启动在:8080端口")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
