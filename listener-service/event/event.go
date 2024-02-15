package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name of the exchange
		"topic",      // type of exchange
		true,         // durable (exchanges survive broker restart)
		false,        // auto-deleted (delete when no consumers)
		false,        // internal (used by other exchanges)
		false,        // no-wait (do not wait for a confirmation from the server)
		nil,          // arguments
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name of the queue
		false, // durable (the queue will survive a broker restart)
		false, // delete when unused
		true,  // exclusive (used by only one connection and the queue will be deleted when that connection closes)
		false, // no-wait (do not wait for a confirmation from the server)
		nil,   // arguments

	)
}
