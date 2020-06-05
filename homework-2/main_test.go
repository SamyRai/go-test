package main

import (
	"testing"
)

func TestIsEvenNumber(t *testing.T) {
	evenNumbers := [6]int{2, 4, 6, 8, 10, 12}
	oddNumbers := [6]int{1, 3, 5, 7, 9, 11}

	for _, evenNumber := range evenNumbers {
		if isEvenNumber(evenNumber) {
			t.Logf("%d is an even number", evenNumber)
		} else {
			t.Errorf("%d is an odd number", evenNumber)
		}
	}
	for _, oddNumber := range oddNumbers {
		if isEvenNumber(oddNumber) {
			t.Errorf("%d is an even number", oddNumber)
		} else {
			t.Logf("%d is an odd number", oddNumber)
		}
	}
}

type testItem struct {
	number  int
	divider int
	result  bool
}

func TestIsDivisible(t *testing.T) {
	testItems := []testItem{
		{3, 3, true},
		{4, 3, false},
		{5, 3, false},
		{9, 3, true},
		{24, 3, true},
	}

	for _, item := range testItems {
		if isDivisible(item.number, item.divider) == item.result {
			t.Logf("isDivisible(%d, %d) returns %t", item.number, item.divider, item.result)
		} else {
			t.Errorf("isDivisible(%d, %d) returns %t", item.number, item.divider, item.result)
		}
	}
}
