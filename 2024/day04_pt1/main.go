package main

import (
	"fmt"
	"os"
	"strings"
)

type coords struct {
	row int
	col int
}

type direction struct {
	h int
	v int
}

func findAllIndices(input []string, sym byte) []coords {
	idxs := make([]coords, 0)
	for row := range len(input) {
		for col := range len(input[0]) {
			if input[row][col] == sym {
				idxs = append(idxs, coords{row: row, col: col})
			}
		}
	}
	return idxs
}

func countPatterns(input []string, pattern string, start coords) int {
	n_patterns := 0
	n_rows := len(input)
	n_cols := len(input[0])

	directions := []direction{
		{h: 1, v: -1},
		{h: 1, v: 0},
		{h: 1, v: 1},
		{h: 0, v: -1},
		{h: 0, v: 1},
		{h: -1, v: -1},
		{h: -1, v: 0},
		{h: -1, v: 1},
	}

	for _, d := range directions {
		matches := true
		for i := range len(pattern) {
			c := coords{
				row: start.row + d.v*i,
				col: start.col + d.h*i,
			}
			if (c.row < 0) || (c.row >= n_rows) ||
				(c.col < 0) || (c.col >= n_cols) {
				matches = false
				break
			}
			if input[c.row][c.col] != pattern[i] {
				matches = false
				break
			}
		}
		if matches {
			n_patterns += 1
		}
	}
	return n_patterns
}

func main() {
	input, _ := os.ReadFile("2024/day04_pt1/input.txt")
	input_rows := strings.Split(string(input), "\n")
	n_patterns := 0
	pattern := "XMAS"

	x_idxs := findAllIndices(input_rows, pattern[0])

	for _, c := range x_idxs {
		n_patterns += countPatterns(input_rows, pattern, c)
	}

	fmt.Println(n_patterns)
}
