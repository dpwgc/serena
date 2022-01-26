package router

import (
	"github.com/gin-gonic/gin"
	"serena/middleware"
	"serena/server"
)

// InitRouters 初始化路由
func InitRouters() (r *gin.Engine) {

	r = gin.Default()
	r.Use(middleware.Cors())

	console := r.Group("/Registry")
	console.Use(middleware.Safe)
	{
		console.POST("/GetNodes", server.GetNodes)
	}
	return
}
