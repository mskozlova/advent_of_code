package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	n_blinks int = 75
)

type status struct {
	stone       int
	in_n_blinks int
}

func main() {
	input, _ := os.ReadFile("2024/day11_pt1/input.txt")
	stones := getStones(string(input))
	precounts := make(map[status]int)
	fmt.Println(getStoneCountAfterNBlinks(stones, n_blinks, &precounts))
}

func getStones(input string) []int {
	stones_str := strings.Split(input, " ")
	stones := make([]int, len(stones_str))

	for i, stone_str := range stones_str {
		stone, _ := strconv.Atoi(stone_str)
		stones[i] = stone
	}
	return stones
}

func applyRules(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	stone_str := strconv.Itoa(stone)
	if len(stone_str)%2 == 0 {
		middle := len(stone_str) / 2
		stone1, _ := strconv.Atoi(stone_str[:middle])
		stone2, _ := strconv.Atoi(stone_str[middle:])
		return []int{stone1, stone2}
	}

	return []int{stone * 2024}
}

func getStoneCountAfterNBlinks(stones []int, n int, precounts *map[status]int) int {
	if n < 0 {
		return 1
	}

	total_stone_count := 0
	for _, stone := range stones {
		stone_count, is_saved := (*precounts)[status{stone: stone, in_n_blinks: n}]
		fmt.Printf("Searching stone %d in %d blinks - %t, %d\n", stone, n, is_saved, stone_count)
		if !is_saved {
			after_blink := applyRules(stone)
			stone_count = getStoneCountAfterNBlinks(after_blink, n-1, precounts)
			(*precounts)[status{stone: stone, in_n_blinks: n}] = stone_count
			fmt.Printf("Saving stone %d in %d blinks - %d\n", stone, n, stone_count)
		}
		total_stone_count += stone_count
	}
	return total_stone_count
}
