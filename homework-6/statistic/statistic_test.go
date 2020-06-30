package statistic

import "testing"

type testpair struct {
	values []float64
	result float64
}

func TestAverageSet(t *testing.T) {
	var tests = []testpair {
		{[]float64{1,2}, 1.5},
		{[]float64{1,1,1,1,1,1}, 1},
		{[]float64{-1,1}, 0},
	}

	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestSum(t *testing.T) {
	var tests = []testpair {
		{[]float64{1,2}, 3},
		{[]float64{1,1,1,1,1,1}, 6},
		{[]float64{-1,1}, 0},
		{[]float64{22,66, 11}, 99},
	}

	for _, pair := range tests {
		result := Sum(pair.values)

		if result != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", result,
			)
		}
	}
}