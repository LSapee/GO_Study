package main

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// )

// var errRequestFailed = errors.New("You can't failed")

// func main() {
// 	var results = map[string]string{}
// 	urls := []string{
// 		"https://www.airbnb.com",
// 		"https://www.google.com",
// 		"https://www.amazon.com",
// 		"https://www.reddit.com",
// 		"https://www.google.com",
// 		"https://soundcloud.com",
// 		"https://www.fackbook.com",
// 		"https://www.instagram.com",
// 		"https://academy.nomadcoders.co/",
// 	}

// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "Failed"
// 		}
// 		results[url] = result
// 		fmt.Println(url, result)
// 	}

// }

// func hitURL(url string) error {
// 	fmt.Println("Checking:", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		// err출력해서 무슨 에러인지 확인한다. err< 에러 내용 resp.StatusCode < 에러 번호
// 		fmt.Println(err, resp.StatusCode)
// 		return errRequestFailed
// 	}
// 	return nil
// }
