package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("2024/day09_pt1/input.txt")

	// checksum := 0
	files, free := parseFileBlocks(string(input))

	fmt.Println("files:", files, "free:", free, "")

	fmt.Println(moveFileBlocks(files, free))
}

func parseFileBlocks(input string) ([]int, []int) {
	files := make([]int, 0)
	free := make([]int, 0)

	for i := range len(input) {
		block_str := string(input[i])
		block, _ := strconv.Atoi(block_str)

		if i%2 == 0 {
			// block is file
			files = append(files, block)
		} else {
			// block is free
			free = append(free, block)
		}
	}

	return files, free
}

func moveFileBlocks(files []int, free []int) int {
	current_pos := 0
	checksum := 0
	free_idx := 0
	file_idx := len(files)
	in_file_idx := 0

	for free_idx < file_idx {
		fmt.Printf("Free #%d\tFile #%d\tChecksum - %d\n", free_idx, file_idx, checksum)
		// adding previous file info to checksum
		for range files[free_idx] {
			fmt.Printf("\tadding file space: pos #%d file id %d, total %d\n", current_pos, free_idx, free_idx*current_pos)
			checksum += free_idx * current_pos
			current_pos += 1
		}

		// moving last file blocks into current free block
		moved := 0
		for moved != free[free_idx] {
			if in_file_idx == 0 {
				file_idx -= 1
				in_file_idx = files[file_idx]
			}

			moving := min(in_file_idx, free[free_idx]-moved)

			for range moving {
				fmt.Printf("\tadding free space: pos #%d file id %d, total %d\n", current_pos, file_idx, file_idx*current_pos)
				checksum += file_idx * current_pos
				current_pos += 1
			}

			moved += moving
			in_file_idx -= moving
		}
		free_idx += 1
	}

	for range in_file_idx {
		fmt.Printf("\tadding left over files: pos #%d file id %d, total %d\n", current_pos, file_idx, file_idx*current_pos)
		checksum += file_idx * current_pos
		current_pos += 1
	}

	return checksum
}
