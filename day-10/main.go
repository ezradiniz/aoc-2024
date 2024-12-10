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

var directions = [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

type Pos struct {
	x, y int
}

func walk(topoMap []string, reachable map[Pos]bool, pos Pos) int {
	if topoMap[pos.x][pos.y] == '9' {
		reachable[pos] = true
		return 1
	}
	count := 0
	for _, dir := range directions {
		nx, ny := pos.x+dir[0], pos.y+dir[1]
		if nx < 0 || nx >= len(topoMap) || ny < 0 || ny >= len(topoMap[0]) {
			continue
		}
		if topoMap[nx][ny]-topoMap[pos.x][pos.y] != 1 {
			continue
		}
		count += walk(topoMap, reachable, Pos{x: nx, y: ny})
	}
	return count
}

func countScores(topoMap []string) (int, int) {
	score1, score2 := 0, 0
	for x, row := range topoMap {
		for y, pos := range row {
			if pos == '0' {
				reachable := make(map[Pos]bool)
				count := walk(topoMap, reachable, Pos{x, y})
				score1 += len(reachable)
				score2 += count
			}
		}
	}
	return score1, score2
}

func main() {
	topoMap := readFile("./input.txt")

	part1, part2 := countScores(topoMap)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
