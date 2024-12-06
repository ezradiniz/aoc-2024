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

var direction = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func isOutOfBounds(x, y, rows, cols int) bool {
	return x < 0 || x >= rows || y < 0 || y >= cols
}

func getInitialPos(grid []string) (int, int) {
	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if grid[i][j] == '^' {
				return i, j
			}
		}
	}
	return 0, 0
}

func countPositions(grid []string, x, y, d int) int {
	m, n := len(grid), len(grid[0])
	seen := make([][]bool, m)
	for i := range m {
		seen[i] = make([]bool, n)
	}
	count := 0
	for !isOutOfBounds(x, y, m, n) {
		if !seen[x][y] {
			count += 1
		}
		seen[x][y] = true
		nx, ny := x+direction[d][0], y+direction[d][1]
		for !isOutOfBounds(nx, ny, m, n) && grid[nx][ny] == '#' {
			d = (d + 1) % len(direction)
			nx, ny = x+direction[d][0], y+direction[d][1]
		}
		x, y = nx, ny
	}
	return count
}

func countLoops(grid []string, startX, startY, dir int) int {
	m, n := len(grid), len(grid[0])
	seen := make([][][4]bool, m)
	for i := range m {
		seen[i] = make([][4]bool, n)
	}
	var recurse func(x, y, d, obsX, obsY int) int
	recurse = func(x, y, d, obsX, obsY int) int {
		if isOutOfBounds(x, y, m, n) {
			return 0
		}
		if seen[x][y][d] {
			return 1
		}
		pd := d
		seen[x][y][pd] = true
		nx, ny := x+direction[d][0], y+direction[d][1]
		for !isOutOfBounds(nx, ny, m, n) && (grid[nx][ny] == '#' || (nx == obsX && ny == obsY)) {
			d = (d + 1) % len(direction)
			nx, ny = x+direction[d][0], y+direction[d][1]
		}
		count := recurse(nx, ny, d, obsX, obsY)
		seen[x][y][pd] = false
		return count
	}
	// Brute force
	var ans int
	for i := range m {
		for j := range n {
			if grid[i][j] == '.' {
				ans += recurse(startX, startY, dir, i, j)
			}
		}
	}
	return ans
}

func main() {
	grid := readFile("./input.txt")

	x, y := getInitialPos(grid)
	part1 := countPositions(grid, x, y, 0)
	part2 := countLoops(grid, x, y, 0)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
