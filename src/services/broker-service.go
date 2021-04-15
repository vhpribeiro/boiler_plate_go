package services

import "github.com/streadway/amqp"

type IBrokerService interface {
	SendMessageToBroker(queueName string) error
}

type brokerService struct {
	Channel *amqp.Channel
}

func (b *brokerService) SendMessageToBroker(queueName string) error {
	err := b.Channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World!"),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func NewBrokerService(channel *amqp.Channel) IBrokerService {
	return &brokerService{
		Channel: channel,
	}
}
