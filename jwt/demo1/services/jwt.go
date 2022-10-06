package services

import (
	"jwt_demo1/middleware"
	"jwt_demo1/modles"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 请求处理函数
func AuthHandle(c *gin.Context) {
	var user modles.UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	//校验用户名和密码是否正确
	if user.UserName == "mh" && user.Password == "123456" {
		//生成Token
		Token, _ := modles.GenToken(user.UserName) //获取到秘钥加密的token
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": Token},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return

}

func UserHandle(c *gin.Context) {
	//headerAuth := c.GetHeader("Authorization")
	//token, _ := jwt.ParseToken(headerAuth)
	authMiddleware := middleware.JWTAuthMiddleware()
	authMiddleware(c)
}
