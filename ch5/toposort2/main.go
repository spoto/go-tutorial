package main

import (
	"fmt"
)

var prereqs = map[string]map[string]bool{
	"algorithms":            {"data structures": true},
	"calculus":              {"linear algebra": true},
	"compilers":             {"data structures": true, "formal languages": true, "computer organization": true},
	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	i := 1
	for course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i, course)
		i = i + 1
	}
}

func topoSort(m map[string]map[string]bool) map[string]bool {
	order := make(map[string]bool)
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)
	visitAll = func(items map[string]bool) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[item] = true
			}
		}
	}
	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}
	//sort.Strings(keys)
	visitAll(keys)
	return order
}
