package main

import "fmt"

//swich

func main() {
	fmt.Println(yourAge(10))
}

func yourAge(age int) bool {
	switch KoreanAge := age + 2; KoreanAge {
	case 5:
		return true
	case 0:
		return false
	}
	return false
}
