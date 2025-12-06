package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type interval struct {
	start int
	end   int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	input_intervals := strings.Split(string(input), "\n\n")[0]
	input_products := strings.Split(string(input), "\n\n")[1]
	intervals := parseIntervals(input_intervals)
	intervals = mergeIntervals(intervals)
	products := parseProducts(input_products)
	fmt.Println(len(products), len(intervals))
	fmt.Println(countFreshProducts(intervals, products))
}

func parseIntervals(input_intervals string) []interval {
	intervals := make([]interval, 0, 0)
	for _, row := range strings.Split(input_intervals, "\n") {
		start_str := strings.Split(row, "-")[0]
		end_str := strings.Split(row, "-")[1]
		start, _ := strconv.Atoi(start_str)
		end, _ := strconv.Atoi(end_str)
		intervals = append(intervals, interval{start, end})
	}
	return intervals
}

func cmpIntervals(a, b interval) int {
	if a.start == b.start {
		return cmp.Compare(a.end, b.end)
	}
	return cmp.Compare(a.start, b.start)
}

func mergeIntervals(intervals []interval) []interval {
	merged_intervals := make([]interval, 0, 0)
	slices.SortFunc(intervals, cmpIntervals)

	merged_intervals = append(merged_intervals, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := len(merged_intervals) - 1

		if intervals[i].start > merged_intervals[last].end+1 {
			// [   ] (   )
			merged_intervals = append(merged_intervals, intervals[i])
		} else if intervals[i].end >= merged_intervals[last].end {
			// [   ( ]   )
			merged_intervals[last].end = intervals[i].end
		} else {
			// [   ( )   ]
			continue
		}
	}
	return merged_intervals
}

func parseProducts(input_products string) []int {
	products := make([]int, 0, 0)
	for _, product_str := range strings.Split(input_products, "\n") {
		product, _ := strconv.Atoi(product_str)
		products = append(products, product)
	}
	slices.Sort(products)
	return products
}

func countFreshProducts(intervals []interval, products []int) int {
	count_fresh := 0
	current_interval := 0
	current_product := 0

	for current_product < len(products) && current_interval < len(intervals) {
		product := products[current_product]
		if product < intervals[current_interval].start {
			current_product++
		} else if product <= intervals[current_interval].end {
			count_fresh += 1
			current_product++
		} else {
			current_interval++
		}
	}
	return count_fresh
}
