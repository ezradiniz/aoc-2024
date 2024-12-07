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

func parseEquation(line string) (int, []int) {
	split := strings.Split(line, ": ")
	result, _ := strconv.Atoi(split[0])
	values := make([]int, 0)
	for _, num := range strings.Split(split[1], " ") {
		val, _ := strconv.Atoi(num)
		values = append(values, val)
	}
	return result, values
}

func concat(a, b int) int {
	val, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return val
}

func canCalibrate(result int, values []int, allowConcat bool) bool {
	var rec func(i int, res int) bool
	rec = func(i, res int) bool {
		if i >= len(values) {
			return res == result
		}
		return rec(i+1, res+values[i]) || (i > 0 && rec(i+1, res*values[i])) || (allowConcat && rec(i+1, concat(res, values[i])))
	}
	return rec(0, 0)
}

func main() {
	lines := readFile("./input.txt")

	var part1 int
	var part2 int
	for _, line := range lines {
		result, values := parseEquation(line)
		if canCalibrate(result, values, false) {
			part1 += result
		}
		if canCalibrate(result, values, true) {
			part2 += result
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
