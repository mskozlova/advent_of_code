package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FIRST_POS int = 50
	TOTAL_POS int = 100
)

type instruction struct {
	direction string
	steps     int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	instructions := make([]instruction, 0, 0)
	zero_counter := 0

	for _, line := range strings.Split(string(input), "\n") {
		direction := string(line[0])
		steps, _ := strconv.Atoi(line[1:])

		instructions = append(instructions, instruction{direction, steps})
	}

	current_pos := FIRST_POS

	for _, instruction := range instructions {
		sign := 1
		if instruction.direction == "L" {
			sign = -1
		}
		current_pos += sign*instruction.steps + TOTAL_POS
		current_pos = current_pos % TOTAL_POS

		if current_pos == 0 {
			zero_counter += 1
		}
	}

	fmt.Printf("%d", zero_counter)
}
