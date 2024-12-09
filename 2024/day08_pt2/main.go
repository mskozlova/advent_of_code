package main

import (
	"fmt"
	"os"
	"strings"
)

type coord struct {
	row, col int
}

func scanAntennas(field []string) map[string][]coord {
	antennas := make(map[string][]coord)

	for row := range len(field) {
		for col := range len(field[0]) {
			if string(field[row][col]) == "." {
				continue
			}

			coords, exists := antennas[string(field[row][col])]
			if exists {
				coords = append(coords, coord{row: row, col: col})
			} else {
				coords = []coord{{row: row, col: col}}
			}
			antennas[string(field[row][col])] = coords
		}
	}
	return antennas
}

func findAntinodes(c1, c2 coord, total_rows int, total_cols int) []coord {
	coords := make([]coord, 0)

	diff_rows := c2.row - c1.row
	diff_cols := c2.col - c1.col

	is_in := true
	i := 0
	for is_in {
		coords = append(coords, coord{row: c1.row - i*diff_rows, col: c1.col - i*diff_cols})
		i += 1

		if (c1.row-i*diff_rows < 0) || (c1.row-i*diff_rows >= total_rows) {
			is_in = false
		}
		if (c1.col-i*diff_cols < 0) || (c1.col-i*diff_cols >= total_cols) {
			is_in = false
		}
	}

	is_in = true
	i = 1
	for is_in {
		coords = append(coords, coord{row: c1.row + i*diff_rows, col: c1.col + i*diff_cols})
		i += 1

		if (c1.row+i*diff_rows < 0) || (c1.row+i*diff_rows >= total_rows) {
			is_in = false
		}
		if (c1.col+i*diff_cols < 0) || (c1.col+i*diff_cols >= total_cols) {
			is_in = false
		}
	}

	return coords
}

func main() {
	input, _ := os.ReadFile("2024/day08_pt2/input.txt")
	field := strings.Split(string(input), "\n")
	total_rows := len(field)
	total_cols := len(field[0])

	antennas := scanAntennas(field)
	antinodes := make(map[coord]bool)
	for id, coords := range antennas {
		fmt.Printf("Processing antenna %s\n", id)
		for i := range len(coords) {
			for j := range i {
				as := findAntinodes(coords[i], coords[j], total_rows, total_cols)
				for _, a := range as {
					antinodes[a] = true
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
