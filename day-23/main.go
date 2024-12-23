package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
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

type Graph = map[string][]string

func createGraph(lines []string) Graph {
	graph := Graph{}
	for _, line := range lines {
		split := strings.Split(line, "-")
		from, to := split[0], split[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	return graph
}

func getSetsOfThree(graph Graph) map[string]bool {
	triad := make(map[string]bool)
	for cur := range graph {
		for _, nxt := range graph[cur] {
			for _, nxtnxt := range graph[nxt] {
				if slices.Contains(graph[cur], nxtnxt) {
					group := []string{cur, nxt, nxtnxt}
					slices.Sort(group)
					triad[strings.Join(group, ",")] = true
				}
			}
		}
	}
	return triad
}

func getCliques(graph Graph) [][]string {
	cliques := [][]string{}

	r := make(map[string]bool)
	p := make(map[string]bool)
	x := make(map[string]bool)

	for node := range graph {
		p[node] = true
	}

	var intersect func(set map[string]bool, neighbors []string) map[string]bool
	intersect = func(set map[string]bool, neighbors []string) map[string]bool {
		result := make(map[string]bool)
		for _, neighbor := range neighbors {
			if set[neighbor] {
				result[neighbor] = true
			}
		}
		return result
	}

	var rec func(r, p, x map[string]bool)
	rec = func(r, p, x map[string]bool) {
		if len(p) == 0 && len(x) == 0 {
			clique := []string{}
			for node := range r {
				clique = append(clique, node)
			}
			cliques = append(cliques, clique)
			return
		}
		for node := range maps.Clone(p) {
			newR := maps.Clone(r)
			newR[node] = true
			newP := intersect(p, graph[node])
			newX := intersect(x, graph[node])
			rec(newR, newP, newX)
			delete(p, node)
			x[node] = true
		}
	}

	rec(r, p, x)
	return cliques
}

func main() {
	lines := readFile("./input.txt")

	graph := createGraph(lines)

	part1 := 0
	for triad := range getSetsOfThree(graph) {
		for _, node := range strings.Split(triad, ",") {
			if node[0] == 't' {
				part1++
				break
			}
		}
	}

	largestClique := []string{}
	for _, clique := range getCliques(graph) {
		if len(clique) > len(largestClique) {
			largestClique = clique
		}
	}
	slices.Sort(largestClique)
	part2 := strings.Join(largestClique, ",")

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
