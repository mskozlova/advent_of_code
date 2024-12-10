package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	row, col int
}

func main() {
	input, _ := os.ReadFile("2024/day10_pt2/input.txt")
	field := parseField(string(input))
	starts := findStarts(field)
	total_score := 0

	for _, start := range starts {
		score := countTrails(field, start)
		total_score += score

		fmt.Printf("Start (%d, %d) - score %d\n", start.row, start.col, score)
	}

	fmt.Println(total_score)
}

func parseField(input string) [][]int {
	rows := strings.Split(input, "\n")
	field := make([][]int, len(rows))

	for i, row := range rows {
		row_num := make([]int, len(row))
		for j, sym := range row {
			num, _ := strconv.Atoi(string(sym))
			row_num[j] = num
		}
		field[i] = row_num
	}

	return field
}

func findStarts(field [][]int) []coord {
	starts := make([]coord, 0)
	for row := range len(field) {
		for col := range len(field[0]) {
			if field[row][col] == 0 {
				starts = append(starts, coord{row: row, col: col})
			}
		}
	}
	return starts
}

func countTrails(field [][]int, start coord) int {
	n_rows := len(field)
	n_cols := len(field[0])

	queue := []coord{start}
	n_trails := 0

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		value := field[c.row][c.col]

		if value == 9 {
			n_trails += 1
			continue
		}

		next_pos := []coord{
			{c.row - 1, c.col},
			{c.row + 1, c.col},
			{c.row, c.col - 1},
			{c.row, c.col + 1},
		}

		for _, n := range next_pos {
			if (n.row < 0) || (n.row >= n_rows) || (n.col < 0) || (n.col >= n_cols) {
				continue
			}

			if field[n.row][n.col] == value+1 {
				queue = append(queue, n)
			}
		}
	}

	return n_trails
}
