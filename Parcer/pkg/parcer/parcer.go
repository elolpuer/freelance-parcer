package parcer

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//Parc парсит с фриланс сайта заказы
func Parc() (string, error) {
	// Request the HTML page.
	res, err := http.Get("https://www.fl.ru/projects")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	var str string
	// Find the review items
	doc.Find("div#projects-list div.b-post").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find("h2 a").Text()
		href, _ := s.Find("h2 a").Attr("href")
		if i > 9 {
			str += fmt.Sprintf("")
		} else {
			str += fmt.Sprintf("~%s\n - |%s\n", title, "https://www.fl.ru"+href)
		}
	})
	return str, nil
}
