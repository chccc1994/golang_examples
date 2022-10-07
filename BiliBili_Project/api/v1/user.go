package v1

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chccc1994/bilibili/models"
	"github.com/chccc1994/bilibili/pkg/e"
	"github.com/chccc1994/bilibili/service"
	"github.com/chccc1994/bilibili/utils"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	code := e.SUCCESS
	// 表单参数
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		code = e.InvaliParams
		c.JSON(http.StatusOK, gin.H{
			"Status": code,
			"Msg":    e.GetMsg(code),
			"Data":   "用户登陆错误",
		})
	}
	var userLoginService service.UserLogin
	_ = c.ShouldBind(&userLoginService)
	res := userLoginService.UserLogin()
	c.JSON(http.StatusOK, gin.H{
		"Status": code,
		"Msg":    e.GetMsg(code),
		"Data":   res,
	})
}

func UserRegister(c *gin.Context) {
	code := e.SUCCESS
	// 表单参数，参数校验
	user_name := c.PostForm("user_name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	userCode := c.PostForm("code")

	if email == "" || userCode == "" ||
		user_name == "" || password == "" {
		code = e.InvaliParams
		c.JSON(http.StatusOK, gin.H{
			"Status": code,
			"Msg":    e.GetMsg(code),
		})
		return
	}
	// 验证码
	sysRDBCode, err := models.InitRedisDB().Get(c, email).Result()
	if err != nil {
		code = e.InvaliParams
		c.JSON(http.StatusOK, gin.H{
			"Status": code,
			"Msg":    e.GetMsg(code),
		})
		return
	}
	if sysRDBCode != userCode {
		code = e.InvaliParams
		c.JSON(http.StatusOK, gin.H{
			"Status": code,
			"Msg":    e.GetMsg(code),
		})
		return
	}

	var userRegisterService service.UserRegister
	_ = c.ShouldBind(&userRegisterService)
	res := userRegisterService.UserRegister()
	c.JSON(http.StatusOK, gin.H{
		"Status": code,
		"Msg":    e.GetMsg(code),
		"Data":   res,
	})
}

func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	fmt.Println("email:", email)
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Parameter error",
		})
		return
	}
	code := utils.GetRand()
	models.InitRedisDB().Set(c, email, code, time.Second*120)

	err := utils.SendCode(email, code)
	if err != nil {
		log.Println("Send code error", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Send code error" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Sending succeeded",
	})
}
