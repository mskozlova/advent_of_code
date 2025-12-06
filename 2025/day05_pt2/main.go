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
	intervals := parseIntervals(input_intervals)
	intervals = mergeIntervals(intervals)
	fmt.Println(countFreshProducts(intervals))
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

func countFreshProducts(intervals []interval) int {
	count_fresh := 0
	for _, interval := range intervals {
		count_fresh += (interval.end - interval.start + 1)
	}
	return count_fresh
}
