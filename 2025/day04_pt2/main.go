package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	row int
	col int
}

type size struct {
	height int
	width  int
}

const ROLL_SYM string = "@"
const NO_ROLL_COUNTER int = -1
const MAX_ROLL_NEIGHBOURS int = 3

var neighbour_deltas [][]int = [][]int{
	{0, 1},
	{0, -1},
	{1, 1},
	{1, 0},
	{1, -1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

func main() {
	input, _ := os.ReadFile("input.txt")

	field := strings.Split(string(input), "\n")
	counter := countNeighbours(field)
	fmt.Println(countRemovals(counter))
}

func countNeighbours(field []string) [][]int {
	field_size := size{len(field), len(field[0])}
	neighour_counter := make([][]int, 0, 0)

	for i, row := range field {
		neighbour_row := make([]int, 0, 0)
		for j, sym := range row {
			if string(sym) != ROLL_SYM {
				neighbour_row = append(neighbour_row, NO_ROLL_COUNTER)
				continue
			}

			n_roll_neighbours := 0
			for _, n := range getNeighbours(position{i, j}, field_size) {
				if string(field[n.row][n.col]) == ROLL_SYM {
					n_roll_neighbours += 1
				}
			}
			neighbour_row = append(neighbour_row, n_roll_neighbours)
		}
		neighour_counter = append(neighour_counter, neighbour_row)
	}
	return neighour_counter
}

func countRemovals(neighour_counter [][]int) int {
	n_removals := 0
	field_size := size{len(neighour_counter), len(neighour_counter[0])}

	for step := 0; step < field_size.width*field_size.height; step++ {
		round_removals := 0
		// if counter <= 3 - remove roll, count removal, update neighbours
		for row := range field_size.height {
			for col := range field_size.width {
				if neighour_counter[row][col] >= 0 && neighour_counter[row][col] <= MAX_ROLL_NEIGHBOURS {
					round_removals += 1
					neighour_counter[row][col] = NO_ROLL_COUNTER
					updateNeighboursAroundRemoval(neighour_counter, position{row, col}, field_size)
				}
			}
		}

		if round_removals == 0 {
			break
		}
		n_removals += round_removals
	}
	return n_removals
}

func getNeighbours(pos position, s size) []position {
	neighbours := make([]position, 0, 0)
	for _, d := range neighbour_deltas {
		n_row := pos.row + d[0]
		n_col := pos.col + d[1]

		if n_row < 0 || n_row >= s.height || n_col < 0 || n_col >= s.width {
			continue
		}
		neighbours = append(neighbours, position{n_row, n_col})
	}
	return neighbours
}

func updateNeighboursAroundRemoval(neighour_counter [][]int, pos position, s size) {
	for _, n := range getNeighbours(pos, s) {
		if neighour_counter[n.row][n.col] > 0 {
			neighour_counter[n.row][n.col] -= 1
		}
	}
}
