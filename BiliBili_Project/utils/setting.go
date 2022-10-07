package utils

import (
	"gopkg.in/ini.v1"
)

var (
	// 服务
	AppMode   string
	HttpPort  string
	JwtSecret string
	// 数据库
	DbUser        string
	DbPassWord    string
	DbHost        string
	DbPort        string
	DbName        string
	DbDefaultPage string
	DbDefaultSize string
	// SMTP邮箱授权码
	EmailPassword string
	// redis
	RedisAddress  string
	RedisPassword string
	RedisDb       int
	// qiniu

	// 登陆账号

)

func init() {
	file, err := ini.Load("conf/conf.ini")
	if err != nil {
		panic(err)
		return
	}
	LoadServer(file)
	LoadMysqlDb(file)
	LoadEmailStmp(file)
	LoadRedisDb(file)
	// LoadQiniu(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":9090")
	JwtSecret = file.Section("server").Key("JwtSecret").MustString("something_very_secret")
}
func LoadMysqlDb(file *ini.File) {
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbName = file.Section("database").Key("DbName").MustString("golangdb")
	DbDefaultPage = file.Section("database").Key("DbDefaultPage").MustString("10")
	DbDefaultSize = file.Section("database").Key("DbDefaultSize").MustString("1")
}
func LoadEmailStmp(file *ini.File) {
	EmailPassword = file.Section("email").Key("EmailPassword").MustString("") // SMTP授权码
}

func LoadRedisDb(file *ini.File) {
	RedisAddress = file.Section("redis").Key("RedisAddress").MustString("127.0.0.1:6379")
	RedisPassword = file.Section("redis").Key("RedisPassword").MustString("")
	RedisDb = file.Section("redis").Key("RedisDb").MustInt(0)
}

// func LoadQiniu(file *ini.File) {

// }
