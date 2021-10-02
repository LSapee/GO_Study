package something

import "fmt"

// 소문자의 경우 프라이빗
// 프라이빗의 경우 외부 패키지에서 접촉 불가
func sayHello() {
	fmt.Println("hello world")
}

// 대문자의 경우 퍼블릭
func SayBye() {
	sayHello()
	fmt.Println("bye")
}
