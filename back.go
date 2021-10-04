package main

import (
	"fmt"

	"github.com/LSapee/GO_Study/tree/master/banking"
)

func main() {
	account := banking.Account{Owner: "jk", Balance: 20}
	fmt.Println(account)
}
