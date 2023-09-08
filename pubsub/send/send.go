package main

import (
	"context"
	"log"
	"os"
	"rabbitmq/utils"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// define a echanger
	err = ch.ExchangeDeclare(
		"logs",
		"fanout", // deliver to all queues
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "exchange declare error")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	q, err := ch.QueueDeclare(
		"",    // name  ------ get a random name
		false, // durable
		false, // delete when unused
		true,  // exclusive  as connection closes, the queue is deleted
		false, // no-wait
		nil,   // arguments
	)
	// bind queue to exchanger
	err = ch.QueueBind(
		q.Name,
		"", // routing key
		"logs",
		false,
		nil,
	)
	body := utils.BodyFrom(os.Args)

	err = ch.PublishWithContext(
		ctx,
		"logs",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	utils.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)

}
