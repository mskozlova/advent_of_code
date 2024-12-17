package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type status struct {
	A, B, C int
	i       int
	output  []int
}

func pow(n, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func main() {
	file_name := "2024/day17_pt1/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	program_raw, program := parseProgram(string(input))

	register := parseRegisters(string(input))
	output := register.runProgram(program)
	output_str := formatOutput(output)
	fmt.Println("output:   ", output_str)
	fmt.Println("should be:", program_raw)
}

func parseRegisters(input string) status {
	registers := strings.Split(input, "\n")[:3]
	A, _ := strconv.Atoi(registers[0][12:])
	B, _ := strconv.Atoi(registers[1][12:])
	C, _ := strconv.Atoi(registers[2][12:])
	output := make([]int, 0)
	return status{A, B, C, 0, output}
}

func parseProgram(input string) (string, []int) {
	rows := strings.Split(input, "\n")
	program_raw := rows[len(rows)-1][9:]
	command_raw := strings.Split(program_raw, ",")
	commands := make([]int, len(command_raw))

	for i, sym := range command_raw {
		num, _ := strconv.Atoi(sym)
		commands[i] = num
	}
	return program_raw, commands
}

func (s *status) runProgram(program []int) []int {
	for s.i < len(program) {
		s.applyCommand(program[s.i], program[s.i+1])
	}
	return s.output
}

func (s *status) inc() {
	s.i += 2
}

func (s *status) getCombo(operand int) int {
	res := -1
	if operand <= 3 {
		res = operand
	} else if operand == 4 {
		res = s.A
	} else if operand == 5 {
		res = s.B
	} else if operand == 6 {
		res = s.C
	} else {
		panic("Unexpected operand value")
	}
	return res
}

func (s *status) applyCommand(command int, operand int) {
	fmt.Printf("\nRunning command %d, operand %d:\tA %d, B %d, C %d\n", command, operand, s.A, s.B, s.C)
	if command == 0 {
		// division -> A
		num := s.A
		denom := pow(2, s.getCombo(operand))
		s.A = num / denom
		s.inc()
	} else if command == 1 {
		// bitwise XOR
		s.B = s.B ^ operand
		s.inc()
	} else if command == 2 {
		// modulo 8
		s.B = s.getCombo(operand) % 8
		s.inc()
	} else if command == 3 {
		// nothing or jump
		if s.A == 0 {
			s.inc()
			return
		} else {
			s.i = operand
		}
	} else if command == 4 {
		// bitwise XOR B ^ C
		s.B = s.B ^ s.C
		s.inc()
	} else if command == 5 {
		// modulo 8 + output
		s.output = append(s.output, s.getCombo(operand)%8)
		s.inc()
	} else if command == 6 {
		// division -> B
		num := s.A
		denom := pow(2, s.getCombo(operand))
		s.B = num / denom
		s.inc()
	} else if command == 7 {
		// division -> C
		num := s.A
		denom := pow(2, s.getCombo(operand))
		s.C = num / denom
		s.inc()
	} else {
		panic("Unknown instruction!")
	}

	fmt.Printf("...result A %d, B %d, C %d, output: %v, index: %d\n", s.A, s.B, s.C, s.output, s.i)
}

func formatOutput(output []int) string {
	output_str := make([]string, len(output))
	for i, num := range output {
		output_str[i] = strconv.Itoa(num)
	}
	return strings.Join(output_str, ",")
}
