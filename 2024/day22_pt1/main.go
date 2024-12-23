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

func main() {
	file_name := "2024/day22_pt1/input.txt"
	input, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	secrets := parseInitSecrets(strings.Split(string(input), "\n"))
	total_sum := 0

	for i, secret := range secrets {
		new_secret := runSecretEvolution(secret)
		total_sum += new_secret
		fmt.Printf("Secret #%d: %d -> %d\n", i, secret, new_secret)
	}
	fmt.Println(total_sum)
}

func parseInitSecrets(input []string) []int {
	secrets := make([]int, len(input))
	for i, str := range input {
		secret, _ := strconv.Atoi(str)
		secrets[i] = secret
	}
	return secrets
}

func runSecretEvolution(secret int) int {
	for range n_steps {
		secret = mix_and_prune(secret, secret*64)
		secret = mix_and_prune(secret, secret/32)
		secret = mix_and_prune(secret, secret*2048)
	}
	return secret
}

func mix_and_prune(result, secret int) int {
	return (result ^ secret) % prune_param
}
