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

var directions = [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func getTotalCost1(grid []string) int {
	m, n := len(grid), len(grid[0])
	seen := make(map[Pos]bool)

	var rec func(pos Pos) (int, int)
	rec = func(pos Pos) (int, int) {
		seen[pos] = true
		area, perimeter := 1, 4
		for _, dir := range directions {
			nx, ny := pos.x+dir[0], pos.y+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] != grid[pos.x][pos.y] {
				continue
			}
			perimeter -= 1
			if ok := seen[Pos{nx, ny}]; ok {
				continue
			}
			a, p := rec(Pos{nx, ny})
			area += a
			perimeter += p
		}
		return area, perimeter
	}

	cost := 0
	for x := range m {
		for y := range n {
			if _, ok := seen[Pos{x, y}]; ok {
				continue
			}
			a, p := rec(Pos{x, y})
			cost += a * p
		}
	}
	return cost
}

type PosDir struct {
	x, y, i int
}

func getTotalCost2(grid []string) int {
	m, n := len(grid), len(grid[0])
	seen := make(map[Pos]bool)

	var rec func(pos Pos, corners map[PosDir]bool) int
	rec = func(pos Pos, corners map[PosDir]bool) int {
		seen[pos] = true
		area := 1
		for i, dir := range directions {
			nx, ny := pos.x+dir[0], pos.y+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] != grid[pos.x][pos.y] {
				corners[PosDir{x: pos.x, y: pos.y, i: i}] = true
				continue
			}
			if ok := seen[Pos{nx, ny}]; ok {
				continue
			}
			area += rec(Pos{nx, ny}, corners)
		}
		return area
	}

	cost := 0
	for x := range m {
		for y := range n {
			if _, ok := seen[Pos{x, y}]; ok {
				continue
			}
			corners := make(map[PosDir]bool)
			area := rec(Pos{x, y}, corners)
			side := 0
			for corner := range corners {
				c := corner
				c.y -= 1
				if _, ok := corners[c]; !ok && c.i != 1 && c.i != 3 {
					side++
				}
				c = corner
				c.x -= 1
				if _, ok := corners[c]; !ok && c.i != 0 && c.i != 2 {
					side++
				}
			}
			cost += area * side
		}
	}
	return cost
}

func main() {
	lines := readFile("./input.txt")

	part1 := getTotalCost1(lines)
	part2 := getTotalCost2(lines)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
