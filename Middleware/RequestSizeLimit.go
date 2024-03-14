package Middleware

import "github.com/gin-gonic/gin"

func SizeLimit(c *gin.Context) {
	if c.Request.ContentLength > 1024*1024*2 {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "请求体过大",
		})
		c.Abort()
	}
	c.Next()
}
