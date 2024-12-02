package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func arrayAtoi(array []string) []int {
	var numbers = make([]int, len(array))
	for idx, e_str := range array {
		num, err := strconv.Atoi(e_str)
		if err != nil {
			panic(err)
		}
		numbers[idx] = num
	}
	return numbers
}

func isSafe(levels []int) bool {
	if len(levels) == 1 {
		return true
	}
	is_ascending := levels[1] >= levels[0]

	for i := 1; i < len(levels); i++ {
		if is_ascending && (levels[i-1] > levels[i]) {
			return false
		}
		if !is_ascending && (levels[i-1] <= levels[i]) {
			return false
		}
		diff := int(math.Abs(float64(levels[i] - levels[i-1])))
		if (diff == 0) || (diff > 3) {
			return false
		}
	}
	return true
}

func equalSlices(lhs []int, rhs []int) bool {
	for i := range len(lhs) {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func checkIncreasing(levels []int) bool {
	sorted_levels := make([]int, len(levels))
	copy(sorted_levels, levels)
	slices.Sort(sorted_levels)

	return equalSlices(sorted_levels, levels)
}

func checkDecreasing(levels []int) bool {
	sorted_levels := make([]int, len(levels))
	copy(sorted_levels, levels)

	slices.Sort(sorted_levels)
	slices.Reverse(sorted_levels)

	return equalSlices(sorted_levels, levels)
}

func main() {
	input, _ := os.ReadFile("2024/day02_pt1/input.txt")
	n_safe := 0

	for idx, line := range strings.Split(string(input), "\n") {
		line = strings.TrimSpace(line)
		levels := arrayAtoi(strings.Split(line, " "))
		if isSafe(levels) {
			// debug:
			if !checkIncreasing(levels) && !checkDecreasing(levels) {
				fmt.Println(idx, levels)
			}
			n_safe += 1
			// fmt.Printf("Line #%d is safe: %s\n", idx, line)
		}
	}
	fmt.Printf("%d\n", n_safe)
}
