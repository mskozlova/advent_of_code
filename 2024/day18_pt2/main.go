package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	field_size int = 71
)

type coord struct {
	row, col int
}
type direction struct {
	v, h int
}
type bt struct {
	c  coord
	ts int
}

func main() {
	file_name := "2024/day18_pt1/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	bts := parseBytes(string(input))
	for i := 1024; i < len(bts); i++ {
		fmt.Println(i)
		bt_set := makeSet(bts, i)
		if findShortestPath(&bt_set, field_size, field_size) < 0 {
			fmt.Printf("%d,%d\n", bts[i-1].c.col, bts[i-1].c.row)
			break
		}
	}
}

func parseBytes(input string) []bt {
	rows := strings.Split(input, "\n")
	bts := make([]bt, len(rows))
	for i, row := range rows {
		coords_str := strings.Split(row, ",")
		row, _ := strconv.Atoi(coords_str[1])
		col, _ := strconv.Atoi(coords_str[0])
		bts[i] = bt{coord{row, col}, i}
	}
	return bts
}

func makeSet(bts []bt, n_first int) map[coord]bool {
	bt_set := make(map[coord]bool)
	for i := range n_first {
		bt_set[bts[i].c] = true
	}
	return bt_set
}

func findShortestPath(bt_set *map[coord]bool, total_rows, total_cols int) int {
	queue := []bt{{coord{0, 0}, 0}}
	visited := make(map[coord]bool)
	visited[coord{0, 0}] = true
	finish := coord{total_rows - 1, total_cols - 1}

	directions := []direction{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if (pos.c.row == finish.row) && (pos.c.col == finish.col) {
			return pos.ts
		}

		for _, d := range directions {
			next_pos := coord{pos.c.row + d.v, pos.c.col + d.h}
			if (next_pos.row < 0) || (next_pos.row >= total_rows) || (next_pos.col < 0) || (next_pos.col >= total_cols) {
				continue
			}
			_, is_visited := visited[next_pos]
			_, is_obstacle := (*bt_set)[next_pos]
			if is_visited || is_obstacle {
				continue
			}
			visited[next_pos] = true
			queue = append(queue, bt{next_pos, pos.ts + 1})
		}
	}
	return -1
}
