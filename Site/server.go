package main

import (
	"fmt"
	"log"

	"github.com/elolpuer/FreelanceParcer/Site/cfg"
	"github.com/elolpuer/FreelanceParcer/Site/pkg/controller"
	"github.com/gin-gonic/gin"
)

//Config  ...
var config cfg.Cfg

func main() {
	router := gin.Default()

	router.GET("/", controller.IndexGet)
	router.GET("/data", controller.IndexPost)

	err := router.Run(fmt.Sprintf("%s:%s", config.HOST, config.PORT))
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	config = cfg.Get()
}
