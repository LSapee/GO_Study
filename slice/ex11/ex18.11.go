package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 0)

	idx := 2

	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}
	slice[idx] = 100

	fmt.Println(slice)
	//1,2,3,4,5,6,0
	//i = 5
	//slice 6 > slice 5
	//1,2,3,4,5,6,6
	//1,2,3,4,5,5,6 ...
	//1,2,3,3,4,5,6 이상태에서 slice[idx] 에 100을 삽입
}
