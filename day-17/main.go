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
	adv int = iota // 0
	bxl            // 1
	bst            // 2
	jnz            // 3
	bxc            // 4
	out            // 5
	bdv            // 6
	cdv            // 7
)

type Machine struct {
	A, B, C int
}

func main() {
	lines := readFile("./input.txt")

	A, _ := strconv.Atoi(lines[0][12:])
	B, _ := strconv.Atoi(lines[1][12:])
	C, _ := strconv.Atoi(lines[2][12:])

	split := strings.Split(lines[4][9:], ",")
	program := make([]int, 0)
	for _, op := range split {
		val, _ := strconv.Atoi(op)
		program = append(program, val)
	}

	machine := Machine{A, B, C}
	output := []string{}
	ip := 0
	for ip < len(program) {
		op := program[ip]
		if ip+1 >= len(program) {
			break
		}
		operand := program[ip+1]
		switch operand {
		case 4:
			operand = machine.A
		case 5:
			operand = machine.B
		case 6:
			operand = machine.C
		}

		switch op {
		case adv:
			machine.A = machine.A / (1 << operand)
		case bxl:
			machine.B ^= operand
		case bst:
			machine.B = operand % 8
		case jnz:
			if machine.A != 0 {
				ip = operand
				continue
			}
		case bxc:
			machine.B ^= machine.C
		case out:
			output = append(output, strconv.Itoa(operand%8))
		case bdv:
			machine.B = machine.A / (1 << operand)
		case cdv:
			machine.C = machine.A / (1 << operand)
		}

		ip += 2
	}

	part1 := strings.Join(output, ",")
	fmt.Println("Part 1:", part1)
	// TODO: Solve part 02
}
