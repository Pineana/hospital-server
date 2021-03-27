package main

import (
	"github.com/gin-gonic/gin"
	"hospital-admin/config"
	adminHttp "hospital-admin/http"
	"hospital-admin/model"
	"hospital-admin/usecase"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Use(Cors())
	db := config.InitGormMysql(config.GlobalConfig.Mysql)
	db.AutoMigrate(&model.User{}, model.Doctor{}, model.Patient{}, model.Drug{}, model.Register{})
	usecase := usecase.NewUseCase(db)
	handle := adminHttp.NewHandle(usecase)
	// 初始化路由
	adminHttp.BuildRouter(engine, handle)
	// 启动服务器 监听80端口
	engine.Run(":8088")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
