package cfg

import (
	"os"
)

//Cfg struct with env variables
type Cfg struct {
	RabbitMQ string
	PORT     string
	HOST     string
}

//Get return struct config with env variables
func Get() Cfg {
	// var err error
	// err = godotenv.Load("./.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env", err.Error())
	// }
	rabbit, _ := os.LookupEnv("RabbitMQ")
	port, _ := os.LookupEnv("PORT")
	host, _ := os.LookupEnv("HOST")

	return Cfg{
		RabbitMQ: rabbit,
		PORT:     port,
		HOST:     host,
	}
}
