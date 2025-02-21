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

func (r *RabbitMQServiceImpl) SendEvent(RoutingKey string, ExchangeName string, ExchangeType string, message string) error {
	ch, err := r.RabbitMQConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// declare Exchange
	err = ch.ExchangeDeclare(
		ExchangeName, // Tên Exchange
		ExchangeType, // Loại Exchange (direct, topic, fanout, headers)
		true,         // Durable (bền vững, không mất khi restart)
		false,        // Auto-delete (không tự động xoá khi không có queue nào sử dụng)
		false,        // Internal
		false,        // No-wait
		nil,          // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declared a exchange: %w", err)
	}

	// Gửi message đến Exchange thay vì trực tiếp đến Queue
	err = ch.Publish(
		ExchangeName, // Gửi đến Exchange
		RoutingKey,   // Routing Key để xác định Queue
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message %w", err)
	}
	log.Printf("send event %s", message)
	return nil
}

func (r *RabbitMQServiceImpl) ConsumeEvent(exchangeName, exchangeType, routingKey, queueName string, handler func(string)) error {
	ch, err := r.RabbitMQConn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}

	// Declare the exchange
	err = ch.ExchangeDeclare(
		exchangeName, // Exchange name
		exchangeType, // Exchange type (direct, topic, fanout, headers)
		true,         // Durable (persistent across restarts)
		false,        // Auto-delete when no queues are bound
		false,        // Internal
		false,        // No-wait
		nil,          // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Declare the queue (it will be bound to the exchange)
	q, err := ch.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Bind queue to exchange with routing key
	err = ch.QueueBind(
		q.Name,       // Queue name
		routingKey,   // Routing key
		exchangeName, // Exchange name
		false,        // No-wait
		nil,          // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	// Register consumer
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer name (empty for auto-generated)
		false,  // Auto-ack (set to false if manual acknowledgment is needed)
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	// Process messages
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
