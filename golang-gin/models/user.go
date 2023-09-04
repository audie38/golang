package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Username string `gorm:"size:200;not null;" json:"username" binding:"required"`
	Password string `gorm:"size:1000;not null;" json:"password" binding:"required"`
}