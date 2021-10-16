package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) widtdrawPoint(amount int) {
	a1.balance -= amount
}
func (a2 account) widtdrawValue(amount int) {
	a2.balance -= amount
}
func (a3 account) widtdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "Joe", "Park"}
	mainA.widtdrawPoint(30)
	fmt.Println(mainA.balance)

	mainA.widtdrawValue(20)
	fmt.Println(mainA.balance)

	var mainB account = mainA.widtdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.widtdrawPoint(30)
	fmt.Println(mainB.balance)

}
