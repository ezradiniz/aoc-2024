package main

import (
	"container/heap"
	"fmt"
	"math"
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
	x, y, dir int
}

type Item struct {
	Pos
	score int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].score < pq[j].score }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any)        { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

var directions = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func isOutOfBounds(maze []string, x, y int) bool {
	return x < 0 || y < 0 || x >= len(maze) || y >= len(maze[0]) || maze[x][y] == '#'
}

func findScoreAndTiles(maze []string, start Pos, end Pos) (int, int) {
	inf := math.MaxInt32
	points := make([][][4]int, len(maze))
	for i := range maze {
		points[i] = make([][4]int, len(maze[0]))
		for j := range maze[0] {
			for d := 0; d < 4; d++ {
				points[i][j][d] = inf
			}
		}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{Pos: start, score: 0})
	points[start.x][start.y][start.dir] = 0

	prev := make(map[Pos][]Pos)
	seen := make(map[Pos]bool)
	var buildPath func(pos Pos, tiles map[Pos]bool)
	buildPath = func(pos Pos, tiles map[Pos]bool) {
		tiles[Pos{pos.x, pos.y, 0}] = true
		for _, nxt := range prev[pos] {
			if _, ok := seen[nxt]; !ok {
				seen[nxt] = true
				buildPath(nxt, tiles)
			}
		}
	}

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		x, y, d, curScore := cur.x, cur.y, cur.dir, cur.score

		if points[x][y][d] < curScore {
			continue
		}

		if x == end.x && y == end.y {
			tiles := make(map[Pos]bool)
			buildPath(cur.Pos, tiles)
			return curScore, len(tiles)
		}

		nx, ny := x+directions[d][0], y+directions[d][1]
		if !isOutOfBounds(maze, nx, ny) {
			newScore := curScore + 1
			if points[nx][ny][d] > newScore {
				points[nx][ny][d] = newScore
				heap.Push(&pq, &Item{Pos: Pos{nx, ny, d}, score: newScore})
				prev[Pos{nx, ny, d}] = []Pos{cur.Pos}
			} else if points[nx][ny][d] == newScore {
				prev[Pos{nx, ny, d}] = append(prev[Pos{nx, ny, d}], cur.Pos)
			}
		}

		for i := -1; i <= 1; i += 2 {
			nd := (d + i + 4) % 4
			newScore := curScore + 1000
			if points[x][y][nd] > newScore {
				points[x][y][nd] = newScore
				heap.Push(&pq, &Item{Pos: Pos{x, y, nd}, score: newScore})
				prev[Pos{x, y, nd}] = []Pos{cur.Pos}
			} else if points[x][y][nd] == newScore {
				prev[Pos{x, y, nd}] = append(prev[Pos{x, y, nd}], cur.Pos)
			}
		}
	}

	return -1, 0
}

func getInitPos(maze []string, target rune) Pos {
	for x, row := range maze {
		for y, col := range row {
			if col == target {
				return Pos{x, y, 0}
			}
		}
	}
	return Pos{0, 0, 0}
}

func main() {
	maze := readFile("./input.txt")

	start := getInitPos(maze, 'S')
	end := getInitPos(maze, 'E')
	score, tiles := findScoreAndTiles(maze, start, end)

	fmt.Println("Part 1:", score)
	fmt.Println("Part 2:", tiles)
}
