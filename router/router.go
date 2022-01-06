package router

import (
	"github.com/gin-gonic/gin"
	"serena/middleware"
	"serena/server"
)

/**
 * 路由
 */

func InitRouters() (r *gin.Engine) {

	r = gin.Default()
	r.Use(middleware.Cors())

	//控制台接口（http post请求，用于查看消息队列的基本信息）
	console := r.Group("/Console")
	console.Use(middleware.Safe)
	{
		console.POST("/GetNode", server.GetNode)
	}
	return
}
