package statistic

// Calculate the average
func Average(xs []float64) float64 {
	return Sum(xs) / float64(len(xs))
}

// Summarize all elements of the array
func Sum(elements []float64) (sum float64) {
	for _, element := range elements {
		sum += element
	}

	return sum
}
