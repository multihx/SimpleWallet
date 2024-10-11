package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	walletApp := WalletApp{
		users: make(map[string]*User),
	}

	// Create users and perform wallet operations
	walletApp.createUser("Alice")
	walletApp.createUser("Bob")

	walletApp.deposit("Alice", decimal.NewFromFloat(100))
	walletApp.sendMoney("Alice", "Bob", decimal.NewFromFloat(50))

	// Check balances
	aliceBalance, _ := walletApp.checkBalance("Alice")
	bobBalance, _ := walletApp.checkBalance("Bob")

	fmt.Printf("Alice's balance: $%s\n", aliceBalance.String())
	fmt.Printf("Bob's balance: $%s\n", bobBalance.String())
}
