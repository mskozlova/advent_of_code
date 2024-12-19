package main

import (
	"fmt"
	"os"
	"strings"
)

type status struct {
	index  int
	prefix string
}

func main() {
	file_name := "2024/day19_pt2/input.txt"
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
		statuses := make(map[status]int)
		possible_options := countDesignOptions(design, towels, 0, &statuses)
		fmt.Printf("Design %d: %s -> %d\n", i, design, possible_options)
		n_possible += possible_options
	}
	fmt.Println(n_possible)
}

func parseTowels(input string) []string {
	return strings.Split(strings.Split(input, "\n\n")[0], ", ")
}

func parseDesigns(input string) []string {
	return strings.Split(strings.Split(input, "\n\n")[1], "\n")
}

func countDesignOptions(design string, towels []string, index int, statuses *map[status]int) int {
	if index == len(design) {
		return 1
	}

	design_options := 0
	for _, towel := range towels {
		if strings.HasPrefix(design[index:], towel) {
			n_options, has_value := (*statuses)[status{index: index, prefix: towel}]
			if !has_value {
				n_options = countDesignOptions(design, towels, index+len(towel), statuses)
				(*statuses)[status{index: index, prefix: towel}] = n_options
			}
			design_options += n_options
		}
	}
	return design_options
}
