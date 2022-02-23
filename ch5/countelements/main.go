package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	counters := make(map[string]int)
	outline(counters, doc)
	for k, v := range counters {
		fmt.Printf("%15s %d\n", k, v)
	}
}

func outline(counters map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counters[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(counters, c)
	}
}
