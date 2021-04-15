package broker

import (
	"fmt"

	"github.com/streadway/amqp"
)

type brokerRabbitMq struct{}

func CreateChannel(queueName string) *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return ch
}
