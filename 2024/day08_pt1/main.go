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

func findAntinodes(c1, c2 coord) (coord, coord) {
	diff_rows := c2.row - c1.row
	diff_cols := c2.col - c1.col

	return coord{row: c1.row - diff_rows, col: c1.col - diff_cols}, coord{row: c2.row + diff_rows, col: c2.col + diff_cols}
}

func checkAntinode(c coord, total_rows int, total_cols int) bool {
	if (c.row < 0) || (c.row >= total_rows) {
		return false
	}
	if (c.col < 0) || (c.col >= total_cols) {
		return false
	}
	return true
}

func main() {
	input, _ := os.ReadFile("2024/day08_pt1/input.txt")
	field := strings.Split(string(input), "\n")
	total_rows := len(field)
	total_cols := len(field[0])

	antennas := scanAntennas(field)
	antinodes := make(map[coord]bool)
	for id, coords := range antennas {
		fmt.Printf("Processing antenna %s\n", id)
		for i := range len(coords) {
			for j := range i {
				a1, a2 := findAntinodes(coords[i], coords[j])
				if checkAntinode(a1, total_rows, total_cols) {
					antinodes[a1] = true
				}
				if checkAntinode(a2, total_rows, total_cols) {
					antinodes[a2] = true
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
