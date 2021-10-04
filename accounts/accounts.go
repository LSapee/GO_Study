package accounts

//Account struct
type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//메소드
func (a Account) Deposit(amount int) {
	a.balance += amount
}