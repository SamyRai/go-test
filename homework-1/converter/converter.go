package converter

import (
	"github.com/SamyRai/go-test/homework-1/mathAdditions"
)

const rubToDollarExchangeRate = 70

func RublesToDollars(amount float64) float64 {
	return mathAdditions.RoundToTwoDecimals(amount / rubToDollarExchangeRate)
}

func DollarsToRubles(amount float64) float64 {
	return amount * rubToDollarExchangeRate
}
