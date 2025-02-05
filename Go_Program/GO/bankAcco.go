package main

import "fmt"

type BankAccount struct {
	balance float64
	owner   string
	accountNumber int
}

func (ba *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
        return fmt.Errorf("Invalid deposit amount. Amount should be greater than zero.")
    }
    ba.balance += amount
    return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
        return fmt.Errorf("Invalid withdrawal amount. Amount should be greater than zero.")
    }
    if amount > ba.balance {
        return fmt.Errorf("Insufficient funds.")
    }
    ba.balance -= amount
    return nil
}

func (ba *BankAccount) GetBalance() error {
	if ba.balance < 0 {
        return fmt.Errorf("Account balance is negative.")
    }
    fmt.Printf("Balance: %.2f\n", ba.balance)
    return nil
}

type AccountBehavior struct {
	Deposit(amount float64) error
    Withdraw(amount float64) error
    CheckBalance() float64
}

func PerformTransaction(account AccountBehavior, action string, amount float64) error {
    switch action {
    case "deposit":
        return account.Deposit(amount)
    case "withdraw":
        return account.Withdraw(amount)
    default:
        return errors.New("invalid action")
    }
}

func main() {
    // Create a new account
    account := Account{
        AccountNumber: "123456",
        HolderName:    "John Doe",
        Balance:       1000.00,
    }

    // Perform transactions
    fmt.Println("Initial Balance:", account.CheckBalance())

    err := PerformTransaction(&account, "deposit", 500.00)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("After depositing $500, Balance:", account.CheckBalance())
    }

    err = PerformTransaction(&account, "withdraw", 200.00)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("After withdrawing $200, Balance:", account.CheckBalance())
    }

    err = PerformTransaction(&account, "withdraw", 2000.00)
    if err != nil {
        fmt.Println("Error:", err) // Output: Error: insufficient funds
    }
}