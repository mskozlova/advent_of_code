package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type node struct {
	id   string
	next map[string]bool
}

func main() {
	file_name := "2024/day23_pt1/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	graph := createGraph(strings.Split(string(input), "\n"))
	connected := findConnectedTriplets(&graph)
	fmt.Println(connected)
	fmt.Println(len(connected))
}

func createGraph(input []string) map[string]node {
	graph := make(map[string]node)
	for _, row := range input {
		id1 := strings.Split(row, "-")[0]
		id2 := strings.Split(row, "-")[1]
		addConnection(id1, id2, &graph)
		addConnection(id2, id1, &graph)
	}
	return graph
}

func addConnection(from, to string, graph *map[string]node) {
	_, from_exists := (*graph)[from]
	if !from_exists {
		next := make(map[string]bool)
		n := node{from, next}
		(*graph)[from] = n
	}
	(*graph)[from].next[to] = true
}

func findConnectedTriplets(graph *map[string]node) map[string]bool {
	triplets := make(map[string]bool)
	for _, n1 := range *graph {
		if !strings.HasPrefix(n1.id, "t") {
			continue
		}
		if len(n1.next) >= 2 {
			for n2 := range n1.next {
				if len(n2) >= 2 {
					for n3 := range (*graph)[n2].next {
						_, is_connected := n1.next[n3]
						if is_connected {
							triplet := []string{n1.id, n2, n3}
							slices.Sort(triplet)
							triplets[strings.Join(triplet, ",")] = true
						}
					}
				}
			}
		}
	}
	return triplets
}
