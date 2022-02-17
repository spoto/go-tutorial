package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate_left(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func rotate_right(s []int, n int) {
	reverse(s)
	reverse(s[:n])
	reverse(s[n:])
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
	rotate_left(a[:], 2)
	fmt.Println(a)
	rotate_right(a[:], 2)
	fmt.Println(a)
}
