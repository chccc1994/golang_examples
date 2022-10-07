package serializer

import (
	"time"

	"github.com/chccc1994/bilibili/models"
)

type User struct {
	ID       uint      `json:"id"`
	Avatar   string    `json:"avatar"`
	UserName string    `json:"user_name"`
	Email    string    `json:"email"`
	Gender   int       `json:"gender"`
	Birthday time.Time `json:"birthday"`
	Sign     string    `json:"sign"`
}

// BuildUser 序列化用户
func BuildUser(user models.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Gender:   user.Gender,
		Birthday: user.Birthday,
		Sign:     user.Sign,
	}
}

func BuildUsers(items []models.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
