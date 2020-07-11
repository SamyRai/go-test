package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	go spinner(50 * time.Millisecond)
	//time.Sleep(time.Second * 10) // Alternative simple solution just sleep on main function for 10 seconds
	tick := 1
	for range time.Tick(1 * time.Second) {
		if tick > 10 {
			break
		}

		fmt.Println(tick, "seconds passed")
		tick++
	}
}
