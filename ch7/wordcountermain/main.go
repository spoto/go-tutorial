package main

import (
	"fmt"
	"github.com/spoto/go-tutorial/ch7/wordcounter"
)

func main() {
	var c wordcounter.WordCounter
	c.Write([]byte("Nel mezzo del cammin di nostra vita, mi ritrovai per una selva oscura"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
