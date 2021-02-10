package controller

import (
	"log"

	"github.com/elolpuer/FreelanceParcer/Site/pkg/models"
	"github.com/elolpuer/FreelanceParcer/Site/pkg/parcer"
	temp "github.com/elolpuer/FreelanceParcer/Site/pkg/tml"
	"github.com/gin-gonic/gin"
)

var tml = temp.GetTemplates()

//IndexGet ...
func IndexGet(ctx *gin.Context) {
	tml.ExecuteTemplate(ctx.Writer, "index.gohtml", struct {
		Title string
		H1    string
	}{
		Title: "Index Page",
		H1:    "Index Page",
	})
}

//IndexPost возвращает данные от парсера
func IndexPost(ctx *gin.Context) {
	data, err := parcer.Get()
	if err != nil {
		log.Fatal(err)
	}
	tml.ExecuteTemplate(ctx.Writer, "indexPost.gohtml", struct {
		Title string
		H1    string
		Data  []models.A
	}{
		Title: "Index Page",
		H1:    "Data",
		Data:  data,
	})
}
