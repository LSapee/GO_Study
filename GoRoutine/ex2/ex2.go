package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done() //작업 완료를 표시
}

func main() {
	wg.Add(10) // 작업 갯수 10
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 100)
	}
	wg.Wait() //작업이 완료되기를 기다림
	fmt.Println("계산 완료")
}
