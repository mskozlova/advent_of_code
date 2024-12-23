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
	file_name := "2024/day23_pt2/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	graph := createGraph(strings.Split(string(input), "\n"))
	nodes := make([]string, 0)
	for _, n := range graph {
		nodes = append(nodes, n.id)
	}
	current_clique := make([]string, 0)
	max_clique := make([]string, 0)
	max_size := 0
	findMaxClique(nodes, 0, &graph, &current_clique, &max_size, &max_clique)
	slices.Sort(max_clique)
	fmt.Println(strings.Join(max_clique, ","))
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

func findMaxClique(nodes []string, index int, graph *map[string]node, current_clique *[]string, max_size *int, max_clique *[]string) {
	if index >= len(nodes) {
		return
	}
	for i := index; i < len(nodes); i++ {
		id := nodes[i]
		size := len(*current_clique)
		*current_clique = append(*current_clique, id)
		is_clique := isClique(*current_clique, graph)
		if is_clique {
			if len(*current_clique) > *max_size {
				fmt.Println("New best:", *current_clique)
				*max_size = len(*current_clique)
				*max_clique = make([]string, *max_size)
				copy(*max_clique, *current_clique)
			}
			findMaxClique(nodes, i+1, graph, current_clique, max_size, max_clique)
		}
		*current_clique = (*current_clique)[:size]
	}
}

func isClique(candidates []string, graph *map[string]node) bool {
	for i := range len(candidates) {
		for j := range len(candidates) {
			if i == j {
				continue
			}

			n, _ := (*graph)[candidates[i]]
			_, is_connected := n.next[candidates[j]]

			if !is_connected {
				return false
			}
		}
	}
	return true
}
