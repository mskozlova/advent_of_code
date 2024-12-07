package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type order struct {
	before int
	after  int
}

func parseOrder(order_input string) map[order]bool {
	rows := strings.Split(strings.TrimSpace(order_input), "\n")
	result := make(map[order]bool)

	for _, row := range rows {
		elements := strings.Split(row, "|")
		before, _ := strconv.Atoi(elements[0])
		after, _ := strconv.Atoi(elements[1])
		result[order{before: before, after: after}] = true
	}
	return result
}

func parsePages(page_input string) [][]int {
	rows := strings.Split(strings.TrimSpace(page_input), "\n")
	result := make([][]int, len(rows))

	for i, row := range rows {
		page := strings.Split(row, ",")
		rule := make([]int, len(page))
		for j, id_str := range page {
			id_num, _ := strconv.Atoi(id_str)
			rule[j] = id_num
		}
		result[i] = rule
	}
	return result
}

func checkPage(page []int, rules map[order]bool) bool {
	for i := 0; i < len(page); i++ {
		current_id := page[i]
		for j := 0; j < i; j++ {
			prev_id := page[j]
			_, is_in := rules[order{before: current_id, after: prev_id}]
			if is_in {
				return false
			}
		}
	}
	return true
}

func getMiddleId(page []int) int {
	idx := len(page) / 2
	return page[idx]
}

func main() {
	input, _ := os.ReadFile("2024/day05_pt1/input.txt")
	input_parts := strings.Split(string(input), "\n\n")

	rules := parseOrder(input_parts[0])
	pages := parsePages(input_parts[1])
	middle_sum := 0

	for i, page := range pages {
		is_ok := checkPage(page, rules)
		if is_ok {
			middle_sum += getMiddleId(page)
		}
		fmt.Printf("Page %d: %v --> %t\n", i, page, is_ok)
	}

	fmt.Println(middle_sum)
}
