package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func diff(first [32]byte, second [32]byte) int {
	counter := 0
	for pos := range first {
		counter += int(pc[first[pos]^second[pos]])
	}
	return counter
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("Number of different bits: %d\n", diff(c1, c2))
}
