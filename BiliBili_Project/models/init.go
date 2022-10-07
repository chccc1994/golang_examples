package models

import (
	"fmt"
	"log"
	"time"

	"github.com/chccc1994/bilibili/utils"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库mysql redis
var Db *gorm.DB

var err error

func InitMySQLDb() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Panicln("Database Connect Error")
		return
	}
	// Db.AutoMigrate(&User{}, &Admin{})
	// 多张表迁移
	migration()
	sqlDb, _ := Db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(10 * time.Second)
}

func InitRedisDb() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
