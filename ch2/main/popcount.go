package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	sum := 0
	for i := 0; i < 64; i += 8 {
		sum += int(pc[byte(x>>i)])
	}

	return sum
}

func PopCount3(x uint64) int {
	counter := 0
	for x != 0 {
		counter++
		x = x & (x - 1)
	}

	return counter
}

func main() {
	const n = 123456
	fmt.Printf("population count of %d = %d\n", n, PopCount(n))
	fmt.Printf("population count of %d = %d\n", n, PopCount2(n))
	fmt.Printf("population count of %d = %d\n", n, PopCount3(n))
}
