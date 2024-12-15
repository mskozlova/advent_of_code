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
	n_seconds    int = 100
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
	input, _ := os.ReadFile("2024/day14_pt1/input.txt")
	rows := strings.Split(string(input), "\n")
	// robots := make([]robot, len(rows))
	robot_positions := make(map[pos]int)

	for i := range len(rows) {
		r := parseRobot(rows[i])
		p := calculatePos(r, n_seconds)
		_, is_occupied := robot_positions[p]
		if !is_occupied {
			robot_positions[p] = 0
		}
		robot_positions[p] += 1
	}
	fmt.Println(calculateSafety(robot_positions))
}

func calculateSafety(r_pos map[pos]int) int {
	total_safety := 1
	q_safety := []int{0, 0, 0, 0}
	for rp, cnt := range r_pos {
		q_idx := -1
		if (rp.x < n_total_cols/2) && (rp.y < n_total_rows/2) {
			q_idx = 0
		} else if (rp.x > n_total_cols/2) && (rp.y < n_total_rows/2) {
			q_idx = 1
		} else if (rp.x < n_total_cols/2) && (rp.y > n_total_rows/2) {
			q_idx = 2
		} else if (rp.x > n_total_cols/2) && (rp.y > n_total_rows/2) {
			// Q1
			q_idx = 3
		}

		if q_idx >= 0 {
			q_safety[q_idx] += cnt
		}
	}

	for _, s := range q_safety {
		total_safety *= s
	}
	return total_safety
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
