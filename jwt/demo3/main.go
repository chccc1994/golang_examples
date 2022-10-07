package main

import (
	"github.com/chccc1994/gin_jwt/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", api.Ping)
	r.Run(":9090")
}
