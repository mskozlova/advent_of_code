package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	total_diff := 0
	input, _ := os.ReadFile("2024/day01_pt1/input.txt")
	left_list := make([]int, 0, 0)
	right_list := make([]int, 0, 0)

	for _, line := range strings.Split(string(input), "\n") {
		elements := strings.Fields(line)
		left_element, _ := strconv.Atoi(elements[0])
		right_element, _ := strconv.Atoi(elements[1])
		left_list = append(left_list, left_element)
		right_list = append(right_list, right_element)
	}

	slices.Sort(left_list)
	slices.Sort(right_list)

	for i := range len(left_list) {
		current_diff := int(math.Abs(float64(left_list[i] - right_list[i])))
		total_diff += current_diff
	}

	fmt.Printf("%d", total_diff)
}
