package main

import (
	"errors"
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

type Reader struct {
	idx   int
	count int
	line  string
}

func (r *Reader) Lookahead() (byte, error) {
	if r.idx+1 >= len(r.line) {
		return 0, errors.New("reader: could not lookahead")
	}
	return r.line[r.idx+1], nil
}

func (r *Reader) Read() (byte, error) {
	if r.idx >= len(r.line) {
		return 0, errors.New("reader: eof")
	}
	ch := r.line[r.idx]
	r.idx++
	r.count++
	return ch, nil
}

func (r *Reader) Checkpoint() {
	r.count = 0
}

func (r *Reader) Rewind() {
	if r.count > r.idx {
		r.idx = 0
	} else {
		r.idx -= r.count
	}
	r.count = 0
}

func (r *Reader) Done() bool {
	return r.idx >= len(r.line)
}

var MulErr = errors.New("mul: could not parse mul")

type Mul struct {
	a int
	b int
}

func parseMul(r *Reader) (Mul, error) {
	r.Checkpoint()
	mul := Mul{}
	const term = "mul"
	for i := 0; i < len(term); i++ {
		ch, err := r.Read()
		if err != nil || ch != term[i] {
			return mul, MulErr
		}
	}
	if ch, err := r.Read(); err != nil || ch != '(' {
		return mul, MulErr
	}
	count := 0
	for {
		ch, err := r.Read()
		if err != nil {
			return mul, err
		}
		if ch == ',' {
			count++
			continue
		}
		if ch == ')' {
			count++
			break
		}
		digit, err := strconv.Atoi(string(ch))
		if err != nil {
			return mul, MulErr
		}
		if count == 0 {
			mul.a = mul.a*10 + digit
		} else if count == 1 {
			mul.b = mul.b*10 + digit
		}
	}
	if count < 2 {
		return Mul{}, MulErr
	}
	return mul, nil
}

func parseInstruction(r *Reader, instruction string) error {
	r.Checkpoint()
	for i := 0; i < len(instruction); i++ {
		ch, err := r.Read()
		if err != nil || ch != instruction[i] {
			return errors.New("parse: could not parse instruction")

		}
	}
	return nil
}

func solve1(lines []string) int {
	ans := 0
	for _, line := range lines {
		r := &Reader{line: line}
		for !r.Done() {
			if mul, err := parseMul(r); err == nil {
				ans += mul.a * mul.b
			}
		}
	}
	return ans
}

func solve2(lines []string) int {
	ans := 0
	enabled := true
	for _, line := range lines {
		r := &Reader{line: line}
		for !r.Done() {
			if mul, err := parseMul(r); err == nil {
				if enabled {
					ans += mul.a * mul.b
				}
				continue
			}
			r.Rewind()
			if err := parseInstruction(r, "do()"); err == nil {
				enabled = true
				continue
			}
			r.Rewind()
			if err := parseInstruction(r, "don't()"); err == nil {
				enabled = false
			}
		}
	}
	return ans
}

func main() {
	lines := readFile("./input.txt")

	part1 := solve1(lines)
	part2 := solve2(lines)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
