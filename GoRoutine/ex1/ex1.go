package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hangul := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hangul {
		//0.3초 마다 생성
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", v)

	}
}

func PrintNumber() {
	for i := 1; i < 5; i++ {
		//0.4 초 마다 생성
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	// Go 루틴 동시 실행
	PrintHangul()
	PrintNumber()

	//Go 루틴이 완료 될때 까지 3초 대기 시간
	time.Sleep(3 * time.Second)
}
