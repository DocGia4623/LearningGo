package config

import (
	"vietanh/gin-gorm-rest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Mở kết nối tới PostgreSQL
	db, err := gorm.Open(postgres.Open("postgres://postgres:12345@database:5433/postgres"), &gorm.Config{})
	if err != nil {
		panic(err) // Báo lỗi nếu không kết nối được
	}
	// Tự động migrate tất cả các bảng
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Device{},
		// Add other models here
	}

	for _, model := range modelsToMigrate {
		db.AutoMigrate(model)
	}
	DB = db
}
