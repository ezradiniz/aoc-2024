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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var directions = [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

type Pos struct {
	x, y int
}

func (p Pos) isOutOfBounds(lim Pos) bool {
	return p.x < 0 || p.x >= lim.x || p.y < 0 || p.y >= lim.y
}

func parseInput(lines []string) (map[Pos]rune, Pos, Pos, Pos) {
	grid := make(map[Pos]rune)
	lim := Pos{len(lines), len(lines[0])}
	start, end := Pos{}, Pos{}
	for x, row := range lines {
		for y, col := range row {
			pos := Pos{x, y}
			grid[pos] = col
			if col == 'S' {
				start = pos
			} else if col == 'E' {
				end = pos
			}
		}
	}
	return grid, lim, start, end
}

func getDist(grid map[Pos]rune, lim, start, end Pos) map[Pos]int {
	queue := make([]Pos, 0)
	queue = append(queue, start)
	dist := make(map[Pos]int)
	dist[queue[0]] = 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == end {
			return dist
		}
		for _, dir := range directions {
			nxt := Pos{cur.x + dir[0], cur.y + dir[1]}
			if nxt.isOutOfBounds(lim) || grid[nxt] == '#' {
				continue
			}
			if d, ok := dist[nxt]; !ok || d > dist[cur]+1 {
				dist[nxt] = dist[cur] + 1
				queue = append(queue, nxt)
			}
		}
	}
	return map[Pos]int{}
}

func countCheats(grid map[Pos]rune, lim, start, end Pos, cheatDist int) int {
	ans := 0
	dist1 := getDist(grid, lim, start, end)
	dist2 := getDist(grid, lim, end, start)
	baseline := dist1[end]
	for p1, d1 := range dist1 {
		for p2, d2 := range dist2 {
			md := abs(p1.x-p2.x) + abs(p1.y-p2.y)
			save := d1 + d2 + md
			if md <= cheatDist && baseline-save >= 100 {
				ans += 1
			}
		}
	}
	return ans
}

func main() {
	lines := readFile("./input.txt")
	grid, lim, start, end := parseInput(lines)

	part1 := countCheats(grid, lim, start, end, 2)
	part2 := countCheats(grid, lim, start, end, 20)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
