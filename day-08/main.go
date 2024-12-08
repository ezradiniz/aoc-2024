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

type Pos struct {
	x, y int
}

func isOutOfBounds(x, y, rows, cols int) bool {
	return x < 0 || x >= rows || y < 0 || y >= cols
}

func countAntinodes(lines []string, newModel bool) int {
	m, n := len(lines), len(lines[0])
	antinodes := make(map[Pos]bool)
	seen := make(map[byte][]Pos)
	for x := range m {
		for y := range n {
			if lines[x][y] == '.' {
				continue
			}
			if newModel {
				antinodes[Pos{x, y}] = true
			}
			if _, ok := seen[lines[x][y]]; !ok {
				seen[lines[x][y]] = make([]Pos, 0)
			}
			for _, pos := range seen[lines[x][y]] {
				dx := x - pos.x
				dy := y - pos.y
				dc := 1
				for !isOutOfBounds(x+dx*dc, y+dy*dc, m, n) {
					antinodes[Pos{x: x + dx*dc, y: y + dy*dc}] = true
					if !newModel {
						break
					}
					dc += 1
				}
				dc = 1
				for !isOutOfBounds(pos.x-dx*dc, pos.y-dy*dc, m, n) {
					antinodes[Pos{x: pos.x - dx*dc, y: pos.y - dy*dc}] = true
					if !newModel {
						break
					}
					dc += 1
				}
			}
			seen[lines[x][y]] = append(seen[lines[x][y]], Pos{x, y})
		}
	}
	return len(antinodes)
}
func main() {
	lines := readFile("./input.txt")

	part1 := countAntinodes(lines, false)
	part2 := countAntinodes(lines, true)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
