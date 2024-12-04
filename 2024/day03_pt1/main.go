package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func performOperation(op []byte) int {
	re_num := regexp.MustCompile(`\d{1,3}`)
	result := 1
	for _, num_b := range re_num.FindAll(op, -1) {
		num, _ := strconv.Atoi(string(num_b))
		result *= num
	}
	return result
}

func main() {
	input, _ := os.ReadFile("2024/day03_pt1/input.txt")

	re_ops := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	total_sum := 0

	for _, op := range re_ops.FindAll(input, -1) {
		total_sum += performOperation(op)
	}

	fmt.Printf("%d\n", total_sum)
}
