package mathAdditions

import "math"

func RoundToTwoDecimals(number float64) float64 {
	return math.Round(number * 100) / 100
}
