package main

import (
	"fmt"

	"github.com/Lsapee/GO_Study/accounts"
)

func main() {
	account := accounts.NewAccount("jk")
	account.Deposit(400)
	fmt.Println(account.Balance())
	err := accounts.Withdraw(30)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
