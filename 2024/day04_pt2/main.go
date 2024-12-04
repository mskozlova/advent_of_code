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

func checkPattern(input []string, pattern string, start coords) bool {
	n_rows := len(input)
	n_cols := len(input[0])

	directions := []direction{
		{h: 1, v: 1},
		{h: -1, v: 1},
	}

	matches := true
	for _, d := range directions {
		has0 := false
		has2 := false

		for _, coef := range []int{1, -1} {
			c := coords{
				row: start.row + coef*d.v,
				col: start.col + coef*d.h,
			}
			if (c.row < 0) || (c.row >= n_rows) ||
				(c.col < 0) || (c.col >= n_cols) {
				matches = false
				break
			}
			if input[c.row][c.col] == pattern[0] {
				has0 = true
			}
			if input[c.row][c.col] == pattern[2] {
				has2 = true
			}
		}
		if !has0 || !has2 {
			matches = false
			break
		}
	}
	return matches
}

func main() {
	input, _ := os.ReadFile("2024/day04_pt2/input.txt")
	input_rows := strings.Split(string(input), "\n")
	n_patterns := 0
	pattern := "MAS"

	a_idxs := findAllIndices(input_rows, pattern[1])

	for _, c := range a_idxs {
		if checkPattern(input_rows, pattern, c) {
			n_patterns += 1
		}
	}

	fmt.Println(n_patterns)
}
