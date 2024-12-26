package main

import (
	"fmt"
	"maps"
	"math/rand"
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
type swap map[string]string

func xor_bool(a, b bool) bool {
	if a == b {
		return false
	}
	return true
}

func pow_int(a, b int) int {
	if a == 1 {
		return 1
	}
	if b == 0 {
		return 1
	}
	res := 1
	for range b {
		res *= a
	}
	return res
}

func bin(a int64) []bool {
	b := make([]bool, 0)
	for a > 0 {
		b = append(b, a%2 == 1)
		a /= 2
	}
	return b
}

func main() {
	file_name := "2024/day24_pt2/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	input_parts := strings.Split(string(input), "\n\n")
	_, x_ids, y_ids := parseInput(strings.Split(input_parts[0], "\n"))
	r := rand.New(rand.NewSource(42))

	swappable_ids := swappableIds(strings.Split(input_parts[1], "\n"))
	fmt.Println("Swappable:", swappable_ids)
	ss := generateSwaps(swappable_ids, 1)
	fmt.Println("Generated swaps!", len(ss))

	s := make(swap)

	errors := 46
	examples := 1000
	if eval := true; eval {
		for run_id := range 4 {
			best_swap, min_errors := runSwaps(strings.Split(input_parts[1], "\n"), s, ss, x_ids, y_ids, r, examples, errors)
			errors = min_errors
			maps.Copy(s, best_swap)
			fmt.Printf(">>>>>> Run #%d finished - Errors %d, best swap %v <<<<<<\n", run_id+1, errors, s)
		}
	}
}

func (n *node) get_args(s *swap) (string, string) {
	arg1 := n.arg1
	arg2 := n.arg2

	new_arg1, has_swap1 := (*s)[n.arg1]
	new_arg2, has_swap2 := (*s)[n.arg2]

	if has_swap1 {
		arg1 = new_arg1
	}
	if has_swap2 {
		arg2 = new_arg2
	}
	return arg1, arg2
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

func parseInput(input []string) (map[string]bool, []string, []string) {
	values := make(map[string]bool)
	x_ids := make([]string, 0)
	y_ids := make([]string, 0)
	for _, row := range input {
		id := strings.Split(row, ": ")[0]
		value_str := strings.Split(row, ": ")[1]
		value, _ := strconv.Atoi(value_str)
		values[id] = (value == 1)

		if strings.HasPrefix(id, "x") {
			x_ids = append(x_ids, id)
		} else {
			y_ids = append(y_ids, id)
		}
	}

	slices.Sort(x_ids)
	slices.Sort(y_ids)
	slices.Reverse(x_ids)
	slices.Reverse(y_ids)
	return values, x_ids, y_ids
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

func getIds(g *graph, prefix string) []string {
	ids := make([]string, 0)
	for id := range *g {
		if strings.HasPrefix(id, prefix) {
			ids = append(ids, id)
		}
	}
	slices.Sort(ids)
	slices.Reverse(ids)
	return ids
}

func dfs(n *node, g *graph, visited *map[string]bool, order *[]string) {
	(*visited)[n.id] = true
	next_nodes := []string{n.arg1, n.arg2}
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
			dfs((*g)[id], g, &visited, &order)
		}
	}
	return order
}

func getNumber(ids []string, values *map[string]bool) (string, int64) {
	answer_str := ""
	for _, id := range ids {
		if (*values)[id] {
			answer_str = answer_str + "1"
		} else {
			answer_str = answer_str + "0"
		}
	}
	answer_num, _ := strconv.ParseInt(answer_str, 2, 64)
	return answer_str, answer_num
}

func fillValues(ids []string, values *map[string]bool, number int) {
	ids_sorted := make([]string, len(ids))
	copy(ids_sorted, ids)
	slices.Sort(ids_sorted)

	for _, id := range ids_sorted {
		(*values)[id] = (number%2 == 1)
		number /= 2
	}
}

func makeGraph(input []string, s swap) graph {
	g := parseGraph(input)
	g.applySwap(s)
	return g
}

func checkAnswer(x_ids, y_ids []string, x, y int, g *graph) map[int]bool {
	output_ids := getIds(g, "z")
	values := make(map[string]bool)

	fillValues(x_ids, &values, x)
	fillValues(y_ids, &values, y)

	calc_order := getCalculationOrder(g, output_ids)

	for _, id := range calc_order {
		(*g)[id].calculate(&values)
	}

	x_bin, _ := getNumber(x_ids, &values)
	y_bin, _ := getNumber(y_ids, &values)
	z_bin, z := getNumber(output_ids, &values)

	if do_print := false; z != int64(x+y) && do_print {
		fmt.Println("x:", x_bin, x)
		fmt.Println("y:", y_bin, y)
		fmt.Println("z:", z_bin, z)
		fmt.Println("should be", x+y)
	}

	return get_1s(bin(z ^ int64(x+y)))
}

func get_1s(b []bool) map[int]bool {
	m := make(map[int]bool)
	for i, f := range b {
		if f {
			m[i] = true
		}
	}
	return m
}

func selectIds(ids []string, start_idx int, target int, current []string, all *[][]string) {
	if len(current) == target {
		add := make([]string, len(current))
		copy(add, current)
		*all = append(*all, add)
		return
	}

	for i := start_idx; i < len(ids); i++ {
		current = append(current, ids[i])
		selectIds(ids, i+1, target, current, all)
		current = current[:len(current)-1]
	}
}

func generatePairs(ids []string, pair_start int, current []int, ss *[]swap) {
	if len(current) == len(ids) {
		s := make(swap)
		for i := 0; i < len(current); i += 2 {
			s[ids[current[i]]] = ids[current[i+1]]
			s[ids[current[i+1]]] = ids[current[i]]
		}
		*ss = append(*ss, s)
		return
	}

	start := 0
	if len(current)%2 == 1 {
		start = pair_start
	}

	for i := start; i < len(ids); i++ {
		if slices.ContainsFunc(current, func(n int) bool {
			return n == i
		}) {
			continue
		}

		if len(current)%2 == 1 {
			current = append(current, i)
			generatePairs(ids, 0, current, ss)
			current = current[:len(current)-1]
		} else {
			current = append(current, i)
			generatePairs(ids, i, current, ss)
			break
		}
	}
}

func generateSwaps(ids []string, n int) []swap {
	swaps := make([][]swap, 0)
	comb_ids := make([][]string, 0)
	current := make([]string, 0)

	selectIds(ids, 0, n*2, current, &comb_ids)
	for _, comb := range comb_ids {
		ss := make([]swap, 0)
		current := make([]int, 0)
		generatePairs(comb, 0, current, &ss)
		swaps = append(swaps, ss)
	}

	return slices.Concat(swaps...)
}

func (g *graph) applySwap(s swap) {
	for _, n := range *g {
		new_arg1, has_swap1 := s[n.arg1]
		new_arg2, has_swap2 := s[n.arg2]

		if has_swap1 {
			n.arg1 = new_arg1
		}
		if has_swap2 {
			n.arg2 = new_arg2
		}
	}
}

func cycleUtil(id string, g *graph, visited *map[string]bool, rec_stack *map[string]bool) bool {
	_, is_visited := (*visited)[id]
	_, is_node := (*g)[id]

	if !is_node {
		return false
	}

	if !is_visited {
		(*visited)[id] = true
		(*rec_stack)[id] = true

		for _, next_id := range []string{(*g)[id].arg1, (*g)[id].arg2} {
			_, is_next_visited := (*visited)[next_id]
			_, is_next_rec_stack := (*rec_stack)[next_id]
			if !is_next_visited && cycleUtil(next_id, g, visited, rec_stack) {
				return true
			} else if is_next_rec_stack {
				return true
			}
		}
	}
	delete((*rec_stack), id)
	return false
}

func hasCycles(g *graph) bool {
	visited := make(map[string]bool)
	rec_stack := make(map[string]bool)
	for id := range *g {
		_, is_visited := visited[id]
		if !is_visited && cycleUtil(id, g, &visited, &rec_stack) {
			return true
		}
	}
	return false
}

type example struct {
	x, y int
}

func generateExamples(min, max, n int, r *rand.Rand) []example {
	examples := make([]example, n)
	for i := range n {
		x := min + r.Intn(max-min)
		y := min + r.Intn(max-min)
		examples[i] = example{x, y}
	}
	return examples
}

func swappableIds(input []string) []string {
	g := parseGraph(input)
	ids := make([]string, len(g))
	i := 0
	for id := range g {
		ids[i] = id
		i += 1
	}
	slices.Sort(ids)
	return ids
}

func runSwaps(input []string, current_swap swap, ss []swap, x_ids, y_ids []string, r *rand.Rand, n_examples int, min_errors int) (swap, int) {
	var best_swap swap
	examples := generateExamples(0, pow_int(2, 45), n_examples, r)

	for sid, s := range ss {
		if sid%100 == 0 {
			fmt.Println("Iterations done:", sid, "Min errors:", min_errors, "Best swap:", best_swap)
		}
		maps.Copy(s, current_swap)

		g := makeGraph(input, s)
		if hasCycles(&g) {
			continue
		}

		wrong_bits := make(map[int]bool)
		for _, e := range examples {
			wb := checkAnswer(x_ids, y_ids, e.x, e.y, &g)
			maps.Copy(wrong_bits, wb)
			if len(wrong_bits) >= min_errors {
				break
			}
		}
		if len(wrong_bits) < min_errors {
			best_swap = s
			min_errors = len(wrong_bits)
		}
	}
	return best_swap, min_errors
}
