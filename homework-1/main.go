package main

import (
	"fmt"
	"github.com/SamyRai/go-test/homework-1/bank"
	"github.com/SamyRai/go-test/homework-1/converter"
	"github.com/SamyRai/go-test/homework-1/triangle"
	"log"
	"strconv"
)

func main() {
	var taskSelector string
	fmt.Println("Please enter which task you want to solve:")
	fmt.Println("1 for converter, 2 for triangle, 3 for deposit calculations")
	fmt.Scanln(&taskSelector)

	switch taskSelector {
	case "1":
		converterTask()
	case "2":
		geometryTask()
	case "3":
		depositTask()
	default:
		fmt.Println("Please select only from 1 to 3")
	}

}

func converterTask() {
	var inputAmountToConvert string
	fmt.Println("How much rubles do you have?")
	fmt.Scanln(&inputAmountToConvert)
	AmountToConvert, err := strconv.ParseFloat(inputAmountToConvert, 64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("From", AmountToConvert, "rubles you get", converter.RublesToDollars(AmountToConvert), "$")
}

func geometryTask() {
	var inputLegOne, inputLegTwo string
	fmt.Println("Please enter the first leg of the right triangle")
	fmt.Scanln(&inputLegOne)
	fmt.Println("Now the second leg")
	fmt.Scanln(&inputLegTwo)

	legOne, err := strconv.ParseFloat(inputLegOne, 64)
	if err != nil {
		log.Fatalln(err)
	}

	legTwo, err := strconv.ParseFloat(inputLegTwo, 64)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Square:", triangle.Square(legOne, legTwo))
	fmt.Println("Hypotenuse:", triangle.Hypotenuse(legOne, legTwo))
	fmt.Println("Perimeter:", triangle.Perimeter(legOne, legTwo))
}

func depositTask() {
	var inputAmount string
	var inputInterestRate string
	fmt.Println("Please enter the amount that you want to deposit")
	fmt.Scanln(&inputAmount)
	fmt.Println("What is your interest rate?")
	fmt.Scanln(&inputInterestRate)

	amount, err := strconv.ParseFloat(inputAmount, 64)
	if err != nil {
		log.Fatalln(err)
	}
	rate, err := strconv.ParseFloat(inputInterestRate, 64)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("In 5 years you will get:", bank.CalculateDeposit(amount, rate, 5))
}
