package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(len(a))
	fmt.Println(cap(a))
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	var q = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	q[1]++
	fmt.Println(r[2])
	r = q
	q[2] = 4
	fmt.Println(r[2])
	fmt.Println(q[2])
	symbol := [...]string{USD: "$", GBP: "£", EUR: "E", RMB: "R"}
	fmt.Println(RMB, symbol[GBP])
}
