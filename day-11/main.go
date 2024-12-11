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

func isEven(num int) bool {
	return len(strconv.Itoa(num))%2 == 0
}

func countStones(line string, blink int) int64 {
	dp := make(map[int]int64, 0)
	for _, num := range strings.Split(line, " ") {
		val, _ := strconv.Atoi(num)
		dp[val] += 1
	}
	for b := 0; b < blink; b++ {
		tmp := make(map[int]int64, 0)
		for stone := range dp {
			if stone == 0 {
				tmp[1] += dp[stone]
			} else if isEven(stone) {
				s := strconv.Itoa(stone)
				s1, _ := strconv.Atoi(s[:len(s)/2])
				s2, _ := strconv.Atoi(s[len(s)/2:])
				tmp[s1] += dp[stone]
				tmp[s2] += dp[stone]
			} else {
				tmp[stone*2024] += dp[stone]
			}
		}
		dp = tmp
	}
	var ans int64
	for _, count := range dp {
		ans += count
	}
	return ans
}

func main() {
	line := readFile("./input.txt")[0]

	part1 := countStones(line, 25)
	part2 := countStones(line, 75)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
