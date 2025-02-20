package service

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQServiceImpl struct {
	RabbitMQConn *amqp091.Connection
}

// NewRabbitMQServiceImpl là factory function để tạo đối tượng RabbitMQServiceImpl
func NewRabbitMQServiceImpl(rabbitMQConn *amqp091.Connection) RabbitMQService {
	return &RabbitMQServiceImpl{
		RabbitMQConn: rabbitMQConn,
	}
}

func (r *RabbitMQServiceImpl) SendEvent(queueName string, message string) error {
	ch, err := r.RabbitMQConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// declare queue

	_, err = ch.QueueDeclare(
		queueName, // Tên hàng đợi
		true,      // Bảo vệ không bị mất nếu RabbitMQ restart
		false,     // Không tự động xóa khi không còn consumer
		false,     // Không chia sẻ với các consumer khác
		false,     // Không tạo hàng đợi bền vững
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declared a queue: %w", err)
	}

	// send
	err = ch.Publish(
		"",
		queueName, // Tên hàng đợi
		false,     // Không yêu cầu acknowledgement
		false,     // Không yêu cầu routing key
		amqp091.Publishing{
			ContentType: "application.json",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message %w", err)
	}
	log.Printf("send event %s", message)
	return nil
}

func (r *RabbitMQServiceImpl) ConsumeEvent(queueName string, handler func(string)) error {
	ch, err := r.RabbitMQConn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}

	// Khai báo hàng đợi (chỉ cần tạo ở phía consumer)
	_, err = ch.QueueDeclare(
		queueName,
		true,  // Durable (bền vững)
		false, // Không tự động xóa khi không có consumer
		false, // Không exclusive (cho phép nhiều consumer)
		false, // Không cần tạo hàng đợi bền vững
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Đăng ký consumer
	msgs, err := ch.Consume(
		queueName,
		"",
		true,  // Auto-ack (nếu cần kiểm soát ACK, nên để false)
		false, // Không độc quyền (exclusive)
		false, // Không cần requeue nếu consumer bị mất
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	// Lắng nghe tin nhắn
	go func() {
		for msg := range msgs {
			log.Printf("📨 Received message: %s", msg.Body)
			if err := handleMessage(handler, string(msg.Body)); err != nil {
				log.Printf("⚠️ Error handling message: %v", err)
			}
		}
	}()
	return nil
}

// Xử lý lỗi khi gọi handler
func handleMessage(handler func(string), message string) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("❌ Panic recovered in handler: %v", r)
		}
	}()
	handler(message)
	return nil
}
