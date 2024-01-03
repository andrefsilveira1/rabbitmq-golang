package main

import (
	"log"
	"rabbitmq/utils"

	"github.com/streadway/amqp"
)

func main() {

	connect, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.HandleError(err, "Failed to connect to RabbitMQ")
	defer connect.Close()

	channel, err := connect.Channel()
	utils.HandleError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"Main queue", // Queue Name
		false,        // Durable
		false,        // Delete when unused
		false,        // Exclusive
		false,        // No-wait
		nil,          // Arguments
	)
	utils.HandleError(err, "Failed to declare a queue")

	// Consumir mensagens da fila
	msgs, err := channel.Consume(
		queue.Name, // Queue
		"",         // Consumer
		true,       // Auto-ack
		false,      // Exclusive
		false,      // No-local
		false,      // No-wait
		nil,        // Args
	)
	utils.HandleError(err, "Failed to register a consumer")

	// Loop para consumir mensagens
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("Listening...")
	<-forever
}
