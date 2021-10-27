package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d", a, b, a/b)
}

func main() {
	divide(3, 4)
	divide(3, 0)
}
