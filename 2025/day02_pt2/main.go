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
	num_str := strconv.Itoa(num)
	n_digits := len(num_str)

	for i := 1; i < n_digits; i++ {
		if n_digits%i == 0 {
			prefix := num_str[:i]
			if isRepeatedPrefix(num_str, prefix) {
				return true
			}
		}
	}
	return false
}

func isRepeatedPrefix(s string, prefix string) bool {
	for i := 0; i < len(s); i += len(prefix) {
		if s[i:i+len(prefix)] != prefix {
			return false
		}
	}
	return true
}
