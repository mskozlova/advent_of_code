package main

import (
	"fmt"
	"os"
	"strings"
)

const ROLL_SYM string = "@"
const MAX_ROLL_NEIGHBOURS int = 3

var neighbours [][]int = [][]int{
	[]int{0, 1},
	[]int{0, -1},
	[]int{1, 1},
	[]int{1, 0},
	[]int{1, -1},
	[]int{-1, 1},
	[]int{-1, 0},
	[]int{-1, -1},
}

func main() {
	input, _ := os.ReadFile("input.txt")

	field := strings.Split(string(input), "\n")
	fmt.Println(countAccessibleRolls(field))
}

func countAccessibleRolls(field []string) int {
	field_height := len(field)
	field_width := len(field[0])
	n_accessible := 0

	for i, row := range field {
		for j, sym := range row {
			if string(sym) != ROLL_SYM {
				continue
			}

			n_roll_neighbours := 0
			for _, d := range neighbours {
				ni := i + d[0]
				nj := j + d[1]

				if ni < 0 || ni >= field_height || nj < 0 || nj >= field_width {
					continue
				}
				if string(field[ni][nj]) == ROLL_SYM {
					n_roll_neighbours += 1
				}
			}
			if n_roll_neighbours <= MAX_ROLL_NEIGHBOURS {
				n_accessible += 1
			}
		}
	}
	return n_accessible
}
