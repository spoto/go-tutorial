package main

import (
	"fmt"
	"github.com/spoto/go-tutorial/ch6/geometry"
)

func main() {
	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(geometry.PathDistance(perim))
	fmt.Println(perim.Distance())

	r := &geometry.Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p := geometry.Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p)

	p = geometry.Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p)
}
