package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func initWalletApp() *WalletApp {
	walletApp := WalletApp{
		users: make(map[string]*User),
	}

	// Create user Alice and Bob
	walletApp.createUser("Alice")
	walletApp.createUser("Bob")

	return &walletApp
}

func TestWalletDeposit(t *testing.T) {
	walletApp := initWalletApp()
	walletApp.deposit("Alice", decimal.NewFromFloat(100))
	aliceBalance, _ := walletApp.checkBalance("Alice")
	if !aliceBalance.Equal(decimal.NewFromFloat(100)) {
		t.Errorf("Deposit failed. Expected Alice's balance to be $100, got: %s", aliceBalance.String())
	}
}

func TestWalletWithdraw(t *testing.T) {
	walletApp := initWalletApp()
	walletApp.deposit("Alice", decimal.NewFromFloat(50))
	withdrawSuccess := walletApp.withdraw("Alice", decimal.NewFromFloat(30))
	aliceBalanceAfterWithdraw, _ := walletApp.checkBalance("Alice")
	if !withdrawSuccess || aliceBalanceAfterWithdraw.Cmp(decimal.NewFromFloat(20)) != 0 {
		t.Errorf("Withdraw failed. Alice's balance should be $20 after withdrawing $30, got: %s", aliceBalanceAfterWithdraw.String())
	}

}
func TestWithdrawMoreThanBalance(t *testing.T) {
	walletApp := initWalletApp() // Init the wallet app

	// Deposit funds to Alice wallet
	walletApp.deposit("Alice", decimal.NewFromFloat(100))

	// Try to withdraw amount exceed the alice wallet balance
	withdrawSuccess := walletApp.withdraw("Alice", decimal.NewFromFloat(120))
	aliceBalance, _ := walletApp.checkBalance("Alice")

	if withdrawSuccess {
		t.Error("Withdrawal of amount greater than balance should fail.")
	}

	if aliceBalance.Cmp(decimal.NewFromFloat(100)) != 0 {
		t.Errorf("Withdrawal of amount greater than balance should not affect balance. Expected Alice's balance to be $100, got: %s", aliceBalance.String())
	}
}
