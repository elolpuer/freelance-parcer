package cfg

import (
	"os"
)

//Cfg struct with env variables
type Cfg struct {
	RabbitMQ string
}

//Get return struct config with env variables
func Get() Cfg {
	// var err error
	// err = godotenv.Load("./.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env", err.Error())
	// }
	rabbit, _ := os.LookupEnv("RabbitMQ")
	return Cfg{
		RabbitMQ: rabbit,
	}
}
