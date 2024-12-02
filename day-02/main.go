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
		fmt.Println(err)
		os.Exit(1)
	}
	content := string(bytes)
	lines := strings.Split(content, "\n")
	return lines[:len(lines)-1]
}

func isSafe(line []string) bool {
	prevDir := 0
	prev, _ := strconv.Atoi(line[0])
	for i := 1; i < len(line); i++ {
		cur, _ := strconv.Atoi(line[i])
		delta := cur - prev
		if delta == 0 || abs(delta) > 3 {
			return false
		}
		curDir := max(-1, min(1, delta))
		if prevDir != 0 && curDir != prevDir {
			return false
		}
		prevDir = curDir
		prev = cur
	}
	return true
}

func isAlmostSafe(line []string) bool {
	for i := 0; i < len(line); i++ {
		if isSafe(remove(line, i)) {
			return true
		}
	}
	return false
}

func main() {
	lines := readFile("./input.txt")

	var safe int
	var almostSafe int

	for _, line := range lines {
		line := strings.Split(line, " ")
		if isSafe(line) {
			safe += 1
		}
		if isAlmostSafe(line) {
			almostSafe += 1
		}
	}

	fmt.Println("Part 1: ", safe)
	fmt.Println("Part 2: ", almostSafe)
}

func remove(arr []string, i int) []string {
	tmp := make([]string, len(arr))
	copy(tmp, arr)
	return append(tmp[:i], tmp[i+1:]...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
