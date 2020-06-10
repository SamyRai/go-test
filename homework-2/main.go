package main

import "fmt"

func main() {
	fmt.Println("is 5 even?", isEvenNumber(5))
	fmt.Println("is 3 divisible by 3?", isDivisible(3, 3))
	printFib(10)
}

func isEvenNumber(number int) bool {
	return isDivisible(number, 2)
}

func isDivisible(number int, divider int) bool {
	return number%divider == 0
}

func printFib(maxNumber uint) {
	fmt.Println("Printing first", maxNumber, "numbers of Fibonacci:")
	var i uint = 0
	for i <= maxNumber {
		fmt.Println(i, "=", fib(i))
		i++
	}
}

func fib(n uint) uint {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
