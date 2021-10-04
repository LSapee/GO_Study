package main

import (
	"fmt"

	/*
		외부 저장소 가져올 때는 go mod tidy를 써서 필요한 외부 저장소를 받아준다
		1.import 에 외부 저장소 주소를 적고
		2.터미널에 go mod tidy
		3. 확실하게 인포트 되면 import주소의 빨간줄이 없어진다.
	*/
	"github.com/Lsapee/GO_Study/banking"
)

func main() {
	account := banking.Account{Owner: "jk", Balance: 20000}
	fmt.Println(account)
}
