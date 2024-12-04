package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type event struct {
	start_idx int
	end_idx   int
	t         string
}

func performOperation(e event, input []byte) int {
	op := input[e.start_idx:e.end_idx]
	re_num := regexp.MustCompile(`\d{1,3}`)
	result := 1
	for _, num_b := range re_num.FindAll(op, -1) {
		num, _ := strconv.Atoi(string(num_b))
		result *= num
	}
	return result
}

func createGrid(op_idxs [][]int, do_idxs [][]int, dont_idxs [][]int) []event {
	grid := make([]event, 0)
	for _, coords := range op_idxs {
		grid = append(
			grid,
			event{start_idx: coords[0], end_idx: coords[1], t: "op"},
		)
	}

	for _, coords := range do_idxs {
		grid = append(
			grid,
			event{start_idx: coords[0], end_idx: coords[1], t: "do"},
		)
	}

	for _, coords := range dont_idxs {
		grid = append(
			grid,
			event{start_idx: coords[0], end_idx: coords[1], t: "dont"},
		)
	}

	slices.SortFunc(grid, func(a, b event) int {
		return a.start_idx - b.start_idx
	})
	return grid
}

func filterOperations(op_idxs [][]int, do_idxs [][]int, dont_idxs [][]int) []event {
	do := true
	filtered := make([]event, 0)

	for _, event := range createGrid(op_idxs, do_idxs, dont_idxs) {
		if do && (event.t == "op") {
			filtered = append(filtered, event)
		} else if event.t == "do" {
			do = true
		} else if event.t == "dont" {
			do = false
		}
	}
	return filtered
}

func main() {
	input, _ := os.ReadFile("2024/day03_pt2/input.txt")

	re_ops := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	re_dos := regexp.MustCompile(`do\(\)`)
	re_donts := regexp.MustCompile(`don't\(\)`)

	ops := re_ops.FindAllIndex(input, -1)
	dos := re_dos.FindAllIndex(input, -1)
	donts := re_donts.FindAllIndex(input, -1)

	filtered_ops := filterOperations(ops, dos, donts)

	total_sum := 0
	for _, op := range filtered_ops {
		total_sum += performOperation(op, input)
	}

	fmt.Printf("%d\n", total_sum)
}
