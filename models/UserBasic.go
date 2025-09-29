package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LastLogin     uint64
	HeartBeatTime uint64
	LogOut        uint64
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
