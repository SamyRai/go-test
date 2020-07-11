package main

import (
	"fmt"
)

// Я тут посто игрался и не могу понять почему len всегда равно 0?

func main() {
	done := make(chan int)
	for i := 0; i < 10; i++ {
		go func(num int) {
			fmt.Println(num)
			done <- num
		}(i)
	}

	for x :=0; x<10; x++ {
		fmt.Println("len", len(done))
		fmt.Println(<-done)
	}

}