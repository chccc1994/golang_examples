package main

import (
	"fmt"

	"github.com/chccc1994/bilibili/models"
	"github.com/chccc1994/bilibili/routers"
)

func main() {
	fmt.Println("哔哩哔哩项目开始")
	models.InitMySQLDb()
	// models.InitRedisDB()
	routers.InitRouter()
}
