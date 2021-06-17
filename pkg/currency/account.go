package currency

import (
	"fmt"
)

type Account struct {
	User    string
	Balance int
}

func NewAccount(user string) *Account {
	return &Account{
		User:    user,
		Balance: 100,
	}
}

func (a *Account) Update(amount int) error {
	if amount < 0 && -amount > a.Balance {
		return fmt.Errorf("Cannot dedecut $%d from account with $%d", amount, a.Balance)
	}
	a.Balance += amount
	return nil
}

func (a *Account) Format() string {
	return fmt.Sprintf("Owner")
}
