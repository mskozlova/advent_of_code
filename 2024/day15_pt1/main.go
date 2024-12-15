package main

import (
	"fmt"
	"os"
	"strings"
)

type direction struct {
	v, h int
}

type position struct {
	row, col int
}

type field struct {
	field []string
	boxes map[position]bool
}

func main() {
	file_name := "2024/day15_pt1/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Printf("No such file %s\n", file_name)
		return
	}
	start, fld, directions := parseInput(string(input))
	fmt.Println(start)
	fmt.Println("boxes", fld.boxes)
	fmt.Println("field\n", fld.field)
	end_pos := runRobot(start, &fld, directions)
	printField(end_pos, &fld)
	fmt.Println(getBoxScore(&fld))
}

func getBoxScore(fld *field) int {
	total_score := 0
	for box := range (*fld).boxes {
		score := 100*box.row + box.col
		total_score += score
	}
	return total_score
}

func printField(pos position, fld *field) {
	print := make([]string, len((*fld).field))
	for row := range len((*fld).field) {
		row_str := ""
		for col := range len((*fld).field[0]) {
			_, is_box := (*fld).boxes[position{row, col}]
			if string((*fld).field[row][col]) == "#" {
				row_str += "#"
			} else if (pos.row == row) && (pos.col == col) {
				row_str += "@"
			} else if is_box {
				row_str += "O"
			} else {
				row_str += "."
			}
		}
		print[row] = row_str
	}
	fmt.Println(strings.Join(print, "\n"))
}

func runRobot(start position, fld *field, directions []direction) position {
	pos := start
	for i, dir := range directions {
		can_do := checkStep(pos, fld, dir)
		if can_do {
			pos = makeStep(pos, fld, dir, false)
		}
		fmt.Printf("Step #%d (%d, %d) - %t - pos after (%d, %d)\n", i, dir.v, dir.h, can_do, pos.row, pos.col)
	}
	return pos
}

func checkStep(pos position, fld *field, dir direction) bool {
	next_pos := position{row: pos.row + dir.v, col: pos.col + dir.h}

	if (next_pos.row < 0) || (next_pos.row >= len((*fld).field)) {
		return false
	}
	if (next_pos.col < 0) || (next_pos.col >= len((*fld).field[0])) {
		return false
	}
	if string((*fld).field[next_pos.row][next_pos.col]) == "#" {
		return false
	}

	_, is_box := (*fld).boxes[next_pos]
	if is_box {
		return checkStep(next_pos, fld, dir)
	}

	return true
}

func makeStep(pos position, fld *field, dir direction, as_box bool) position {
	next_pos := position{row: pos.row + dir.v, col: pos.col + dir.h}

	_, is_box := (*fld).boxes[next_pos]
	if is_box {
		if !as_box {
			delete((*fld).boxes, next_pos)
		}
		makeStep(next_pos, fld, dir, true)
	} else if as_box && !is_box {
		(*fld).boxes[next_pos] = true
	}
	return next_pos
}

func parseInput(input string) (position, field, []direction) {
	parts := strings.Split(input, "\n\n")
	pos, fld := parseField(parts[0])
	return pos, fld, parseDirections(parts[1])
}

func parseDirections(dir_raw string) []direction {
	directions := make([]direction, len(dir_raw))
	for i, sym := range dir_raw {
		if string(sym) == "<" {
			directions[i] = direction{v: 0, h: -1}
		} else if string(sym) == ">" {
			directions[i] = direction{v: 0, h: 1}
		} else if string(sym) == "^" {
			directions[i] = direction{v: -1, h: 0}
		} else if string(sym) == "v" {
			directions[i] = direction{v: 1, h: 0}
		}
	}
	return directions
}

func parseField(field_raw string) (position, field) {
	rows_raw := strings.Split(field_raw, "\n")
	fld := make([]string, len(rows_raw))
	start := position{}
	boxes := make(map[position]bool)

	for r, row_raw := range rows_raw {
		row := ""
		for c, sym := range row_raw {
			if string(sym) == "@" {
				start.row = r
				start.col = c
				row += "."
			} else if string(sym) == "O" {
				boxes[position{r, c}] = true
				row += "."
			} else if string(sym) == "#" {
				row += "#"
			} else {
				row += "."
			}
		}
		fld[r] = row
	}
	return start, field{fld, boxes}
}
