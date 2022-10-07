package routers

import (
	"github.com/chccc1994/bilibili/api"
	v1 "github.com/chccc1994/bilibili/api/v1"
	"github.com/chccc1994/bilibili/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/ping", api.Ping)
	r.POST("/send-code", v1.SendCode)
	auth := r.Group("/api/v1")
	{
		// 用户操作
		auth.POST("/user/register", v1.UserRegister)
		auth.POST("/user/login", v1.UserLogin)

		authed := auth.Group("/")
		{
			//用户操作
			authed.PUT("user/update", v1.UserUpdate)
			authed.GET("user/show", v1.UserInfo)
			authed.POST("user/search", v1.UserSearch)
		}
		//
	}
	r.Run(utils.HttpPort)
}
