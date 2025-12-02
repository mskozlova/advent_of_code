package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	from int
	to   int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	input_str := string(input)
	intervals := make([]interval, 0, 0)
	sum_invalid_intervals := 0

	for _, s := range strings.Split(input_str, ",") {
		intervals = append(intervals, parseInterval(s))
	}

	for _, interval := range intervals {
		sum_invalid_intervals += findSumInvalidInInterval(interval)
	}
	fmt.Println(sum_invalid_intervals)
}

func parseInterval(s string) interval {
	from, _ := strconv.Atoi(strings.Split(s, "-")[0])
	to, _ := strconv.Atoi(strings.Split(s, "-")[1])
	return interval{from, to}
}

func findSumInvalidInInterval(i interval) int {
	sum_invalid := 0
	for num := i.from; num <= i.to; num++ {
		if isInvalid(num) {
			sum_invalid += num
		}
	}
	return sum_invalid
}

func isInvalid(num int) bool {
	n_digits := len(strconv.Itoa(num))
	if n_digits%2 == 1 {
		return false
	}
	first_half := num / IntPow(10, n_digits/2)
	second_half := num - first_half*IntPow(10, n_digits/2)
	return first_half == second_half
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
