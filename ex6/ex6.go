package main

import "fmt"

func superAdd(numbers ...int) int {
	// 출력을 위한 변수 선언
	total := 0
	//  range값을 nil로 하기위해 _ 로 선언
	for _, number := range numbers {
		total += number
	}
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	return total

	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
}

func main() {
	//total값은
	total := superAdd(1, 2, 3, 4, 5, 6, 6, 55, 6, 7, 7, 84, 545, 54, 6, 4123, 13, 4, 15, 534, 34, 5, 6)
	fmt.Println(total)
}
