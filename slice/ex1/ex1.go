package main

import "fmt"

func main() {
	var slice []int
	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}
	// slice 길이가 0인 배열에서 2번째 요소를 변경하려 하니 에러 발생

	slice[1] = 10
	fmt.Println(slice)
}
