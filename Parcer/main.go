package main

import (
	"log"

	"github.com/elolpuer/FreelanceParcer/Parcer/cfg"
	"github.com/elolpuer/FreelanceParcer/Parcer/pkg/parcer"
	"github.com/streadway/amqp"
)

var config cfg.Cfg

func main() {
	conn, err := amqp.Dial(config.RabbitMQ)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"putTelegram", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	msgsTelegram, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}
	q2, err := ch.QueueDeclare(
		"putSite", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	msgsSite, err := ch.Consume(
		q2.Name, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)
	var text string

	go func() {
		for d := range msgsTelegram {
			text, err = parcer.Parc()
			if err != nil {
				log.Fatal(err)
			}
			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(text),
				})
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	go func() {
		for d := range msgsSite {
			text, err = parcer.Parc()
			if err != nil {
				log.Fatal(err)
			}
			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(text),
				})
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func init() {
	config = cfg.Get()
}
