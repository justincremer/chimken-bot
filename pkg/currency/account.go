package currency

import (
	"fmt"
)

type Account struct {
	ID      string
	User    string
	Balance int
}

func NewAccount(user string) *Account {
	return &Account{
		ID:      "",
		User:    user,
		Balance: 0,
	}
}

func (a *Account) Update(amount int) error {
	if amount < 0 && -amount > a.Balance {
		return fmt.Errorf("Cannot dedecut $%s from account with $%s", amount, a.Balance)
	}
	a.Balance += amount
	return nil
}

func (a *Account) Format() string {
	return fmt.Sprintf("Owner")
}
