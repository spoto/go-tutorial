package main

import (
	"fmt"
	"github.com/spoto/go-tutorial/ch7/bytecounter"
)

func main() {
	var c bytecounter.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
