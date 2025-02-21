package service

type RabbitMQService interface {
	SendEvent(RoutingKey string, ExchangeName string, ExchangeType string, message string) error
	ConsumeEvent(exchangeName, exchangeType, routingKey, queueName string, handler func(string)) error
}
