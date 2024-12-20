package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	cheat_steps int = 20
	n_saved         = 100
)

type coord struct {
	row, col int
}

type direction struct {
	v, h int
}

type step struct {
	c coord
	n int
}

func main() {
	file_name := "2024/day20_pt2/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	start, end, field := parseField(string(input))
	n_steps_normal, _ := countStepsNormal(field, start)[end]
	fmt.Println("# steps without cheats: ", n_steps_normal)
	cheats := countCheats(field, start, end)

	printCheats(&cheats, n_steps_normal)

}

func parseField(input string) (coord, coord, []string) {
	field := strings.Split(input, "\n")
	start := coord{}
	end := coord{}

	for i, row := range field {
		for j, sym := range row {
			if string(sym) == "S" {
				start.row = i
				start.col = j
			}

			if string(sym) == "E" {
				end.row = i
				end.col = j
			}
		}
	}

	return start, end, field
}

func countStepsNormal(field []string, from coord) map[coord]int {
	queue := []step{{from, 0}}
	visited := make(map[coord]bool)
	visited[queue[0].c] = true
	distances := make(map[coord]int)

	directions := []direction{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		pos := queue[0].c
		n_steps := queue[0].n
		queue = queue[1:]

		distances[pos] = n_steps

		for _, d := range directions {
			next_coord := coord{pos.row + d.v, pos.col + d.h}

			if (next_coord.row < 0) || (next_coord.row >= len(field)) || (next_coord.col < 0) || (next_coord.col >= len(field[0])) {
				continue
			}
			if string(field[next_coord.row][next_coord.col]) == "#" {
				continue
			}
			_, is_visited := visited[next_coord]
			if is_visited {
				continue
			}
			queue = append(queue, step{next_coord, n_steps + 1})
			visited[next_coord] = true
		}
	}

	return distances
}

func findAllCheats(field []string, pos coord) map[coord]int {
	cheats := make(map[coord]int)
	visited := make(map[coord]bool)
	queue := []step{{pos, 0}}
	visited[pos] = true

	for len(queue) > 0 {
		c := queue[0].c
		n := queue[0].n
		queue = queue[1:]

		if n > cheat_steps {
			continue
		}

		if string(field[c.row][c.col]) != "#" {
			if c != pos {
				cheats[c] = n
			}
		}

		directions := []direction{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, d := range directions {
			next_c := coord{c.row + d.v, c.col + d.h}

			if (next_c.row < 0) || (next_c.row >= len(field)) || (next_c.col < 0) || (next_c.col >= len(field[0])) {
				continue
			}
			_, is_visited := visited[next_c]
			if is_visited {
				continue
			}
			visited[next_c] = true
			queue = append(queue, step{next_c, n + 1})
		}
	}
	return cheats
}

func countCheats(field []string, start, end coord) map[int]int {
	map_to_start := countStepsNormal(field, start)
	map_to_end := countStepsNormal(field, end)
	cheat_secs := make(map[int]int)

	for row := range len(field) {
		for col := range len(field[0]) {
			pos := coord{row, col}
			steps_to_here, path_exists := map_to_start[pos]
			if !path_exists {
				continue
			}

			cheats := findAllCheats(field, pos)
			for cheat, cheat_steps := range cheats {
				steps_to_finish, path_exists := map_to_end[cheat]
				if !path_exists {
					continue
				}
				total_steps := steps_to_here + cheat_steps + steps_to_finish
				value, has_value := cheat_secs[total_steps]
				if !has_value {
					value = 0
				}
				cheat_secs[total_steps] = value + 1

				if total_steps == 84-66 {
					fmt.Printf("Cheat %d, %d -> %d, %d\n", row, col, cheat.row, cheat.col)
					fmt.Printf("From start %d, to end %d, cheat %d\n\n", steps_to_here, steps_to_finish, cheat_steps)
				}
			}
		}
	}
	return cheat_secs
}

func printCheats(cheats *map[int]int, n_steps_normal int) {
	n_good_cheats := 0
	steps_sorted := make([]int, len(*cheats))
	i := 0
	for steps := range *cheats {
		steps_sorted[i] = steps
		i += 1
	}
	slices.Sort(steps_sorted)
	slices.Reverse(steps_sorted)
	for _, steps := range steps_sorted {
		cnt, _ := (*cheats)[steps]
		if steps < n_steps_normal {
			fmt.Printf("%d cheats save %d seconds\n", cnt, n_steps_normal-steps)
		}
		if n_steps_normal-steps >= n_saved {
			n_good_cheats += cnt
		}
	}
	fmt.Printf("Cheats saving >=%d seconds: %d\n", n_saved, n_good_cheats)
}
