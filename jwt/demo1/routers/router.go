package routers

import (
	"jwt_demo1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "ping测试",
		})
	})
	// jwt
	r.POST("/auth", services.AuthHandle)
	r.GET("/user", services.UserHandle)
	r.Run(":9090")
}
