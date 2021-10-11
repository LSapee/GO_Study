package main

import "fmt"

//메모리 관리 (포인터)

func main() {

	a := 2
	b := a // b = 2
	c := &a
	a = 10
	// 메모리 주소 확인할때는 변,상수 앞에 &붙이면 된다.
	// fmt.Println(&a, &b)
	/*
		*를 붙이면 메모리의 값을 나타냄 이경우 c는 a의 주소를 나타내고 있기에
		a의 값을 변경하면 c의 값도 변경된다.
	*/
	//fmt.Println(*c)
	fmt.Println(a, b, *c)
}
