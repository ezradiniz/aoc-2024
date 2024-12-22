package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(name string) []string {
	bytes, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	content := strings.TrimSpace(string(bytes))
	if content == "" {
		return []string{}
	}
	return strings.Split(content, "\n")
}

func calcSecret(secret int) int {
	const mod int = 16777216
	secret = ((secret * 64) ^ secret) % mod
	secret = ((secret / 32) ^ secret) % mod
	secret = ((secret * 2048) ^ secret) % mod
	return secret
}

func generateSecret(secret int) int {
	for i := 0; i < 2000; i++ {
		secret = calcSecret(secret)
	}
	return secret
}

type Sequence struct {
	a, b, c, d int
}

func getBestPrice(sequences map[Sequence]int, secret int) int {
	changes := make([]int, 0)
	seen := make(map[Sequence]bool)
	best := 0
	for i := 0; i < 2000; i++ {
		pSecret := secret
		secret = calcSecret(secret)
		changes = append(changes, secret%10-pSecret%10)
		if len(changes) >= 4 {
			seq := Sequence{changes[i-3], changes[i-2], changes[i-1], changes[i]}
			if _, ok := seen[seq]; !ok {
				seen[seq] = true
				sequences[seq] += secret % 10
				best = max(best, sequences[seq])
			}
		}
	}
	return best
}

func main() {
	lines := readFile("./input.txt")

	part1 := 0
	part2 := 0
	sequences := make(map[Sequence]int)
	for _, line := range lines {
		initialSecret, _ := strconv.Atoi(line)
		part1 += generateSecret(initialSecret)
		part2 = max(part2, getBestPrice(sequences, initialSecret))
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
