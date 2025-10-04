package models

import (
	"flashchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LastLogin     time.Time
	HeartBeatTime time.Time
	LogOutTime    time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUsers() []*UserBasic {
	users := make([]*UserBasic, 10)
	utils.DB.Find(&users)
	// for _, v := range users {
	// 	fmt.Println(v)
	// }
	return users
}

func CreateUsers(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}
