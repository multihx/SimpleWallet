# Wallet Application

This Go application simulates a wallet system that keeps track of users and their respective wallet balances. Users can deposit money, withdraw funds, send money to other users, and check their balances.

## Application Architecture

The wallet application follows a simple yet modular architecture to handle user interactions, wallet operations, and data management efficiently. It consists of the following components:

Main Application: Controls the flow of the application, handles user creation, and delegates operations to the appropriate modules.
User Management Module: Responsible for creating, storing, and managing user information, including user names and wallet balances.

Wallet Operations Module: Handles all wallet-related functionalities such as deposit, withdrawal, sending money, and checking balances.

Transaction History Module(To be implemented): Manages the transaction history for each user, recording details of all monetary transactions for reference and auditing purposes.

## Setup

1. Ensure you have Go installed on your machine.
2. Clone this repository to your local machine.
3. Install the required dependencies using `go mod tidy`.

## Usage

1. Run the application by executing `go run main.go`.
2. The application creates users and performs wallet operations such as deposit, withdraw, send money, and check balance.
3. Check the console output to see the results of the wallet operations.

## Functionality

- `createUser(name)`: Creates a new user with the specified name.
- `deposit(name, amount)`: Deposits the specified amount into the user's wallet.
- `withdraw(name, amount)`: Withdraws the specified amount from the user's wallet.
- `sendMoney(sender, recipient, amount)`: Sends the specified amount from the sender's wallet to the recipient's wallet.
- `checkBalance(name)`: Retrieves the balance of the user's wallet.

## Sample Usage

```go
walletApp := WalletApp{
    users: make(map[string]User),
}

walletApp.createUser("Alice")
walletApp.createUser("Bob")

walletApp.deposit("Alice", decimal.NewFromFloat(100))
walletApp.sendMoney("Alice", "Bob", decimal.NewFromFloat(50))

aliceBalance, _ := walletApp.checkBalance("Alice")
bobBalance, _ := walletApp.checkBalance("Bob")

fmt.Printf("Alice's balance: $%s\n", aliceBalance.String())
fmt.Printf("Bob's balance: $%s\n", bobBalance.String())
```

## Further Development

- Transaction History: Implement a feature to keep track of transaction history for each user, including details such as deposit, withdrawal, and transfers.
-  Error Handling: Enhance error handling mechanisms to provide informative messages to users in case of failed transactions or invalid inputs.
Authentication: Introduce user authentication mechanisms to secure user accounts and prevent unauthorized access to wallets.
- Concurrency: Optimize the application for concurrency by exploring ways to improve performance in scenarios with multiple simultaneous transactions.
- Web Interface: Develop a web interface using a framework like Gin or Echo to interact with the wallet system through a user-friendly interface.
- Database Integration: Integrate a database (e.g., PostgreSQL, MySQL) to persist user data and transaction history for improved data management and scalability.
- Notifications: Implement a notification system to alert users about important events, such as successful transactions, low balances, or account activities.
- Unit Testing: Expand the test suite to cover more edge cases and ensure robustness and correctness of the application.

- Logging: Add logging functionality to track application events, errors, and user activities for troubleshooting and auditing purposes.