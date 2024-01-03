package main

import (
	"log"
	"rabbitmq/utils"

	"github.com/streadway/amqp"
)

func main() {
	connect, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.HandleError(err, "Connection failed")
	defer connect.Close()

	channel, err := connect.Channel()
	utils.HandleError(err, "Channel creation has failed")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"Main queue", // Queue Name
		false,        // Duralbe
		false,        // Delete when unused
		false,        // Exclusive
		false,        // No wait
		nil,          // Arguments
	)

	utils.HandleError(err, "Failed to create a queue")

	body := "Hi, there!"

	err = channel.Publish(
		"",         // Exchange
		queue.Name, // Routing key
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.HandleError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
