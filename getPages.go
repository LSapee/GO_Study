package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// type extractedJob struct {
// 	id       string
// 	title    string
// 	location string
// 	salary   string
// 	summary  string
// }

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	totalPages := getPages()
	// fmt.Println(totalPages)
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	// fmt.Println("Requesting", pageURL
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")

	searchCards.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("id")
		fmt.Println(id)
	})
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()

		// fmt.Println(s.Find("a").Length())
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

// func cleanString(str string) string {

// }
