package main

import (
	"fmt"
	"strings"
)

//func 함수명 (입력값 타입) (리텁값 타입,리턴값 타입)
func lenAndUpper(name string) (lenght int, uppercase string) {
	//defer은 함수가 다 실행된 후에 추가 실행됨
	defer fmt.Println("I do it")
	// 여기서 :=를 하지 않는 이유는 위에 lenght, uppercase는 선언 되어있다.
	// name의 길이
	lenght = len(name)
	//name 을 대문자로
	uppercase = strings.ToUpper(name)
	return

}

func main() {
	fmt.Println("aaaa")
	//lanAdUpper = lenght & up = uppercase
	lanAndUpper, up := lenAndUpper("jk")
	fmt.Println(lanAndUpper, up)
}

/*
결과값
aaaa
i do it
2 JK
i do it이 2 JK 보다 먼저 출력되는 이유는 21라인에서 리턴값을 반환해주고 실행되서

*/
