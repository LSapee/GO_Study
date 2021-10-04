package main

import (
	"fmt"
	"goproject/git_try/banking"
)

func main() {
	account := banking.Account{Owner: "jk", Balance: 20000}
	fmt.Println(account)
}
