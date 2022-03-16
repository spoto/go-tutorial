package main

import (
	"fmt"
	"github.com/spoto/go-tutorial/ch8/thumbnail"
	"os"
)

func main() {
	args := os.Args[1:]
	input := make(chan string, len(args))
	for _, thumb := range args {
		input <- thumb
	}
	close(input)
	fmt.Printf("the input channel has length %d\n", len(input))
	thumbnail.MakeThumbnails6(input)
}
