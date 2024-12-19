package main

import (
	"fmt"
	"os"
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

func countWays(design string, patterns map[string]bool) int {
	n := len(design)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			sub := design[j-1 : i]
			if patterns[sub] {
				dp[i] += dp[j-1]
			}
		}
	}
	return dp[n]
}

func main() {
	lines := readFile("./input.txt")
	patterns := make(map[string]bool, 0)
	for _, pattern := range strings.Split(lines[0], ", ") {
		patterns[pattern] = true
	}

	part1 := 0
	part2 := 0
	for _, design := range lines[2:] {
		count := countWays(design, patterns)
		if count > 0 {
			part1++
		}
		part2 += count
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
