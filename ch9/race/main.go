package main

import (
	"fmt"
	"time"
)

func main() {
	var x, y int

	go func() {
		x = 1
		fmt.Print("y:", y, " ")
	}()

	go func() {
		y = 1
		fmt.Print("x:", x, " ")
	}()

	time.Sleep(1 * time.Second)
}
