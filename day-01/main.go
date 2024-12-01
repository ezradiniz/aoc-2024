package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(name string) []string {
	bytes, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")
	return lines[:len(lines)-1]
}

func main() {
	lines := readFile("./input.txt")

	list1 := make([]int, 0)
	list2 := make([]int, 0)
	freq2 := make(map[int]int)

	for _, line := range lines {
		pairs := strings.Split(line, "   ")

		if n, err := strconv.Atoi(pairs[0]); err == nil {
			list1 = append(list1, n)
		}
		if n, err := strconv.Atoi(pairs[1]); err == nil {
			list2 = append(list2, n)
			freq2[n] += 1
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)

	dist := 0
	score := 0
	for i := range list1 {
		dist += abs(list2[i] - list1[i])
		score += list1[i] * freq2[list1[i]]
	}

	fmt.Println("Part 1: ", dist)
	fmt.Println("Part 2: ", score)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
