package main

import (
	"fmt"
	"goproject/Nomad_Coders/banking"
)

func main() {
	account := banking.Account{Owner: "jk", Balance: 20}
	fmt.Println(account)
}
