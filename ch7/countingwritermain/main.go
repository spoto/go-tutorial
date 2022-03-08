package main

import (
	"fmt"
	"github.com/spoto/go-tutorial/ch7/countingwriter"
	"os"
)

func main() {
	w, counter := countingwriter.CountingWriter(os.Stdout)
	w.Write([]byte("Nel mezzo del cammin di nostra vita, mi ritrovai per una selva oscura"))
	fmt.Printf("\ncounter = %d\n", *counter)
}
