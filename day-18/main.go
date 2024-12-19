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

type Coord struct {
	x, y int
}

var directions = [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func countSteps(size int, start Coord, end Coord, corrupted map[Coord]bool) int {
	queue := make([]Coord, 0)
	queue = append(queue, start)
	dist := make(map[Coord]int)
	dist[start] = 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == end {
			return dist[end]
		}
		for _, dir := range directions {
			nxt := Coord{cur.x + dir[0], cur.y + dir[1]}
			if nxt.x < 0 || nxt.x >= size || nxt.y < 0 || nxt.y >= size || corrupted[nxt] {
				continue
			}
			if d, ok := dist[nxt]; !ok || d > dist[cur]+1 {
				dist[nxt] = dist[cur] + 1
				queue = append(queue, nxt)
			}
		}
	}
	return -1
}

func findCorrupted(size int, start Coord, end Coord, corruptedList []Coord, corrupted map[Coord]bool) Coord {
	lo, hi := 0, len(corruptedList)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		for i := mid + 1; i < len(corruptedList); i++ {
			corrupted[corruptedList[i]] = false
		}
		steps := countSteps(size, start, end, corrupted)
		if steps == -1 {
			hi = mid
		} else {
			lo = mid + 1
		}
		for i := mid + 1; i < len(corruptedList); i++ {
			corrupted[corruptedList[i]] = true
		}
	}
	return corruptedList[lo]
}

func main() {
	lines := readFile("./input.txt")

	corrupted1024 := make(map[Coord]bool)
	corruptedFull := make(map[Coord]bool)
	corruptedList := make([]Coord, 0)
	for i, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		if i < 1024 {
			corrupted1024[Coord{x, y}] = true
		}
		corruptedFull[Coord{x, y}] = true
		corruptedList = append(corruptedList, Coord{x, y})
	}

	size := 71
	start := Coord{0, 0}
	end := Coord{70, 70}

	part1 := countSteps(size, start, end, corrupted1024)
	part2 := findCorrupted(size, start, end, corruptedList, corruptedFull)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
