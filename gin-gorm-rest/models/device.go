package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	ID       uint   `json:"ID"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	UserID   uint   `json:"user_id"`
	User     User   `gorm:"foreignKey:UserID"`
}
