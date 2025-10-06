package models

import (
	"flashchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name     string
	Password string
	// Phone         string `valid: "matches(^1[3-9]{1}\\d9$)"`
	Email         string `valid: "email"`
	Identity      string
	ClientIP      string
	ClientPort    string
	Salt          string
	LastLogin     time.Time
	HeartBeatTime time.Time
	LogOutTime    time.Time `gorm:"column:log_out_time" json"column:log_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUsers() []*UserBasic {
	users := make([]*UserBasic, 10)
	utils.DB.Find(&users)

	return users
}

func FindUserByName(name string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("name = ", name).First(&user)
}

func FindUserByNameAndPwd(name, password string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("name = ? and password = ?", name, password).First(&user)
}

func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ", email).First(&user)
}

func CreateUsers(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{
		Name:     user.Name,
		Password: user.Password,
	})
}
