package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	n_total_rows int = 103
	n_total_cols int = 101
)

type pos struct {
	x, y int
}
type velocity struct {
	x, y int
}
type robot struct {
	p pos
	v velocity
}

func main() {
	input, _ := os.ReadFile("2024/day14_pt2/input.txt")
	rows := strings.Split(string(input), "\n")
	robots := make([]robot, len(rows))
	for i := range len(rows) {
		r := parseRobot(rows[i])
		robots[i] = r
	}

	max_symmetrical := 0
	max_symmetrical_seconds := 0

	for n_seconds := range 200000 {
		robot_positions := calculateMap(robots, n_seconds)
		n_symmetrical := checkMap(robot_positions)
		if n_symmetrical > max_symmetrical {
			max_symmetrical = n_symmetrical
			max_symmetrical_seconds = n_seconds
		}
		fmt.Printf("# Symmetrical after %d seconds: %d out of %d\n", n_seconds, n_symmetrical, len(robot_positions))
	}

	fmt.Println(max_symmetrical, max_symmetrical_seconds)
	robot_positions := calculateMap(robots, max_symmetrical_seconds)
	fmt.Println(drawMap(robot_positions))
}

func drawMap(robot_positions map[pos]int) string {
	full_map := make([]string, n_total_rows)
	for row := range n_total_rows {
		str := ""
		for col := range n_total_cols {
			_, exists := robot_positions[pos{col, row}]
			if exists {
				str += "*"
			} else {
				str += "."
			}
		}
		full_map[row] = str
	}
	return strings.Join(full_map, "\n")
}

func checkMap(robot_positions map[pos]int) int {
	n_symmetrical := 0
	for r := range robot_positions {
		_, exists := robot_positions[pos{n_total_cols - r.x, r.y}]
		if exists {
			n_symmetrical += 1
		}
	}
	return n_symmetrical / 2
}

func calculateMap(robots []robot, n_seconds int) map[pos]int {
	robot_positions := make(map[pos]int)
	for i := range len(robots) {
		r := robots[i]
		p := calculatePos(r, n_seconds)
		_, is_occupied := robot_positions[p]
		if !is_occupied {
			robot_positions[p] = 0
		}
		robot_positions[p] += 1
	}
	return robot_positions
}

func calculatePos(r robot, secs int) pos {
	next_x := (r.p.x + secs*r.v.x) % n_total_cols
	next_y := (r.p.y + secs*r.v.y) % n_total_rows

	if next_x < 0 {
		next_x += n_total_cols
	}
	if next_y < 0 {
		next_y += n_total_rows
	}
	return pos{next_x, next_y}
}

func parseRobot(row string) robot {
	elems := strings.Split(row, " ")

	p_raw := strings.Split(elems[0][2:], ",")
	px, _ := strconv.Atoi(p_raw[0])
	py, _ := strconv.Atoi(p_raw[1])

	v_raw := strings.Split(elems[1][2:], ",")
	vx, _ := strconv.Atoi(v_raw[0])
	vy, _ := strconv.Atoi(v_raw[1])

	return robot{pos{px, py}, velocity{vx, vy}}
}
