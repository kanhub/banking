package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errWithdraw = errors.New("Can't withdraw")

// NewAccount creator
func NewAccount(name string) *Account {
	account := Account{owner: name, balance: 0}
	return &account
}

// Deposit func with receiver a of type Account for setting up deposits
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance displays account balance
func (a *Account) Balance() int {
	return a.balance
}

// Withdraw from our account
// errors have two values, error and nil so must return nil
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errWithdraw
	}
	a.balance -= amount
	return nil
}

// ChangeOwner changees current owner of an account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner returns the owner of an account
func (a *Account) Owner() string {
	return a.owner
}

// String representation of object
func (a *Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account has: $", a.Balance())
}
