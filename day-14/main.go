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

const (
	rows int = 103
	cols int = 101
)

type Robot struct {
	x, y, dx, dy int
}

func parseRobot(line string) Robot {
	parts := strings.Split(line, " ")
	p1 := strings.Index(parts[0], ",")
	y, _ := strconv.Atoi(parts[0][2:p1])
	x, _ := strconv.Atoi(parts[0][p1+1:])
	p2 := strings.Index(parts[1], ",")
	dy, _ := strconv.Atoi(parts[1][2:p2])
	dx, _ := strconv.Atoi(parts[1][p2+1:])
	return Robot{x, y, dx, dy}
}

type Pos struct {
	x, y int
}

func getRobotPos(robot Robot, seconds int) Pos {
	robot.x = (robot.x + robot.dx*seconds) % rows
	if robot.x < 0 {
		robot.x += rows
	}
	robot.y = (robot.y + robot.dy*seconds) % cols
	if robot.y < 0 {
		robot.y += cols
	}
	return Pos{robot.x, robot.y}
}

func getQuadrant(pos Pos) int {
	dx, dy := rows/2, cols/2
	if pos.x >= 0 && pos.x < dx && pos.y >= 0 && pos.y < dy {
		return 0
	}
	if pos.x >= rows-dx && pos.x < rows && pos.y >= 0 && pos.y < dy {
		return 1
	}
	if pos.x >= 0 && pos.x < dx && pos.y >= cols-dy && pos.y < cols {
		return 2
	}
	if pos.x >= rows-dx && pos.x < rows && pos.y >= cols-dy && pos.y < cols {
		return 3
	}
	return -1
}

func getSafetyFactor(lines []string) int {
	seconds := 100
	counter := make(map[Pos]int)
	for _, line := range lines {
		robot := parseRobot(line)
		pos := getRobotPos(robot, seconds)
		counter[pos]++
	}
	quadrants := [4]int{0, 0, 0, 0}
	for pos, count := range counter {
		if quad := getQuadrant(pos); quad != -1 {
			quadrants[quad] += count
		}
	}
	factor := 1
	for _, count := range quadrants {
		factor *= count
	}
	return factor
}

type Grid [cols][rows]bool

func (g Grid) Display() {
	for j := range rows {
		for i := range cols {
			if !g[i][j] {
				fmt.Print(" ")
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}
}

func getMinTimeForEasterEgg(lines []string) int {
	// I tried to simulate for each second starting
	// from 4 and incrementing 103 (sec = 4; ; sec += 103)
	seconds := 6493
	grid := Grid{}
	for _, line := range lines {
		robot := parseRobot(line)
		pos := getRobotPos(robot, seconds)
		grid[pos.y][pos.x] = true
	}
	grid.Display()
	return seconds
}

func main() {
	lines := readFile("./input.txt")

	part1 := getSafetyFactor(lines)
	part2 := getMinTimeForEasterEgg(lines)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
