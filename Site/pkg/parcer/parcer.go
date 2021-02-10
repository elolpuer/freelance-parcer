package parcer

import (
	"github.com/elolpuer/FreelanceParcer/Site/cfg"
	"github.com/elolpuer/FreelanceParcer/Site/pkg/models"
	"github.com/streadway/amqp"
)

var config cfg.Cfg

//Get запрашивает данные парсера с сервера
func Get() ([]models.A, error) {
	conn, err := amqp.Dial(config.RabbitMQ)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	err = ch.Publish(
		"",
		"putSite",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: string(2),
			ReplyTo:       q.Name,
			Body:          []byte("/startSite"),
		},
	)
	var str string

	for d := range msgs {
		if d.CorrelationId == string(2) {
			str += string(d.Body)
			break
		}
	}
	//go func() {

	//}()
	//defer fmt.Println("Here", str)
	var arr []string
	var brr []models.A
	for i, a := range str {
		if string(a) == "~" {
			var arrOne string
			s := str[i+1:]
			for _, b := range s {
				if string(b) == "\n" {
					break
				}
				arrOne += string(b)
			}
			arr = append(arr, arrOne)
		}
		if string(a) == "|" {
			var arrOne string
			s := str[i+1:]
			for _, b := range s {
				if string(b) == "\n" {
					break
				}
				arrOne += string(b)
			}
			arr = append(arr, arrOne)
		} else {
			continue
		}
	}
	for i := 0; i < len(arr); i += 2 {
		a := models.A{
			Href: arr[i+1],
			Text: arr[i],
		}
		brr = append(brr, a)
	}
	return brr, nil
}

func init() {
	config = cfg.Get()
}
