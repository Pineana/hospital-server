package main

import (
	"github.com/gin-gonic/gin"
	"hospital-admin/config"
	"hospital-admin/http"
	"hospital-admin/usecase"
)

func main() {
	engine := gin.Default()
	db := config.InitMysql(config.GlobalConfig.Mysql)

	usecase := usecase.NewUseCase(db)
	handle := http.NewHandle(usecase)
	// 初始化路由
	http.BuildRouter(engine, handle)
	// 启动服务器 监听80端口
	engine.Run(":8123")
}
