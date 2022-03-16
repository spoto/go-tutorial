package main

import (
	"flag"
	"fmt"
	"github.com/spoto/go-tutorial/ch5/links"
	"log"
)

var maxDepth = flag.Uint("depth", 3, "max depth")

func main() {
	flag.Parse()

	type batch struct {
		urls  []string
		depth uint
	}

	workList := make(chan batch)
	n := 1
	go func() { workList <- batch{flag.Args(), 1} }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		b := <-workList
		depth := b.depth
		if depth <= *maxDepth {
			for _, link := range b.urls {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						workList <- batch{crawl(link, depth), depth + 1}
					}(link)
				}
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string, depth uint) []string {
	fmt.Printf("%d: %s\n", depth, url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
