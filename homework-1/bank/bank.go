package bank

import "github.com/SamyRai/go-test/homework-1/mathAdditions"

func CalculateDeposit(amount, rate float64, years int) float64 {
	yearlyInterest := amount * rate / 100
	return mathAdditions.RoundToTwoDecimals(amount + yearlyInterest * 5)
}
