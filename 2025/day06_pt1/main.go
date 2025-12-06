package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	numbers []int64
	op      string
}

func main() {
	input, _ := os.ReadFile("input.txt")
	numbers := parseNumbers(string(input))
	operations := parseOperations(string(input))
	problems := createProblems(numbers, operations)
	fmt.Println(problems)

	var answer int64 = 0
	for _, p := range problems {
		answer += p.Solve()
	}
	fmt.Println(answer)
}

func parseNumbers(input string) [][]int64 {
	rows := strings.Split(input, "\n")
	numbers := make([][]int64, 0, 0)
	for i := range len(rows) - 1 {
		number_set := make([]int64, 0, 0)
		for _, num_str := range strings.Fields(rows[i]) {
			num, _ := strconv.Atoi(num_str)
			number_set = append(number_set, int64(num))
		}
		numbers = append(numbers, number_set)
	}
	return numbers
}

func parseOperations(input string) []string {
	rows := strings.Split(input, "\n")
	operations := make([]string, 0, 0)

	for _, sym := range strings.Fields(rows[len(rows)-1]) {
		operations = append(operations, sym)
	}
	return operations
}

func createProblems(numbers [][]int64, operations []string) []problem {
	problems := make([]problem, 0, 0)
	for i := range len(operations) {
		problems = append(problems, problem{make([]int64, 0, 0), operations[i]})
	}

	for _, num_set := range numbers {
		for i, num := range num_set {
			problems[i].numbers = append(problems[i].numbers, num)
		}
	}
	return problems
}

func (p *problem) Solve() int64 {
	var answer int64 = 0
	if p.op == "*" {
		answer = 1
	}
	for _, num := range p.numbers {
		if p.op == "*" {
			answer *= num
		} else {
			answer += num
		}
	}
	return answer
}
