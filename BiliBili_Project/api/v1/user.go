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
	"github.com/chccc1994/bilibili/utils/jwt"
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

func UserUpdate(c *gin.Context) {
	// 用户更新

	// 检查当前用户是否存在 email
	user_name := c.PostForm("user_name")
	if user_name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Parameter error",
		})
		return
	}
	// 更新
	var userUpdateService service.UserUpdate
	_ = c.ShouldBind(&userUpdateService)
	_, chaim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	res := userUpdateService.UserUpdate(chaim.UserId)
	c.JSON(http.StatusOK, gin.H{
		"Status": e.SUCCESS,
		"Msg":    e.GetMsg(e.SUCCESS),
		"Data":   res,
	})
}
func UserInfo(c *gin.Context) {
	var userInfoService service.UserInfo
	_ = c.ShouldBind(&userInfoService)
	_, chaim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	fmt.Println("Ok")
	res := userInfoService.UserShow(chaim.UserId)
	c.JSON(http.StatusOK, gin.H{
		"Status": e.SUCCESS,
		"Msg":    e.GetMsg(e.SUCCESS),
		"Data":   res,
	})
}
func UserSearch(c *gin.Context) {
	var userSearchService service.UserSearch
	_ = c.ShouldBind(&userSearchService)
	res := userSearchService.UserSearch()
	c.JSON(http.StatusOK, gin.H{
		"Status": e.SUCCESS,
		"Msg":    e.GetMsg(e.SUCCESS),
		"Data":   res,
	})
}
