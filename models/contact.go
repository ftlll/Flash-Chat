package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	OwnerId     string
	FromId      string
	Type        int
	Description string
}

func (table *Contact) TableName() string {
	return "contact"
}
