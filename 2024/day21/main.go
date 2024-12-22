package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//     | 0 | A |
//     +---+---+

//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+

const (
	inf      int = 1000000000000
	n_robots int = 25
)

type coord struct {
	row, col int
}

type direction struct {
	v, h int
}

var code2coord = map[string]coord{
	"7": {0, 0},
	"8": {0, 1},
	"9": {0, 2},
	"4": {1, 0},
	"5": {1, 1},
	"6": {1, 2},
	"1": {2, 0},
	"2": {2, 1},
	"3": {2, 2},
	"0": {3, 1},
	"A": {3, 2},
}

var digit_coords = map[coord]bool{
	{0, 0}: true,
	{0, 1}: true,
	{0, 2}: true,
	{1, 0}: true,
	{1, 1}: true,
	{1, 2}: true,
	{2, 0}: true,
	{2, 1}: true,
	{2, 2}: true,
	{3, 1}: true,
	{3, 2}: true,
}

var arrow2coord = map[string]coord{
	"^": {0, 1},
	"A": {0, 2},
	"<": {1, 0},
	"v": {1, 1},
	">": {1, 2},
}

var arrow_coords = map[coord]bool{
	{0, 1}: true,
	{0, 2}: true,
	{1, 0}: true,
	{1, 1}: true,
	{1, 2}: true,
}

var direction2arrow = map[direction]string{
	{1, 0}:  "v",
	{-1, 0}: "^",
	{0, 1}:  ">",
	{0, -1}: "<",
}

type status struct {
	from  coord
	to    coord
	level int
}

func main() {
	file_name := "2024/day21/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	precalc := make(map[status]int)
	codes := strings.Split(string(input), "\n")
	complexity := 0

	for _, code := range codes {
		fmt.Printf("--- CODE %s ----", code)
		code_parsed := parseCode(code)
		full_shortest_path := 0

		for i := range len(code_parsed) {
			from := coord{}
			if i == 0 {
				from = code2coord["A"]
			} else {
				from = code_parsed[i-1]
			}
			to := code_parsed[i]
			moves := makeMove(from, to, &digit_coords)
			shortest_path := inf
			fmt.Println("from:", from, "to:", to)
			for _, move := range moves {
				fmt.Println("Digit move:", move)
				shortest_path = min(shortest_path, getShortestPathLength(move, 0, n_robots, &precalc))
			}
			full_shortest_path += shortest_path
		}

		fmt.Println(code, "--->", full_shortest_path)
		complexity += full_shortest_path * getNumber(code)
	}
	fmt.Println("Complexity:", complexity)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func parseCode(code string) []coord {
	coords := make([]coord, len(code))
	for i, sym := range code {
		coords[i] = code2coord[string(sym)]
	}
	return coords
}

func getNumber(code string) int {
	num, err := strconv.Atoi(code[:len(code)-1])
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return num
}

func makeMove(from coord, to coord, all_coords *map[coord]bool) [][]direction {
	all_paths := make([][]direction, 0)
	d := direction{v: 0, h: 0}
	if from.row > to.row {
		d.v = -1
	} else if from.row < to.row {
		d.v = 1
	}
	if from.col > to.col {
		d.h = -1
	} else if from.col < to.col {
		d.h = 1
	}

	shortest_path := abs(from.row-to.row) + abs(from.col-to.col)
	path := make([]direction, 0)
	getMoves(from, to, shortest_path, all_coords, d, &path, &all_paths)
	return all_paths
}

func getMoves(pos, to coord, max_steps int, all_coords *map[coord]bool, d direction, path *[]direction, paths *[][]direction) {
	if (pos.row == to.row) && (pos.col == to.col) {
		// we're here
		final_path := make([]direction, len(*path))
		copy(final_path, *path)
		*paths = append(*paths, final_path)
		return
	}

	if len(*path) == max_steps {
		return
	}

	next_positions := make([]coord, 0)
	l := len(*path)

	if d.h != 0 {
		next_pos := coord{pos.row, pos.col + d.h}
		_, is_allowed := (*all_coords)[next_pos]
		if is_allowed {
			next_positions = append(next_positions, next_pos)
			if is_allowed {
				*path = append(*path, direction{0, d.h})
				getMoves(next_pos, to, max_steps, all_coords, d, path, paths)
				*path = (*path)[:l]
			}
		}
	}

	if d.v != 0 {
		next_pos := coord{pos.row + d.v, pos.col}
		_, is_allowed := (*all_coords)[next_pos]
		if is_allowed {
			*path = append(*path, direction{d.v, 0})
			getMoves(next_pos, to, max_steps, all_coords, d, path, paths)
			*path = (*path)[:l]
		}
	}
}

func getShortestPathLength(move []direction, level, max_level int, precalc *map[status]int) int {
	defer fmt.Println(*precalc)
	if level == max_level {
		return len(move) + 1
	}

	min_path_length := 0
	var from, to coord

	for i := range len(move) + 1 {
		if i == 0 {
			from = arrow2coord["A"]
		} else {
			from = arrow2coord[direction2arrow[move[i-1]]]
		}
		if i == len(move) {
			to = arrow2coord["A"]
		} else {
			to = arrow2coord[direction2arrow[move[i]]]
		}

		min_path, is_calculated := (*precalc)[status{from, to, level}]
		if is_calculated {
			min_path_length += min_path
			continue
		}

		arrow_moves := makeMove(from, to, &arrow_coords)
		min_path = inf
		for _, arrow_move := range arrow_moves {
			min_path = min(
				min_path,
				getShortestPathLength(arrow_move, level+1, max_level, precalc),
			)
		}

		min_path_length += min_path
		(*precalc)[status{from, to, level}] = min_path
	}
	return min_path_length
}
