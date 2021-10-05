package main

import (
	"fmt"

	"github.com/Lsapee/GO_Study/accounts"
)

func main() {
	account := accounts.NewAccount("jk")
	account.Deposit(400)
	fmt.Println(account.Balance())
	err := account.Withdraw(450)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
