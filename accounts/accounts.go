package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	owner   string
	balance int
}

var errNomoney = errors.New("Can't withdraw")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//메소드
func (a *Account) Deposit(amount int) {
	a.balance += amount

}
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNomoney
	}
	a.balance -= amount
	return nil
}

func (a Account) Balance() int {
	return a.balance
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "`s account.\nHas: ", a.Balance())
}
