package event

import (
	ampq "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *ampq.Channel) error {
	return ch.ExchangeDeclare("logs_topic", "topic", true, false, false, false, nil)
}

func declareRandomQueue(ch *ampq.Channel) (ampq.Queue, error) {
	return ch.QueueDeclare("", false, false, true, false, nil)
}
