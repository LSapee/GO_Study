package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// )

// type extractedJob struct {
// 	id       string
// 	title    string
// 	location string
// 	salary   string
// 	summary  string
// }

// //기본 URL
// var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

// //실행 함수
// func main() {
// 	var jobs []extractedJob
// 	c := make(chan []extractedJob)
// 	totalPages := getPages()

// 	for i := 0; i < totalPages; i++ {
// 		go getPage(i, c)
// 	}

// 	for i := 0; i < totalPages; i++ {
// 		extractedJobs := <-c
// 		jobs = append(jobs, extractedJobs...)
// 	}

// 	writeJobs(jobs)
// 	fmt.Println("Done, extraced", len(jobs))
// }

// func writeJobs(jobs []extractedJob) {
// 	file, err := os.Create("jobs.csv")
// 	checkErr(err)

// 	w := csv.NewWriter(file)
// 	defer w.Flush()

// 	headers := []string{"ID", "Title", "Location", "salary", "Summary"}

// 	wErr := w.Write(headers)
// 	checkErr(wErr)

// 	for _, job := range jobs {
// 		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
// 		jwErr := w.Write(jobSlice)
// 		checkErr(jwErr)
// 	}
// }

// func getPage(page int, mainC chan<- []extractedJob) {
// 	c := make(chan extractedJob)
// 	var jobs []extractedJob
// 	// page당 url 변경 되는것
// 	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
// 	// 처음 페이지 URL 출력
// 	fmt.Println("Requesting", pageURL)
// 	// 페이지 변수 넣기
// 	res, err := http.Get(pageURL)
// 	// 에러 발생시 에러 소환 입력값 err
// 	checkErr(err)
// 	// 정상 작동시 checkCode 작동 입력값 res
// 	checkCode(res)
// 	// 에러 발생해도 나중에 함수를 끝내기 위해 close 마지막에 실행
// 	defer res.Body.Close()

// 	doc, err := goquery.NewDocumentFromReader(res.Body)
// 	checkErr(err)
// 	//탐색할 html element  선택자는 css 선택자로
// 	searchCards := doc.Find(".tapItem")
// 	//탐색하기
// 	searchCards.Each(func(i int, card *goquery.Selection) {
// 		go extractJob(card, c)
// 	})
// 	for i := 0; i < searchCards.Length(); i++ {
// 		job := <-c
// 		jobs = append(jobs, job)
// 	}
// 	mainC <- jobs
// }

// func extractJob(card *goquery.Selection, c chan<- extractedJob) {
// 	id, _ := card.Attr("id")
// 	id = strings.TrimLeft(id, "job_")
// 	title := cleanString(card.Find("h2>span").Text())
// 	location := cleanString(card.Find("div pre").Text())
// 	salary := cleanString(card.Find(".salary-snippet").Text())
// 	summary := cleanString(card.Find(".job-snippet").Text())
// 	c <- extractedJob{
// 		id:       id,
// 		title:    title,
// 		location: location,
// 		salary:   salary,
// 		summary:  summary}
// }

// //페이지 가져오기 URL
// func getPages() int {
// 	pages := 0
// 	res, err := http.Get(baseURL)
// 	checkErr(err)
// 	checkCode(res)

// 	defer res.Body.Close()

// 	doc, err := goquery.NewDocumentFromReader(res.Body)
// 	checkErr(err)
// 	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
// 		pages = s.Find("a").Length()

// 		// fmt.Println(s.Find("a").Length())
// 	})
// 	return pages
// }

// func checkErr(err error) {
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func checkCode(res *http.Response) {
// 	if res.StatusCode != 200 {
// 		log.Fatalln("Request failed with Status: ", res.StatusCode)
// 	}
// }

// func cleanString(str string) string {
// 	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
// }
