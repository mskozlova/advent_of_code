package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type problem struct {
	op      string
	width   int
	numbers []int64
}

func main() {
	input, _ := os.ReadFile("input.txt")
	problems := parseProblems(string(input))
	parseNumbers(string(input), problems)
	fmt.Println(problems)

	var answer int64 = 0
	for _, p := range problems {
		answer += p.Solve()
	}
	fmt.Println(answer)
}

func parseNumbers(input string, problems []problem) {
	rows := strings.Split(input, "\n")
	for i := range len(rows) - 1 {
		current_sym := 0
		problem_id := 0
		digit_id := 0

		for current_sym < len(rows[i]) {
			if digit_id == problems[problem_id].width {
				problem_id += 1
				current_sym += 1
				digit_id = 0
			} else {
				sym := string(rows[i][current_sym])
				if sym != " " {
					num, _ := strconv.Atoi(sym)
					problems[problem_id].numbers[digit_id] = problems[problem_id].numbers[digit_id]*10 + int64(num)
				}
				digit_id += 1
				current_sym += 1
			}
		}
	}
}

func parseProblems(input string) []problem {
	rows := strings.Split(input, "\n")
	operations := make([]problem, 0, 0)

	for _, sym := range rows[len(rows)-1] {
		if !unicode.IsSpace(sym) {
			operations = append(operations, problem{string(sym), 0, make([]int64, 0, 0)})
		} else {
			operations[len(operations)-1].width += 1
			operations[len(operations)-1].numbers = append(operations[len(operations)-1].numbers, 0)
		}
	}
	operations[len(operations)-1].width += 1
	operations[len(operations)-1].numbers = append(operations[len(operations)-1].numbers, 0)

	return operations
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
