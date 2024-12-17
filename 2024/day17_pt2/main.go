package main

import (
	"fmt"
	"slices"
)

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

type status struct {
	A int
	i int
}

func main() {
	program := []int{2, 4, 1, 5, 7, 5, 1, 6, 0, 3, 4, 6, 5, 5, 3, 0}
	slices.Reverse(program)

	queue := []status{{A: 0, i: 0}}

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		for A_modulo_8 := range 8 {
			A := 8*s.A + A_modulo_8
			B_cand := A_modulo_8 ^ 5 ^ 6 ^ (A / pow(2, A_modulo_8^5))
			if B_cand%8 == program[s.i] {
				if s.i+1 == len(program) {
					fmt.Printf("A: %d\tB: %d\tOUT: %d\n", A, B_cand, B_cand%8)
				} else {
					queue = append(queue, status{A: A, i: s.i + 1})
				}
			}
		}
	}
}
