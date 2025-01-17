package config

import (
	"fmt"
	"log"
	"vietanh/gin-gorm-rest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	AppConfig Config
)

func Connect(config *Config) *gorm.DB {

	// Tạo chuỗi kết nối PostgreSQL
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.DBHost,
		config.DBPort,
		config.PostgresDB,
	)

	// Mở kết nối tới PostgreSQL
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
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
	return db
}
