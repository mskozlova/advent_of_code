package main

import (
	"fmt"
	"os"
	"strings"
)

type direction struct {
	v, h int
}

type coord struct {
	row, col int
}

type plant struct {
	area, perimeter int
}

func main() {
	input, _ := os.ReadFile("2024/day12_pt1/input.txt")
	field := strings.Split(string(input), "\n")

	processed := make(map[coord]bool)
	price := 0

	for row := range len(field) {
		for col := range len(field[0]) {
			_, is_processed := processed[coord{row, col}]
			if !is_processed {
				p := getPlantStats(field, coord{row, col}, &processed)
				fmt.Printf("ID: %s, area: %d, perimeter: %d\n", string(field[row][col]), p.area, p.perimeter)
				price += p.area * p.perimeter
			}
		}
	}
	fmt.Println(price)
}

func getPlantStats(field []string, start coord, processed *map[coord]bool) plant {
	id := field[start.row][start.col]
	area := make(map[coord]bool)
	perimeter := 0

	directions := []direction{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	queue := []coord{start}
	area[start] = true

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		for _, d := range directions {
			next_pos := coord{pos.row + d.v, pos.col + d.h}
			if (next_pos.row < 0) || (next_pos.row >= len(field)) || (next_pos.col < 0) || (next_pos.col >= len(field[0])) {
				perimeter += 1
				continue
			}
			if field[next_pos.row][next_pos.col] != id {
				perimeter += 1
				continue
			}
			_, is_visited := area[next_pos]
			if !is_visited {
				queue = append(queue, next_pos)
				area[next_pos] = true
			}
		}
	}

	for c := range area {
		(*processed)[c] = true
	}
	return plant{len(area), perimeter}
}
