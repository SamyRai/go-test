package main

import (
	"fmt"
	Account "github.com/SamyRai/go-test/homework-4/account"
)

type Money interface {
	YearLeftover() float64
}

func TotalMoneyLeft(accounts ...Money) (result float64) {
	for _, account := range accounts {
		result = result + account.YearLeftover()
	}

	return result
}

func main() {
	debitAccount := Account.Debit{
		Details: Account.Details{
			InitialSum:   6000.00,
			Transactions: []float64{1000.26, -200, -300, 100.10},
			Owner:        "Mister Twister",
		},
		Overdraft: 500,
	}

	creditAccount := Account.Credit{
		Details: Account.Details{
			InitialSum:   0,
			Transactions: []float64{-100, -500, -900, -1000},
		},
		YearlyInterestRate: 5,
	}

	fmt.Printf("Debit Account: %0.2f\n", debitAccount.YearLeftover())
	fmt.Printf("Credit Account: %0.2f\n", creditAccount.YearLeftover())
	fmt.Printf("In total you have: %0.2f\n", TotalMoneyLeft(debitAccount, creditAccount))
}
