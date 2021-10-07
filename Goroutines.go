package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	//채널 만들어서 소통하기
// 	// c := make(chan string)
// 	// people := [2]string{"jk", "je"}
// 	// for _, person := range people {
// 	// 	go sexyCount(person, c)
// 	// }

// 	// resultOne := <-c
// 	// resultTwo := <-c

// 	// fmt.Println("Waiting for messages")
// 	// fmt.Println("Received this message: ", resultOne)
// 	// fmt.Println("Received this message: ", resultTwo)
// 	c := make(chan string)
// 	people := [5]string{"jk", "je", "jl", "jm", "ju"}
// 	for _, person := range people {
// 		go sexyCount(person, c)
// 	}

// 	for i := 0; i < len(people); i++ {
// 		fmt.Println(<-c)
// 	}
// }

// func sexyCount(person string, c chan string) {
// 	time.Sleep(time.Second * 10)
// 	c <- person + " is sexy"
// }
