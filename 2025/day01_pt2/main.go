package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FIRST_POSITION  int = 50
	TOTAL_POSITIONS int = 100
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

	current_pos := FIRST_POSITION

	for _, instruction := range instructions {
		sign := 1
		if instruction.direction == "L" {
			sign = -1
		}

		rounds := instruction.steps / TOTAL_POSITIONS
		zero_counter += rounds
		leftover_steps := instruction.steps - TOTAL_POSITIONS*rounds

		// current pos always in [0, TOTAL_POS) interval
		// leftover steps is always positive and < TOTAL_POSITIONS
		next_pos := current_pos + sign*leftover_steps

		if (current_pos == 0) && (next_pos < 0) {
			next_pos += TOTAL_POSITIONS
		} else if (current_pos == 0) && (next_pos < 0) {
			continue
		} else if next_pos >= TOTAL_POSITIONS {
			zero_counter += 1
			next_pos -= TOTAL_POSITIONS
		} else if next_pos <= 0 {
			zero_counter += 1
			next_pos += TOTAL_POSITIONS
		}
		fmt.Println(instruction)
		fmt.Printf(
			"prev pos %d, next pos %d, rounds %d, zero_counter %d\n",
			current_pos, next_pos, rounds, zero_counter,
		)
		current_pos = next_pos % TOTAL_POSITIONS
	}

	fmt.Printf("%d", zero_counter)
}
