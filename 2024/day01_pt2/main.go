package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func updateCounter(counter *map[int]int, id int) {
	cnt, ok := (*counter)[id]
	if ok {
		(*counter)[id] = cnt + 1
	} else {
		(*counter)[id] = 1
	}
}

func getValue(counter *map[int]int, id int) int {
	cnt, ok := (*counter)[id]
	if ok {
		return cnt
	}
	return 0
}

func main() {
	similarity_score := 0
	input, _ := os.ReadFile("2024/day01_pt2/input.txt")
	left_ids := make(map[int]int)
	right_ids := make(map[int]int)

	for _, line := range strings.Split(string(input), "\n") {
		elements := strings.Fields(line)
		left_element, _ := strconv.Atoi(elements[0])
		right_element, _ := strconv.Atoi(elements[1])

		updateCounter(&left_ids, left_element)
		updateCounter(&right_ids, right_element)
	}

	for id, cnt := range left_ids {
		similarity_score += id * cnt * getValue(&right_ids, id)
	}

	fmt.Printf("%d", similarity_score)
}
