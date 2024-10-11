package main

import (
	"sync"

	"github.com/shopspring/decimal"
)

type User struct {
	name   string
	wallet *Wallet
}

type WalletApp struct {
	users map[string]*User
	mu    sync.Mutex
}

func (wa *WalletApp) createUser(name string) {
	wa.mu.Lock()
	defer wa.mu.Unlock()
	if _, exists := wa.users[name]; !exists {
		wa.users[name] = &User{name: name, wallet: &Wallet{balance: decimal.NewFromFloat(0)}}
	}
}

func (wa *WalletApp) getUserWallet(name string) (*Wallet, bool) {
	wa.mu.Lock()
	defer wa.mu.Unlock()
	user, exists := wa.users[name]
	if !exists {
		return &Wallet{}, false
	}
	return user.wallet, true
}

func (wa *WalletApp) deposit(name string, amount decimal.Decimal) {
	wa.mu.Lock()
	defer wa.mu.Unlock()
	user, exists := wa.users[name]
	if exists {
		user.wallet.deposit(amount)
		wa.users[name] = user
	}
}

func (wa *WalletApp) withdraw(name string, amount decimal.Decimal) bool {
	wa.mu.Lock()
	defer wa.mu.Unlock()
	user, exists := wa.users[name]
	if exists {
		if user.wallet.withdraw(amount) {
			wa.users[name] = user
			return true
		}
	}
	return false
}

func (wa *WalletApp) sendMoney(sender string, recipient string, amount decimal.Decimal) bool {
	wa.mu.Lock()
	defer wa.mu.Unlock()
	senderUser, senderExists := wa.users[sender]
	recipientUser, recipientExists := wa.users[recipient]
	if senderExists && recipientExists {
		if senderUser.wallet.sendMoney(recipientUser.wallet, amount) {
			wa.users[sender] = senderUser
			wa.users[recipient] = recipientUser
			return true
		}
	}
	return false
}

func (wa *WalletApp) checkBalance(name string) (decimal.Decimal, bool) {
	wa.mu.Lock()
	defer wa.mu.Unlock()
	user, exists := wa.users[name]
	if exists {
		return user.wallet.balance, true
	}
	return decimal.Zero, false
}
