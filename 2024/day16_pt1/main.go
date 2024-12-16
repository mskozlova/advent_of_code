package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

const (
	step_point int = 1
	turn_point int = 1000
	inf_dist       = 100000000000
)

type direction struct {
	v, h int
}

type coord struct {
	row, col int
}

type node struct {
	c   coord
	dir direction
}

type QueueNode struct {
	node  node
	dist  int
	index int
}

// https://pkg.go.dev/container/heap#example-package-PriorityQueue
type PriorityQueue []*QueueNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*QueueNode)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *QueueNode, n node, dist int) {
	item.node = n
	item.dist = dist
	heap.Fix(pq, item.index)
}

func main() {
	file_name := "2024/day16_pt1/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	start, end, field := parseInput(string(input))
	fmt.Println(findMinDistance(start, end, field))
}

func initQueue(start coord, field []string) (PriorityQueue, map[node]*QueueNode) {
	directions := []direction{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	pq := make(PriorityQueue, len(field)*len(field[0])*len(directions))
	links := make(map[node]*QueueNode)

	i := 0
	for row := range len(field) {
		for col := range len(field[0]) {
			for _, d := range directions {
				dist := inf_dist
				if (row == start.row) && (col == start.col) && (d.v == 0) && (d.h == 1) {
					dist = 0
				}
				links[node{coord{row, col}, d}] = &QueueNode{
					node{coord{row, col}, d},
					dist,
					i,
				}
				pq[i] = links[node{coord{row, col}, d}]
				i += 1
			}
		}
	}
	heap.Init(&pq)

	return pq, links
}

func findMinDistance(start coord, end coord, field []string) int {
	pq, links := initQueue(start, field)

	for {
		unvis_node := heap.Pop(&pq).(*QueueNode)
		fmt.Printf("Visiting node row %d, col %d, v %d, h %d - dist %d\n",
			unvis_node.node.c.row, unvis_node.node.c.col,
			unvis_node.node.dir.v, unvis_node.node.dir.h, unvis_node.dist)
		if (unvis_node.node.c.row == end.row) && (unvis_node.node.c.col == end.col) {
			return unvis_node.dist
		}

		neighbours := getNeighbours(unvis_node.node, field)
		fmt.Println("Neighbours: ", neighbours)
		for _, neighbour := range neighbours {
			n_link, _ := links[neighbour.node]
			new_dist := min(n_link.dist, unvis_node.dist+neighbour.dist)
			pq.update(n_link, n_link.node, new_dist)
		}
	}
}

func (d direction) getNeighbours() []direction {
	if d.h != 0 {
		return []direction{{v: 1, h: 0}, {v: -1, h: 0}}
	}
	return []direction{{v: 0, h: 1}, {v: 0, h: -1}}
}

func getNeighbours(n node, field []string) []QueueNode {
	neighbours := make([]QueueNode, 0)

	// 1. step option
	next_pos := coord{row: n.c.row + n.dir.v, col: n.c.col + n.dir.h}
	if (next_pos.row >= 0) && (next_pos.row < len(field)) && (next_pos.col >= 0) && (next_pos.col < len(field[0])) {
		// we're in the field
		if string(field[next_pos.row][next_pos.col]) != "#" {
			// can make a step
			neighbours = append(neighbours, QueueNode{node{next_pos, n.dir}, step_point, -1})
		}
	}

	// 2. turn option
	for _, new_dir := range n.dir.getNeighbours() {
		neighbours = append(neighbours, QueueNode{node{n.c, new_dir}, turn_point, -1})
	}

	return neighbours
}

func parseInput(input string) (coord, coord, []string) {
	field := strings.Split(input, "\n")
	start := coord{}
	end := coord{}

	for row := range len(field) {
		for col := range len(field[0]) {
			if string(field[row][col]) == "S" {
				start.row = row
				start.col = col
			} else if string(field[row][col]) == "E" {
				end.row = row
				end.col = col
			}
		}
	}
	return start, end, field
}
