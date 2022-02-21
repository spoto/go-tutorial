package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("ciao", "amico")
	addEdge("amico", "caro")
	addEdge("ciao", "bello")
	fmt.Printf("%t\n", hasEdge("amico", "caro"))
	fmt.Printf("%t\n", hasEdge("amico", "cara"))
	fmt.Printf("%t\n", hasEdge("amica", "caro"))
	fmt.Printf("%t\n", hasEdge("amica", "cara"))
}
