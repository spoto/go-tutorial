package main

import "fmt"

func noAdjacent(strings []string) []string {
	i := 0
	previous := ""
	first := true
	for _, s := range strings {
		if first || s != previous {
			previous = s
			strings[i] = s
			i++
		} else {
			previous = s
		}

		first = false
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "one", "one", "", "", "one", "three", "three"}
	fmt.Printf("%q\n", noAdjacent(data))
	fmt.Printf("%q\n", data)
}
