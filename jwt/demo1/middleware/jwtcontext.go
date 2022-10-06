package middleware

import (
	"jwt_demo1/modles"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 基于JWT认证的中间件   验证token的中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//携带Token有三种方式
		//1.放在请求头
		//2.放在请求体
		//3.放在URI
		//这里实现的方法是Token放在header的Authorization，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization") //获取请求中头部的token
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "header中auth为空",
			})
			c.Abort() //授权失败，调用Abort以确保没有调用此请求的其余处理程序
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "header中auth格式有误",
			})
			c.Abort()
			return
		}

		myclaim, err := modles.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		//将当前请求的username信息保存到请求的上下文c
		//c.Set("username", myclaim.UserName)
		//c.Next()
		//后续的处理函数可以用c.Get("username")来获取当前请求的用户信息

		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"myClaim": myclaim},
		})
	}
}
