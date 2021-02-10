package main

import (
	"fmt"
	"log"

	"github.com/elolpuer/FreelanceParcer/TelegramBot/cfg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/streadway/amqp"
)

var config cfg.Cfg

func main() {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Fatal(err)
	}
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
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if update.Message.Text == "/start" {
			q, err := ch.QueueDeclare(
				"",    // name
				false, // durable
				false, // delete when unused
				true,  // exclusive
				false, // no-wait
				nil,   // arguments
			)
			if err != nil {
				log.Fatal(err)
			}
			msgs, err := ch.Consume(
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
			err = ch.Publish(
				"",
				"putTelegram",
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: string(1),
					ReplyTo:       q.Name,
					Body:          []byte("/startTelegram"),
				},
			)
			for d := range msgs {
				if d.CorrelationId == string(1) {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s", d.Body))
					bot.Send(msg)
					break
				}
			}
		}
	}
}

func init() {
	config = cfg.Get()
}
