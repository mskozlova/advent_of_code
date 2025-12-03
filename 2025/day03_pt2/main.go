package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const MAX_BATTERIES int = 12

func main() {
	input, _ := os.ReadFile("input.txt")
	sum_joltage := int64(0)
	for _, line := range strings.Split(string(input), "\n") {
		battery := readBattery(line)
		sum_joltage += getMaxJoltage(battery)
	}
	fmt.Println(sum_joltage)
}

func readBattery(s string) []int {
	battery := make([]int, 0, 0)
	for _, sym := range s {
		digit, _ := strconv.Atoi(string(sym))
		battery = append(battery, digit)
	}
	return battery
}

func getMaxJoltage(battery []int) int64 {
	max_joltage := int64(0)
	current_idx := 0

	for i := range MAX_BATTERIES {
		required_suffix_len := MAX_BATTERIES - i - 1
		for digit := 9; digit >= 0; digit-- {
			pos := slices.Index(battery[current_idx:], digit)

			if pos < 0 {
				continue
			}

			pos += current_idx
			if len(battery)-pos-1 >= required_suffix_len {
				max_joltage = max_joltage*10 + int64(digit)
				current_idx = pos + 1
				break
			}
		}
	}
	return max_joltage
}
