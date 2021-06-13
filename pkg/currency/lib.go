package currency

import (
	"fmt"
	"time"
)

type Bank struct {
	Accounts  []*Account
	UserTable map[string]bool
	clock     *time.Ticker
}

func New() *Bank {
	return &Bank{
		Accounts:  []*Account{},
		UserTable: make(map[string]bool),
		clock:     time.NewTicker(5 * time.Second),
	}
}

func (b *Bank) CreateAccount(user string) (*Account, error) {
	if b.UserTable[user] == false {
		b.UserTable[user] = true
		return NewAccount(user), nil
	}
	return nil, fmt.Errorf("User %s already has an account.", user)
}
