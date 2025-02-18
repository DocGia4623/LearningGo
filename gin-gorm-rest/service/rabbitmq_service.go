package service

type RabbitMQService interface {
	SendEvent(queueName string, message string) error
	ConsumeEvent(queueName string, handler func(string)) error
}
