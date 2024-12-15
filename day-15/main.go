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

type Dir string

func (d Dir) ToVec(rev bool) [2]int {
	offset := 0
	if rev {
		offset = 2
	}
	switch d {
	case "<":
		return directions[(1+offset)%4]
	case ">":
		return directions[(3+offset)%4]
	case "v":
		return directions[(2+offset)%4]
	case "^":
		return directions[(0+offset)%4]
	default:
		return [2]int{-1, -1}
	}
}

type Grid [][]Cell

type Cell struct {
	token string
	dy    int
}

type Pos struct {
	x, y int
}

func isOutOfBounds(grid Grid, pos Pos) bool {
	return pos.x < 0 || pos.x >= len(grid) || pos.y < 0 || pos.y >= len(grid[0]) || grid[pos.x][pos.y].token == "#"
}

func moveTo(grid Grid, pos Pos, dir Dir) Pos {
	cdir := dir.ToVec(false)
	newPos := Pos{pos.x + cdir[0], pos.y + cdir[1]}
	if isOutOfBounds(grid, pos) {
		return pos
	}
	if grid[newPos.x][newPos.y].token == "O" {
		return moveBox(grid, newPos, dir)
	}
	if grid[newPos.x][newPos.y].token == "." {
		grid[newPos.x][newPos.y], grid[pos.x][pos.y] = grid[pos.x][pos.y], grid[newPos.x][newPos.y]
		return newPos
	}
	return pos
}

func moveBox(grid Grid, pos Pos, dir Dir) Pos {
	seen := make(map[Pos]bool)
	want := make([]Pos, 0)
	ndir := dir.ToVec(false)
	pdir := dir.ToVec(true)
	canMove := true

	var dfs func(cur Pos)
	dfs = func(cur Pos) {
		if !canMove {
			return
		}
		if grid[cur.x][cur.y].token == "." {
			want = append(want, cur)
			return
		}
		for _, dy := range [2]int{0, grid[cur.x][cur.y].dy} {
			nxt := Pos{cur.x + ndir[0], cur.y + ndir[1] + dy}
			if isOutOfBounds(grid, nxt) {
				canMove = false
				continue
			}
			if _, ok := seen[nxt]; ok {
				continue
			}
			seen[nxt] = true
			dfs(nxt)
		}
	}

	seen[pos] = true
	dfs(pos)

	if !canMove {
		return Pos{pos.x + pdir[0], pos.y + pdir[1]}
	}

	for len(want) > 0 {
		nxtWant := make([]Pos, 0)
		for _, p := range want {
			grid[p.x][p.y] = grid[p.x+pdir[0]][p.y+pdir[1]]
			p.x += pdir[0]
			p.y += pdir[1]
			grid[p.x][p.y] = Cell{token: ".", dy: 0}
			if _, ok := seen[p]; ok {
				nxtWant = append(nxtWant, p)
			}
		}
		want = nxtWant
	}
	return pos
}

func getInitialPos(grid Grid) Pos {
	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j].token == "@" {
				return Pos{i, j}
			}
		}
	}
	return Pos{0, 0}
}

func sumGPSCoord(grid Grid, wide bool) int {
	sum := 0
	for i := range len(grid) {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].token == "O" {
				sum += 100*i + j
				if wide {
					j += 1
				}
			}
		}
	}
	return sum
}

func parseCell(row []string, wide bool) []Cell {
	cells := make([]Cell, 0)
	for _, r := range row {
		if r == "@" {
			if wide {
				cells = append(cells, Cell{token: r, dy: 0})
				cells = append(cells, Cell{token: ".", dy: 0})
			} else {
				cells = append(cells, Cell{token: r, dy: 0})
			}
		} else {
			if wide {
				cells = append(cells, Cell{token: r, dy: 1})
				cells = append(cells, Cell{token: r, dy: -1})
			} else {
				cells = append(cells, Cell{token: r, dy: 0})
			}
		}
	}
	return cells
}

func main() {
	lines := readFile("./input.txt")

	grid1 := make([][]Cell, 0)
	grid2 := make([][]Cell, 0)

	parseMoves := false
	moves := make([]string, 0)
	for _, line := range lines {
		split := strings.Split(line, "")
		if line == "" {
			parseMoves = true
			continue
		}
		if parseMoves {
			moves = append(moves, split...)
		} else {
			grid1 = append(grid1, parseCell(split, false))
			grid2 = append(grid2, parseCell(split, true))
		}
	}

	pos1 := getInitialPos(grid1)
	pos2 := getInitialPos(grid2)
	for _, move := range moves {
		pos1 = moveTo(grid1, pos1, Dir(move))
		pos2 = moveTo(grid2, pos2, Dir(move))
	}

	part1 := sumGPSCoord(grid1, false)
	part2 := sumGPSCoord(grid2, true)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
