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

func countWords(lines []string, word string) int {
	rows, cols := len(lines), len(lines[0])
	directions := [][2]int{
		{0, 1}, {0, -1}, {-1, 0}, {1, 0},
		{-1, 1}, {-1, -1}, {1, 1}, {1, -1},
	}
	count := 0
	for i := range rows {
		for j := range cols {
			for _, dir := range directions {
				idx := 0
				ni, nj := i, j
				di, dj := dir[0], dir[1]
				for idx < len(word) && ni >= 0 && ni < rows && nj >= 0 && nj < cols {
					if lines[ni][nj] != word[idx] {
						break
					}
					ni, nj = ni+di, nj+dj
					idx += 1
				}
				if idx >= len(word) {
					count += 1
				}
			}
		}
	}
	return count
}

func countShapes(lines []string, xShapes [][]string) int {
	rows, cols := len(lines), len(lines[0])
	isValid := func(r, c int, shape []string) bool {
		for i := range len(shape) {
			for j := range len(shape[0]) {
				if shape[i][j] == '.' {
					continue
				}
				if lines[r+i][c+j] != shape[i][j] {
					return false
				}
			}
		}
		return true
	}
	count := 0
	for i := range rows {
		for j := range cols {
			for _, shape := range xShapes {
				m, n := len(shape), len(shape[0])
				if i+m > rows || j+n > cols {
					continue
				}
				if isValid(i, j, shape) {
					count += 1
					break
				}
			}
		}
	}
	return count
}

func main() {
	lines := readFile("./input.txt")

	part1 := countWords(lines, "XMAS")
	part2 := countShapes(lines, [][]string{
		{"M.S", ".A.", "M.S"},
		{"S.M", ".A.", "S.M"},
		{"M.M", ".A.", "S.S"},
		{"S.S", ".A.", "M.M"},
	})

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
