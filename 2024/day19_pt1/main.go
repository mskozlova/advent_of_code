package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file_name := "2024/day19_pt1/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	designs := parseDesigns(string(input))
	towels := parseTowels(string(input))
	n_possible := 0
	fmt.Println("Towels:", towels)

	for i, design := range designs {
		is_possible := checkDesign(design, towels, 0)
		fmt.Printf("Design %d: %s -> %t\n", i, design, is_possible)
		if is_possible {
			n_possible += 1
		}
	}
	fmt.Println(n_possible)
}

func parseTowels(input string) []string {
	return strings.Split(strings.Split(input, "\n\n")[0], ", ")
}

func parseDesigns(input string) []string {
	return strings.Split(strings.Split(input, "\n\n")[1], "\n")
}

func checkDesign(design string, towels []string, index int) bool {
	if index == len(design) {
		return true
	}

	can_continue := false
	for _, towel := range towels {
		if strings.HasPrefix(design[index:], towel) {
			can_continue = can_continue || checkDesign(design, towels, index+len(towel))
		}
		if can_continue {
			break
		}
	}
	return can_continue
}
