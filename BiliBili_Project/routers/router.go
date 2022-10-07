package routers

import (
	"github.com/chccc1994/bilibili/api"
	"github.com/chccc1994/bilibili/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/ping", api.Ping)

	r.Run(utils.HttpPort)
}
