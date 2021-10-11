package main

import "fmt"

//  if else

func canIDrink(age int) bool {
	if KoreanAge := age + 2; KoreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(13))
}
