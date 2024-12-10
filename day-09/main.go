package main

import (
	"fmt"
	"os"
	"slices"
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

const empty int = -1

type Block struct {
	id   int
	len  int
	free int
}

func parseBlocks(diskmap []string) []Block {
	blocks := make([]Block, 0)
	fileId := 0
	for i := range diskmap {
		if i%2 == 1 {
			continue
		}
		sFile := diskmap[i]
		file, _ := strconv.Atoi(sFile)
		block := Block{id: fileId, len: file}
		if i+1 < len(diskmap) {
			sFree := diskmap[i+1]
			free, _ := strconv.Atoi(sFree)
			block.free = free
		}
		blocks = append(blocks, block)
		fileId += 1
	}
	return blocks
}

func checksum(buf []int) uint64 {
	var sum uint64
	for i, v := range buf {
		if v == empty {
			continue
		}
		sum += uint64(i) * uint64(v)
	}
	return sum
}

func calcChecksum1(blocks []Block) uint64 {
	buf := make([]int, 0)
	spaces := make([]int, 0)
	for _, block := range blocks {
		for i := 0; i < block.len; i++ {
			buf = append(buf, block.id)
		}
		for i := 0; i < block.free; i++ {
			spaces = append(spaces, len(buf))
			buf = append(buf, empty)
		}
	}
	for j := len(buf) - 1; j >= 0; j-- {
		if buf[j] == empty {
			continue
		}
		i := spaces[0]
		spaces = spaces[1:]
		if i >= j {
			break
		}
		buf[i], buf[j] = buf[j], buf[i]
	}
	return checksum(buf)
}

func calcChecksum2(blocks []Block) uint64 {
	buf := make([]int, 0)
	spaces := make([]int, 0)
	for _, block := range blocks {
		for i := 0; i < block.len; i++ {
			buf = append(buf, block.id)
		}
		for i := 0; i < block.free; i++ {
			spaces = append(spaces, len(buf))
			buf = append(buf, empty)
		}
	}
	freeSpaces := make([][2]int, 0)
	index, count := -1, 0
	for i := len(spaces) - 1; i >= 0; i-- {
		if index != -1 && index-spaces[i] != 1 {
			freeSpaces = append(freeSpaces, [2]int{index, count})
			count = 0
		}
		index = spaces[i]
		count++
	}
	freeSpaces = append(freeSpaces, [2]int{index, count})
	slices.Reverse(freeSpaces)
	size := 0
	for i := len(blocks) - 1; i >= 0; i-- {
		block := blocks[i]
		size += block.free + block.len
		for idx := range freeSpaces {
			if len(buf)-size <= freeSpaces[idx][0] {
				break
			}
			if freeSpaces[idx][1]-block.len >= 0 {
				freeSpaces[idx][1] -= block.len
				for j := 0; j < block.len; j++ {
					buf[freeSpaces[idx][0]] = block.id
					buf[len(buf)-size+j] = empty
					freeSpaces[idx][0] += 1
				}
				break
			}
		}
	}
	return checksum(buf)
}

func main() {
	lines := readFile("./input.txt")

	blocks := parseBlocks(strings.Split(lines[0], ""))

	part1 := calcChecksum1(blocks)
	part2 := calcChecksum2(blocks)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
