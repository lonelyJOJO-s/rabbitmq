package main

import (
	"context"
	"fmt"
	"os"
	"rabbitmq/utils"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// we have one exchanger and push msg to consumer

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"msg",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "declare exchange error")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := utils.BodyFrom(os.Args)

	err = ch.PublishWithContext(
		ctx,
		"msg",
		severityFrom(os.Args), // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[2]
		fmt.Println(s)
	}
	return s
}
