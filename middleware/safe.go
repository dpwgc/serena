package middleware

import (
	"github.com/gin-gonic/gin"
	"serena/config"
)

// Safe 安全验证中间件
func Safe(c *gin.Context) {

	secretKey := c.GetHeader("secretKey")

	if secretKey != config.Get.Registry.SecretKey {
		c.JSON(-1, gin.H{
			"code": -1,
			"msg":  "Secret key matching error",
		})
		c.Abort()
	}
	return
}
