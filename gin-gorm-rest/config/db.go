package config

import (
	"context"
	"fmt"
	"log"
	"vietanh/gin-gorm-rest/models"

	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	RedisClient  *redis.Client
	RabbitMQConn *amqp091.Connection
	AppConfig    Config
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
		&models.RefreshToken{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.RolePermission{},
		// Add other models here
	}

	for _, model := range modelsToMigrate {
		db.AutoMigrate(model)
	}
	DB = db
	return db
}

// ConnectRabbitMQ thiết lập kết nối RabbitMQ
func ConnectRabbitMQ(config *Config) *amqp091.Connection {
	// Chuỗi kết nối RabbitMQ
	RabbitMQConnStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitMqUser, config.RabbitMQPassword, config.RabbitMQHost, config.RabbitMQPort)
	// Kết nối RabbitMQ
	conn, err := amqp091.Dial(RabbitMQConnStr)
	if err != nil {
		log.Fatalf("Failed to connect rabbitMQ: %v", err)
	}
	RabbitMQConn = conn
	return conn
}

// ConnectRedis thiết lập kết nối Redis
func ConnectRedis(config *Config) *redis.Client {
	redisAddr := fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   config.RedisDB,
	})

	// Kiểm tra kết nối Redis
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	RedisClient = client
	return client
}

// InitConfig khởi tạo cả PostgreSQL và Redis
func InitConfig(config *Config) {
	Connect(config)
	ConnectRedis(config)
}
