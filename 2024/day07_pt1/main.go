package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	expected int
	nums     []int
}

func parseEquations(raw string) equation {
	parts := strings.Split(raw, ":")
	res, _ := strconv.Atoi(parts[0])
	nums_str := strings.Split(strings.TrimSpace(parts[1]), " ")
	nums := make([]int, len(nums_str))

	for i, num_str := range nums_str {
		num, _ := strconv.Atoi(num_str)
		nums[i] = num
	}

	return equation{
		expected: res,
		nums:     nums,
	}
}

func evaluateEquation(eq equation, idx int, current_val int) bool {
	if (current_val == eq.expected) && (idx == len(eq.nums)) {
		return true
	}
	if current_val > eq.expected {
		return false
	}
	if idx == len(eq.nums) {
		return false
	}

	next_val := eq.nums[idx]

	return evaluateEquation(eq, idx+1, current_val*next_val) ||
		evaluateEquation(eq, idx+1, current_val+next_val)
}

func main() {
	input, _ := os.ReadFile("2024/day07_pt1/input.txt")
	equations_raw := strings.Split(string(input), "\n")
	total_sum := 0

	for i, raw := range equations_raw {
		eq := parseEquations(raw)
		is_ok := evaluateEquation(eq, 1, eq.nums[0])
		if is_ok {
			total_sum += eq.expected
		}
		fmt.Printf("Eq #%d: %s -> %t\n", i, raw, is_ok)
	}

	fmt.Println(total_sum)
}
