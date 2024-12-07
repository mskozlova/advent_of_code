package main

import (
	"fmt"
	"os"
	"strings"
)

type coord struct {
	row, col int
}

type position struct {
	row, col     int
	dir_v, dir_h int
}

func findStart(field []string, obstacle string, empty string) position {
	for row := range len(field) {
		for col := range len(field[0]) {
			sym := string(field[row][col])
			if (sym != obstacle) && (sym != empty) {
				dir_h, dir_v := 0, 0
				if sym == "^" {
					dir_v = -1
				} else if sym == ">" {
					dir_h = 1
				} else if sym == "v" {
					dir_v = 1
				} else if sym == "<" {
					dir_h = -1
				}
				return position{
					row: row, col: col,
					dir_v: dir_v, dir_h: dir_h,
				}
			}
		}
	}
	return position{}
}

func turn(pos position) position {
	if pos.dir_h != 0 {
		return position{
			row:   pos.row,
			col:   pos.col,
			dir_v: pos.dir_h,
			dir_h: 0,
		}
	}

	if pos.dir_v != 0 {
		return position{
			row:   pos.row,
			col:   pos.col,
			dir_v: 0,
			dir_h: -pos.dir_v,
		}
	}

	return position{}
}

func doStep(field []string, pos position, obstacle string) (bool, position) {
	rows := len(field)
	cols := len(field[0])

	next_row := pos.row + pos.dir_v
	next_col := pos.col + pos.dir_h

	if (next_row < 0) || (next_row >= rows) || (next_col < 0) || (next_col >= cols) {
		return true, position{}
	}

	if string(field[next_row][next_col]) == obstacle {
		return false, turn(pos)
	}

	return false, position{
		row:   next_row,
		col:   next_col,
		dir_v: pos.dir_v,
		dir_h: pos.dir_h,
	}
}

func countVisited(field []string, start position, obstacle string) int {
	visited := make(map[coord]bool)

	do_exit := false
	pos := start

	for !do_exit {
		visited[coord{pos.row, pos.col}] = true
		do_exit, pos = doStep(field, pos, obstacle)
		fmt.Println(do_exit, pos)
	}

	return len(visited)
}

func main() {
	input, _ := os.ReadFile("2024/day06_pt1/input.txt")
	field := strings.Split(string(input), "\n")
	obstacle := "#"
	empty := "."

	start := findStart(field, obstacle, empty)
	n_visited := countVisited(field, start, obstacle)
	fmt.Println(n_visited)
}
