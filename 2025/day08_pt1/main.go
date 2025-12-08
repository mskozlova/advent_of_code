package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const N_ITER int = 1000

type Box struct {
	id      int
	x, y, z int
}

type Edge struct {
	id1, id2 int
	dist     int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	boxes := parseBoxes(string(input))
	edges := getEdges(boxes)
	graph := buildGraph(edges, len(boxes))
	circuits := countCircuits(&graph)
	fmt.Println(circuits)
	fmt.Println(circuits[0] * circuits[1] * circuits[2])
}

func parseBoxes(input string) []Box {
	boxes := make([]Box, 0, 0)
	for box_id, row := range strings.Split(input, "\n") {
		coords := strings.Split(row, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		boxes = append(boxes, Box{box_id, x, y, z})
	}
	return boxes
}

func sq(a int) int {
	return a * a
}

func (b1 *Box) dist(b2 Box) int {
	return sq(b1.x-b2.x) + sq(b1.y-b2.y) + sq(b1.z-b2.z)
}

func cmpEdges(e1, e2 Edge) int {
	return cmp.Compare(e1.dist, e2.dist)
}

func getEdges(boxes []Box) []Edge {
	distances := make([]Edge, 0, 0)
	for _, box1 := range boxes {
		for _, box2 := range boxes {
			if box1.id >= box2.id {
				continue
			}
			distances = append(distances, Edge{box1.id, box2.id, box1.dist(box2)})
		}
	}
	slices.SortFunc(distances, cmpEdges)
	return distances[:N_ITER]
}

func buildGraph(edges []Edge, total_nodes int) map[int][]int {
	graph := make(map[int][]int)
	for _, e := range edges {
		addEdge(&graph, e.id1, e.id2)
		addEdge(&graph, e.id2, e.id1)
	}

	for id := range total_nodes {
		_, exists := graph[id]
		if !exists {
			empty_list := make([]int, 0, 0)
			graph[id] = empty_list
		}
	}

	return graph
}

func addEdge(graph *map[int][]int, id1, id2 int) {
	node, exists := (*graph)[id1]
	if !exists {
		node = make([]int, 0, 0)
	}
	node = append(node, id2)
	(*graph)[id1] = node
}

func countCircuits(graph *map[int][]int) []int {
	circuits := make([]int, 0, 0)
	seen := make(map[int]int)

	for id := range *graph {
		_, e := seen[id]
		if e {
			continue
		}
		seen[id] = id
		queue := make([]int, 0, 0)
		queue = append(queue, id)
		circuit_size := 0

		for len(queue) > 0 {
			cur_id := queue[len(queue)-1]
			circuit_size += 1
			queue = queue[:len(queue)-1]
			neighbours, e := (*graph)[cur_id]
			for _, n := range neighbours {
				_, e = seen[n]
				if e {
					continue
				}
				seen[n] = n
				queue = append(queue, n)
			}
		}
		circuits = append(circuits, circuit_size)
	}
	slices.Sort(circuits)
	slices.Reverse(circuits)
	return circuits
}
