package accounts

import "errors"

//Account struct
type Account struct {
	owner   string
	balance int
}

var ErrNoMoney = errors.New("your balance is lack")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//메소드
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return ErrNoMoney
	}
	a.balance -= amount
	return nil
}
