package service

import (
	"github.com/chccc1994/bilibili/models"
	"github.com/chccc1994/bilibili/pkg"
	"github.com/chccc1994/bilibili/pkg/e"
	"github.com/chccc1994/bilibili/serializer"
	"github.com/chccc1994/bilibili/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserRegister struct {
	UserName string `json:"user_name" form:"user_name" bind:"required"`
	Email    string `json:"email" form:"email" bind:"required"`
	Password string `json:"password" form:"password" bind:"required"`
	Code     string `json:"code" form:"code"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" bind:"required"`
	Password string `json:"password" form:"password" bind:"required"`
}
type UserUpdate struct {
	Name     string `form:"name" json:"name" bind:"required"`
	Gender   int    `form:"gender" json:"gender" `
	Birthday string `form:"birthday" json:"birthday" time_format:"2006-01-02"`
	Sign     string `form:"sign" json:"sign" `
}
type UserInfo struct {
}
type UserSearch struct {
	UserName string `form:"user_name" json:"user_name"`
}

func (service *UserRegister) UserRegister() serializer.Response {
	var user models.User
	var count int64

	code := e.SUCCESS
	// 验证邮箱格式是否正确
	if !pkg.VerifyEmailFormat(service.Email) {
		code = e.InvaliParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "邮箱格式错误",
		}
	}
	// 验证邮箱是否已经注册
	models.Db.Model(&models.User{}).Where("email=?", service.Email).Count(&count)
	if count > 0 {
		code = e.InvaliParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "邮箱已注册",
		}
	}
	// 密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(service.Password), bcrypt.DefaultCost)
	user = models.User{
		Email:    service.Email,
		UserName: service.UserName,
		Password: string(hashedPassword),
	}
	// 发送验证码

	// 验证验证码
	// 创建用户
	models.Db.Create(&user)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "创建用户成功",
	}

	// 注册成功
}

func (service *UserLogin) UserLogin() serializer.Response {
	code := e.SUCCESS
	var user models.User
	if !pkg.VerifyEmailFormat(service.Email) {
		code = e.InvaliParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "不是正确的邮箱格式",
		}
	}
	models.Db.Model(&models.User{}).Where("email = ?", service.Email).First(&user)
	if user.ID == 0 {
		code = e.InvaliParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在",
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(service.Password)); err != nil {
		code = e.InvaliParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码不正确",
		}
	}
	user.Authority = 0
	// 生成Token
	token, _ := jwt.GenToken(user)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}
}
func (service *UserUpdate) UserUpdate(id uint) serializer.Response {
	code := e.SUCCESS
	err := models.Db.Model(models.User{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"user_name": service.Name, "gender": service.Gender,
			"birthday": service.Birthday, "sign": service.Sign}).Error
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "修改信息失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "修改信息成功",
	}
}
func (service *UserInfo) UserShow(id uint) serializer.Response {
	code := e.SUCCESS
	var user models.User
	models.Db.Model(&models.User{}).First(&user.ID)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}

}

func (service *UserSearch) UserSearch() serializer.Response {
	code := e.SUCCESS
	var user []models.User
	models.Db.Model(&models.User{}).
		Where("user_name LIKE ?", "%"+service.UserName+"%").
		Find(&user)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUsers(user),
	}
}
