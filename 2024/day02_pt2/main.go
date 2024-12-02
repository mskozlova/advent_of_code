package main

import (
	"fmt"
	"math"
	"os"
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

func removeIthElement(array []int, i int) []int {
	array_short := make([]int, i)
	copy(array_short, array[:i])
	return append(array_short, array[i+1:]...)
}

func equalSlices(lhs []int, rhs []int) bool {
	for i := range len(lhs) {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func main() {
	input, _ := os.ReadFile("2024/day02_pt2/input.txt")
	n_safe := 0

	for _, line := range strings.Split(string(input), "\n") {
		line = strings.TrimSpace(line)
		levels := arrayAtoi(strings.Split(line, " "))
		fmt.Println("--", levels, "--")

		for i := range len(levels) {
			levels_trunc := removeIthElement(levels, i)
			fmt.Println(levels_trunc)

			if isSafe(levels_trunc) {
				n_safe += 1
				break
			}
		}
	}
	fmt.Printf("Total safe: %d\n", n_safe)
}
