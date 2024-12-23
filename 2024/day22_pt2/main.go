package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	n_steps     int = 2000
	prune_param int = 16777216
)

type pattern struct {
	n0, n1, n2, n3 int
}

func main() {
	file_name := "2024/day22_pt2/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	secrets := parseInitSecrets(strings.Split(string(input), "\n"))
	wins := make(map[pattern]int)

	for _, secret := range secrets {
		runSecretEvolution(secret, &wins)
	}

	max_win := -1
	var best_pattern pattern
	for p, value := range wins {
		if value > max_win {
			max_win = value
			best_pattern = p
		}
	}
	fmt.Println(max_win, best_pattern)
}

func (p *pattern) fill(observed []int) {
	p.n0 = observed[1] - observed[0]
	p.n1 = observed[2] - observed[1]
	p.n2 = observed[3] - observed[2]
	p.n3 = observed[4] - observed[3]
}

func update(observed []int, new_value int) []int {
	observed = append(observed, new_value)
	for len(observed) > 5 {
		observed = observed[1:]
	}
	return observed
}

func parseInitSecrets(input []string) []int {
	secrets := make([]int, len(input))
	for i, str := range input {
		secret, _ := strconv.Atoi(str)
		secrets[i] = secret
	}
	return secrets
}

func runSecretEvolution(secret int, wins *map[pattern]int) {
	occurences := make(map[pattern]bool)
	var p pattern
	prices := make([]int, 0)
	for range n_steps {
		secret = mix_and_prune(secret, secret*64)
		secret = mix_and_prune(secret, secret/32)
		secret = mix_and_prune(secret, secret*2048)

		price := getLastDigit(secret)
		prices = update(prices, price)

		if len(prices) < 5 {
			continue
		}
		p.fill(prices)
		_, seen := occurences[p]

		if !seen {
			occurences[p] = true
			value, has_value := (*wins)[p]
			if !has_value {
				value = 0
			}
			(*wins)[p] = value + price
		}
	}
}

func mix_and_prune(result, secret int) int {
	return (result ^ secret) % prune_param
}

func getLastDigit(num int) int {
	return num % 10
}
