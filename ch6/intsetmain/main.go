package main

import (
	"fmt"
	"github.com/spoto/go-tutorial/ch6/intset"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	fmt.Printf("%v has length %d\n", x.String(), x.Len())
	y.Add(9)
	y.Add(42)
	fmt.Printf("%v has length %d\n", y.String(), y.Len())
	x.UnionWith(&y)
	fmt.Printf("%v has length %d\n", x.String(), x.Len())
	x.Remove(42)
	fmt.Printf("%v has length %d\n", x.String(), x.Len())
	x.Remove(43)
	fmt.Printf("%v has length %d\n", x.String(), x.Len())
	fmt.Println(x.Has(9), x.Has(123))
	fmt.Println(x.Elems())
	copy := x.Copy()
	fmt.Printf("%v has length %d\n", copy.String(), copy.Len())
	x.Clear()
	fmt.Printf("%v has length %d\n", x.String(), x.Len())
}
