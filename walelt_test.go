package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestWalletOperations(t *testing.T) {
	// Init the wallet
	wallet := &Wallet{balance: decimal.NewFromFloat(100)}

	// Deposit test
	wallet.deposit(decimal.NewFromFloat(50))
	if wallet.balance.Cmp(decimal.NewFromFloat(150)) != 0 {
		t.Errorf("Deposit failed. Expected balance to be $150, got: %s", wallet.balance.String())
	}

	// Withdraw test
	withdrawSuccess := wallet.withdraw(decimal.NewFromFloat(70))
	if !withdrawSuccess || wallet.balance.Cmp(decimal.NewFromFloat(80)) != 0 {
		t.Errorf("Withdrawal failed. Expected balance to be $80 after withdrawing $70, got: %s", wallet.balance.String())
	}

	// Init another wallet
	recipientWallet := &Wallet{balance: decimal.NewFromFloat(50)}

	// Transfer test
	sendSuccess := wallet.sendMoney(recipientWallet, decimal.NewFromFloat(30))
	if !sendSuccess || wallet.balance.Cmp(decimal.NewFromFloat(50)) != 0 || recipientWallet.balance.Cmp(decimal.NewFromFloat(80)) != 0 {
		t.Errorf("Send money failed. Incorrect balances. Sender: %s, Recipient: %s", wallet.balance.String(), recipientWallet.balance.String())
	}
}
