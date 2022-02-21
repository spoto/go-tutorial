package main

import "fmt"

func myRotateLeft(s []int, n int) []int {
	overflow := make([]int, n)
	copy(overflow, s[:n])
	copy(s, s[n:])
	copy(s[len(s)-n:], overflow)
	return s
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	b := myRotateLeft(a[:], 2)
	fmt.Println(b)
}
