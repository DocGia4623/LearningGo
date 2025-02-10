package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint     `json:"ID"`
	UserName string   `json:"username"`
	FullName string   `json:"fullname"`
	Role     string   `json:"role"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Devices  []Device `gorm:"foreignKey:UserID"`
}
