package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint     `json:"ID"`
	UserName string   `gorm:"unique" json:"username"`
	FullName string   `json:"fullname"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Devices  []Device `gorm:"foreignKey:UserID"`
	Roles    []Role   `gorm:"many2many:user_roles;" json:"roles"`
}

type Role struct {
	gorm.Model
	ID         uint         `json:"ID"`
	Name       string       `gorm:"unique;not null"`
	Permission []Permission `gorm:"many2many:role_permissions"`
}

type UserRole struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`
}

type Permission struct {
	gorm.Model
	Name string `gorm:"unique;not null" json:"name"`
}

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}
