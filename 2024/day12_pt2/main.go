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

type border struct {
	c coord
	d direction
}

func main() {
	input, _ := os.ReadFile("2024/day12_pt2/input.txt")
	field := strings.Split(string(input), "\n")

	processed := make(map[coord]bool)
	price := 0

	for row := range len(field) {
		for col := range len(field[0]) {
			_, is_processed := processed[coord{row, col}]
			if !is_processed {
				area, perimeter_h, perimeter_v := getArea(field, coord{row, col}, &processed)
				sides := getSides(perimeter_h, direction{v: 0, h: -1})
				sides += getSides(perimeter_v, direction{v: -1, h: 0})

				fmt.Printf("ID: %s, area: %d, perimeter: %d\n", string(field[row][col]), len(area), sides)
				fmt.Println(perimeter_h, "\n", perimeter_v)
				price += len(area) * sides
			}
		}
	}
	fmt.Println(price)
}

func getArea(field []string, start coord, processed *map[coord]bool) ([]coord, map[border]bool, map[border]bool) {
	id := field[start.row][start.col]
	area := make(map[coord]bool)
	perimeter_v := make(map[border]bool)
	perimeter_h := make(map[border]bool)

	directions := []direction{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	queue := []coord{start}
	area[start] = true

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		for _, d := range directions {
			next_pos := coord{pos.row + d.v, pos.col + d.h}
			is_border := false
			if (next_pos.row < 0) || (next_pos.row >= len(field)) || (next_pos.col < 0) || (next_pos.col >= len(field[0])) {
				is_border = true
			} else if field[next_pos.row][next_pos.col] != id {
				is_border = true
			}
			if is_border && (d.h != 0) {
				perimeter_v[border{next_pos, d}] = true
			} else if is_border {
				perimeter_h[border{next_pos, d}] = true
			}

			if is_border {
				continue
			}

			_, is_visited := area[next_pos]
			if !is_visited {
				queue = append(queue, next_pos)
				area[next_pos] = true
			}
		}
	}

	area_list := make([]coord, 0)
	for c := range area {
		(*processed)[c] = true
		area_list = append(area_list, c)
	}
	return area_list, perimeter_h, perimeter_v
}

func getSides(perimeter map[border]bool, axis direction) int {
	sides := 0
	for tile := range perimeter {
		prev_tile := coord{tile.c.row + axis.v, tile.c.col + axis.h}
		_, exists := perimeter[border{prev_tile, tile.d}]
		if !exists {
			sides += 1
		}
	}
	return sides
}
