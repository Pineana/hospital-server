package http

import "github.com/gin-gonic/gin"

func BuildRouter(engine *gin.Engine, handle *Handle) {
	SetBaseRouter(engine, handle)
	SetPatientRouter(engine, handle)
}

func SetBaseRouter(engine *gin.Engine, handle *Handle) {
	r := engine.Group("/base")
	r.POST("/login", handle.UserLogin)
	r.POST("/register", handle.UserRegister)
}

func SetPatientRouter(engine *gin.Engine, handle *Handle) {
	r := engine.Group("/patient")
	r.GET("/getList", handle.GetPatientList)
}
