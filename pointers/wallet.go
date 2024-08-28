package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin that represents a number of bitcoins
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet represents a wallet containing a balance of bitcoins
type Wallet struct {
	balance Bitcoin
}

// Deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds means a wallet does not have enough Bitcoin to perform a withdraw.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw subtracts some Bitcoin from the wallet, returning an error if it cannot be performed.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
