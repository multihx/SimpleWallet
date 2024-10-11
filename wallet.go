package main

import (
	"sync"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	balance decimal.Decimal
	mu      sync.Mutex
}

func (w *Wallet) deposit(amount decimal.Decimal) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance = w.balance.Add(amount)
}

func (w *Wallet) withdraw(amount decimal.Decimal) bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.balance.GreaterThanOrEqual(amount) {
		w.balance = w.balance.Sub(amount)
		return true
	}
	return false
}

func (w *Wallet) sendMoney(recipient *Wallet, amount decimal.Decimal) bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.balance.GreaterThanOrEqual(amount) {
		w.balance = w.balance.Sub(amount)
		recipient.mu.Lock()
		defer recipient.mu.Unlock()
		recipient.balance = recipient.balance.Add(amount)
		return true
	}
	return false
}
