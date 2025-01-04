package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;unique;not null"`
	Password string `gorm:"size:255;not null"`
	Role     string `gorm:"size:50;not null" default:"user"`
}

func (u *User) TableName() string {
	return "user"
}
