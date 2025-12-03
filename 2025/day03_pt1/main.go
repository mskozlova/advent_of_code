package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	sum_joltage := 0
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

func getMaxJoltage(battery []int) int {
	max_idx := slices.Index(battery, slices.Max(battery[:len(battery)-1]))
	max_joltage := battery[max_idx]*10 + slices.Max(battery[max_idx+1:])
	return max_joltage
}
