package something

import "fmt"

// 소문자의 경우 프라이빗
func sayHello() {
	fmt.Println("hello world")
}

// 대문자의 경우 퍼블릭
func SayBye() {
	sayHello()
	fmt.Println("bye")
}
