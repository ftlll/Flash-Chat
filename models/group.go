package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name        string
	AdminId     uint
	Avator      string
	Description string
}

func (table *Group) TableName() string {
	return "group"
}
