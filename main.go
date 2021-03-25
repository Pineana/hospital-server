package main

import (
	"github.com/gin-gonic/gin"
	"hospital-admin/config"
	"hospital-admin/http"
	"hospital-admin/model"
	"hospital-admin/usecase"
)

func main() {
	engine := gin.Default()
	db := config.InitGormMysql(config.GlobalConfig.Mysql)
	db.AutoMigrate(&model.User{}, model.Doctor{}, model.Patient{})
	usecase := usecase.NewUseCase(db)
	handle := http.NewHandle(usecase)
	// 初始化路由
	http.BuildRouter(engine, handle)
	// 启动服务器 监听80端口
	engine.Run(":8123")
}
