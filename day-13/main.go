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

type Coor struct {
	x, y int
}

func parseButton(line string) Coor {
	x, _ := strconv.Atoi(line[strings.Index(line, "+")+1 : strings.Index(line, ",")])
	y, _ := strconv.Atoi(line[strings.LastIndex(line, "+")+1:])
	return Coor{x, y}
}

func parsePrize(line string, offset int) Coor {
	x, _ := strconv.Atoi(line[strings.Index(line, "=")+1 : strings.Index(line, ",")])
	y, _ := strconv.Atoi(line[strings.LastIndex(line, "=")+1:])
	return Coor{x + offset, y + offset}
}

const (
	costA int = 3
	costB     = 1
)

const inf int = 1e15

func countTokens(btnA, btnB, prize Coor) int {
	// btnA.x*a + btnB.x*b = prize.x
	// btnB.y*a + btnB.y*b = prize.y
	//
	// btnA.x*a = prize.x - btnB.x*b
	// a = (prize.x - btnB.x*b) / btnA.x
	//
	// btnB.y * (prize.x - btnB.x*b) / btnA.x + btnB.y*b = prize.y
	// btnB.y*b = prize.y - btnB.y * (prize.x - btnB.x*b) / btnA.x
	// b = (prize.y - btnB.y * (prize.x - btnB.x*b) / btnA.x) / btnB.y
	// b = (prize.y - btnA.y * prize.x / btnA.x) / (btnB.y - btnA.y * btnB.x / btnA.x)
	// a = -(prize.y * btnB.x - btnB.y * prize.x) / (btnB.y * btnA.x - btnA.y * btnB.x)
	//
	a := prize.x*btnB.y - prize.y*btnB.x
	b := prize.y*btnA.x - prize.x*btnA.y
	d := btnB.y*btnA.x - btnA.y*btnB.x
	if a%d != 0 || b%d != 0 {
		return inf
	}
	a /= d
	b /= d
	return costA*a + b*costB
}

func main() {
	lines := readFile("./input.txt")

	part1 := 0
	part2 := 0
	for i := 0; i < len(lines); i += 4 {
		btnA := parseButton(lines[i])
		btnB := parseButton(lines[i+1])
		prize1 := parsePrize(lines[i+2], 0)
		prize2 := parsePrize(lines[i+2], 10000000000000)
		if tokens1 := countTokens(btnA, btnB, prize1); tokens1 < inf {
			part1 += tokens1
		}
		if tokens2 := countTokens(btnA, btnB, prize2); tokens2 < inf {
			part2 += tokens2
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
