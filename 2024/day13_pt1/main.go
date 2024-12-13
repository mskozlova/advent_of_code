package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	a_cost int = 3
	b_cost int = 1
)

type button struct {
	X, Y int
}
type prize struct {
	X, Y int
}
type machine struct {
	A, B button
	p    prize
}

func main() {
	input, _ := os.ReadFile("2024/day13_pt1/input.txt")
	total_price := 0

	for i, m := range parseMachines(string(input)) {
		m_str, _ := json.Marshal(m)
		price := solveMachine(m)
		fmt.Printf("Machine #%d: %s, price: %d\n", i, string(m_str), price)
		if price >= 0 {
			total_price += price
		}
	}

	fmt.Println(total_price)
}

func solveMachine(m machine) int {
	nb_denom := (m.B.X*m.A.Y - m.B.Y*m.A.X)
	nb_num := (m.p.X*m.A.Y - m.p.Y*m.A.X)

	if nb_num%nb_denom != 0 {
		return -1
	}

	nb := nb_num / nb_denom
	na_denom := m.A.X
	na_num := m.p.X - nb*m.B.X

	if na_num%na_denom != 0 {
		return -1
	}
	na := na_num / na_denom
	return na*a_cost + nb*b_cost
}

func parseMachines(input string) []machine {
	machines_raw := strings.Split(input, "\n\n")
	machines := make([]machine, len(machines_raw))

	for i, machine_raw := range machines_raw {
		rows := strings.Split(machine_raw, "\n")
		m := machine{}
		for j, row := range rows {
			if j == 0 {
				m.A = parseButton(row)
			} else if j == 1 {
				m.B = parseButton(row)
			} else {
				m.p = parsePrize(row)
			}
		}
		machines[i] = m
	}
	return machines
}

func parseButton(row string) button {
	elements := strings.Split(row[10:], ", ")
	X, _ := strconv.Atoi(elements[0][2:])
	Y, _ := strconv.Atoi(elements[1][2:])
	return button{X, Y}
}
func parsePrize(row string) prize {
	elements := strings.Split(row[7:], ", ")
	X, _ := strconv.Atoi(elements[0][2:])
	Y, _ := strconv.Atoi(elements[1][2:])
	return prize{X, Y}
}
