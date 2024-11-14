package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email 	 string `gorm:"Unique"`
	Password string
}
