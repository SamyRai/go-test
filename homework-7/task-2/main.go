package main

import (
	"fmt"
)

const maxNumber = 50

func main() {
	naturals := make(chan int, maxNumber)
	squares := make(chan int, maxNumber)

	// генерация
	go func() {
		for x := 0; x<=maxNumber; x++ {
			naturals <- x
		}
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for x := 0; x<=maxNumber; x++  { // Пытался тут использовать range, но почему-то выводит в дэдлок...
		fmt.Println(<-squares)
	}

}
