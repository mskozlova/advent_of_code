package main

import (
	"fmt"
	"os"
	"strings"
)

type key []int
type lock []int

func main() {
	file_name := "2024/day25_pt1/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	elems := strings.Split(string(input), "\n\n")
	keys, locks := parseInput(elems)
	max_height := len(strings.Split(elems[0], "\n")) - 2
	fmt.Println("Keys:", len(keys))
	fmt.Println("Locks:", len(locks))

	ok := 0
	for ik, k := range keys {
		for il, l := range locks {
			if !k.overlaps(l, max_height) {
				ok += 1
			} else if (ik < 10) && (il < 10) {
				fmt.Println("key", ik, "-", k, "lock", il, "-", l)
			}
		}
	}
	fmt.Println(ok)
}

func parseInput(input []string) ([]key, []lock) {
	keys := make([]key, 0)
	locks := make([]lock, 0)

	for _, elem := range input {
		rows := strings.Split(elem, "\n")
		is_key := false
		k := make(key, len(rows[0]))
		l := make(lock, len(rows[0]))
		for row_num, row := range rows {
			if row_num == 0 {
				if string(row[0]) == "." {
					is_key = true
				}
			}

			for i, sym := range row {
				if !is_key {
					if string(sym) == "#" {
						l[i] = row_num
					}
				} else {
					if string(sym) == "." {
						k[i] = len(rows) - row_num - 2
					}
				}
			}
		}

		if is_key {
			keys = append(keys, k)
		} else {
			locks = append(locks, l)
		}
	}
	return keys, locks
}

func (k *key) overlaps(l lock, max_height int) bool {
	for i := range len(l) {
		if (*k)[i]+l[i] > max_height {
			return true
		}
	}
	return false
}
