package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	content := strings.TrimSpace(string(data))
	if content == "" {
		return []string{}
	}
	return strings.Split(content, "\n")
}

func parseRules(lines []string) (map[string]map[string]bool, int) {
	var lineIdx int
	rules := make(map[string]map[string]bool)
	for idx, line := range lines {
		if len(line) == 0 {
			lineIdx = idx + 1
			break
		}
		parts := strings.Split(line, "|")
		x, y := parts[0], parts[1]
		if _, ok := rules[y]; !ok {
			rules[y] = make(map[string]bool)
		}
		rules[y][x] = true
	}
	return rules, lineIdx
}

func topoSort(seq []string, rules map[string]map[string]bool) []string {
	graph := make(map[string][]string)
	deg := make(map[string]int)
	for _, u := range seq {
		for _, v := range seq {
			if u == v || rules[u][v] {
				continue
			}
			deg[v]++
			graph[u] = append(graph[u], v)
		}
	}
	queue := make([]string, 0)
	for node := range graph {
		if deg[node] == 0 {
			queue = append(queue, node)
		}
	}
	ans := make([]string, 0)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		ans = append(ans, cur)
		for _, nxt := range graph[cur] {
			deg[nxt]--
			if deg[nxt] == 0 {
				queue = append(queue, nxt)
			}
		}
	}
	return ans
}

func main() {
	lines := readFile("./input.txt")
	rules, startIdx := parseRules(lines)

	var part1 int
	var part2 int

	for _, line := range lines[startIdx:] {
		seq1 := strings.Split(line, ",")
		seq2 := topoSort(seq1, rules)
		if slices.Compare(seq1, seq2) == 0 {
			mid, _ := strconv.Atoi(seq1[len(seq1)/2])
			part1 += mid
		} else {
			mid, _ := strconv.Atoi(seq2[len(seq2)/2])
			part2 += mid
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
