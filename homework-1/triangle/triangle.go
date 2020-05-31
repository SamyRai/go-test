package triangle

import (
	"github.com/SamyRai/go-test/homework-1/mathAdditions"
	"math"
)

func Square(legOne, legTwo float64) float64 {
	return mathAdditions.RoundToTwoDecimals((legOne * legTwo) / 2)
}

func Hypotenuse(legOne, legTwo float64) float64 {
	hypotenuse := math.Sqrt(math.Pow(legOne, 2) + math.Pow(legTwo, 2))

	return mathAdditions.RoundToTwoDecimals(hypotenuse)
}

func Perimeter(legOne, legTwo float64) float64 {
	return legOne + legTwo + Hypotenuse(legOne, legTwo)
}
