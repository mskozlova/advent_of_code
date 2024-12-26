package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	and int = iota
	or  int = iota
	xor int = iota
)

type node struct {
	id      string
	arg1    string
	arg2    string
	command int
}

type graph map[string]*node

func xor_bool(a, b bool) bool {
	if a == b {
		return false
	}
	return true
}

func main() {
	file_name := "2024/day24_pt1/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	input_parts := strings.Split(string(input), "\n\n")
	values := parseInput(strings.Split(input_parts[0], "\n"))
	g := parseGraph(strings.Split(input_parts[1], "\n"))
	start_ids := getStartIds(&g)

	fmt.Println(values)
	fmt.Println(g)
	// fmt.Println(start_ids)

	calc_order := getCalculationOrder(&g, start_ids)
	fmt.Println(calc_order)

	for i, id := range calc_order {
		fmt.Printf("%d Calculating node %s -> ", i+1, id)
		g[id].calculate(&values)
		fmt.Printf("%t\n", values[id])
	}

	answer_str := ""
	for _, id := range start_ids {
		if values[id] {
			answer_str = answer_str + "1"
		} else {
			answer_str = answer_str + "0"
		}
	}
	fmt.Println(answer_str)
	answer_num, err := strconv.ParseInt(answer_str, 2, 64)
	fmt.Println(answer_num)
}

func (n *node) calculate(values *map[string]bool) {
	var result bool
	arg1_val, arg1_exists := (*values)[n.arg1]
	arg2_val, arg2_exists := (*values)[n.arg2]

	if !arg1_exists || !arg2_exists {
		fmt.Printf("Node: %s: arg1: %s, arg2: %s, command: %d.\nValues: %v\n", n.id, n.arg1, n.arg2, n.command, *values)
		panic("Trying to calculate node with some of the args not ready")
	}

	switch op := n.command; op {
	case or:
		result = arg1_val || arg2_val
	case xor:
		result = xor_bool(arg1_val, arg2_val)
	case and:
		result = arg1_val && arg2_val
	default:
		panic("Undefined command")
	}
	(*values)[n.id] = result
}

func parseInput(input []string) map[string]bool {
	values := make(map[string]bool)
	for _, row := range input {
		id := strings.Split(row, ": ")[0]
		value_str := strings.Split(row, ": ")[1]
		value, _ := strconv.Atoi(value_str)
		values[id] = (value == 1)
	}
	return values
}

func parseGraphRow(row string) node {
	var n node
	n.id = strings.Split(row, " -> ")[1]
	formula := strings.Split(strings.Split(row, " -> ")[0], " ")
	n.arg1 = formula[0]
	n.arg2 = formula[2]

	switch command := formula[1]; command {
	case "OR":
		n.command = or
	case "XOR":
		n.command = xor
	case "AND":
		n.command = and
	default:
		n.command = -1
	}
	return n
}

func parseGraph(input []string) graph {
	g := make(graph)
	for _, row := range input {
		n := parseGraphRow(row)
		g[n.id] = &n
	}
	return g
}

func getStartIds(g *graph) []string {
	ids := make([]string, 0)
	for id := range *g {
		if strings.HasPrefix(id, "z") {
			ids = append(ids, id)
		}
	}
	slices.Sort(ids)
	slices.Reverse(ids)
	return ids
}

func dfs(n *node, g *graph, visited *map[string]bool, order *[]string) {
	fmt.Println(n.id)
	(*visited)[n.id] = true
	next_nodes := []string{n.arg1, n.arg2}
	fmt.Println("\tnext:", next_nodes)
	for _, nid := range next_nodes {
		_, is_visited := (*visited)[nid]
		_, has_node := (*g)[nid]
		if !is_visited && has_node {
			dfs((*g)[nid], g, visited, order)
		}
	}
	*order = append(*order, n.id)
}

func getCalculationOrder(g *graph, start_ids []string) []string {
	order := make([]string, 0)
	visited := make(map[string]bool)

	for _, id := range start_ids {
		_, is_visited := visited[id]
		if !is_visited {
			// order = append(order, id)
			dfs((*g)[id], g, &visited, &order)
		}
	}
	return order
}
