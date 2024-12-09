package main

import (
	"fmt"
	"os"
	"strconv"
)

type block struct {
	size      int
	id        int
	start_idx int
	is_moved  bool
}

func main() {
	input, _ := os.ReadFile("2024/day09_pt2/input.txt")

	// checksum := 0
	files, free := parseFileBlocks(string(input))

	fmt.Println("files prefix:", files[:10], "\nfree prefix:", free[:10], "")
	// fmt.Println("files suffix:", files[len(files)-10:], "\nfree suffix:", free[len(free)-10:], "")
	fmt.Println("files len:", len(files), "\nfree len:", len(free), "")

	fmt.Println(moveFileBlocks(files, free))
}

func parseFileBlocks(input string) ([]block, []block) {
	files := make([]block, 0)
	free := make([]block, 0)
	idx := 0

	for i := range len(input) {
		b_str := string(input[i])
		b, _ := strconv.Atoi(b_str)

		if i%2 == 0 {
			// block is file
			files = append(files, block{size: b, id: i / 2, start_idx: idx, is_moved: false})
		} else {
			// block is free
			free = append(free, block{size: b, id: i / 2, start_idx: idx, is_moved: false})
		}
		idx += b
	}

	return files, free
}

func moveFileBlocks(files []block, free []block) int {
	checksum := 0
	free_idx := 0

	for free_idx < len(free) {
		fmt.Printf("Free #%d\tChecksum - %d\n", free_idx, checksum)
		// adding previous file info to checksum
		if !files[free_idx].is_moved {
			for i := range files[free_idx].size {
				current_pos := files[free_idx].start_idx + i
				fmt.Printf("\tadding file space: pos #%d file id %d, total %d\n", current_pos, free_idx, free_idx*current_pos)
				checksum += free_idx * current_pos
			}
		}

		// moving last file blocks into current free block
		file_idx := len(files) - 1
		moved := 0
		for (moved != free[free_idx].size) && (file_idx > free_idx) {
			if (files[file_idx].size <= free[free_idx].size-moved) && !files[file_idx].is_moved {
				fmt.Printf("\t\tspace left to move %d\n", free[free_idx].size-moved)
				fmt.Printf("\t\tmoving file #%d -> free #%d\n", file_idx, free_idx)
				// moving
				for i := range files[file_idx].size {
					current_pos := free[free_idx].start_idx + moved + i
					fmt.Printf("\tadding free space: pos #%d file id %d, total %d\n", current_pos, file_idx, file_idx*current_pos)
					checksum += file_idx * current_pos
				}
				files[file_idx].is_moved = true
				moved += files[file_idx].size
			}
			file_idx--
			if file_idx < 0 {
				break
			}
		}
		free_idx += 1
	}

	return checksum
}
