package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	one()
	two()
}

func one() {
	absPath, _ := filepath.Abs("src/day08/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	acc := 0
	pointer := 0
	seen := make(map[int]bool)
	for true {
		line := lines[pointer]
		split := strings.Split(line, " ")
		cmd := split[0]
		num, _ := strconv.Atoi(split[1])

		if seen[pointer] {
			break
		}
		seen[pointer] = true

		if cmd == "acc" {
			acc += num
			pointer++
		} else if cmd == "jmp" {
			pointer += num
		} else if cmd == "nop" {
			pointer++
		}
	}

	fmt.Println(acc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type instruction struct {
	line int
	cmd  string
	num  int
	acc  int
}

func two() {
	absPath, _ := filepath.Abs("src/day08/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []instruction

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		cmd := split[0]
		num, _ := strconv.Atoi(split[1])
		lines = append(lines, instruction{lineNum, cmd, num, 0})
		lineNum++
	}

	acc := 0
	pointer := 0
	seen := make(map[int]bool)
	var stack []instruction
	stacking := true
	for true {
		if pointer >= len(lines) {
			break
		}
		instr := lines[pointer]

		if seen[pointer] { // we entered a loop
			if stacking {
				stacking = false
			}
			// go back to the last jmp or nop command
			instr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pointer = instr.line
			acc = instr.acc
			if instr.cmd == "jmp" {
				instr.cmd = "nop"
			} else if instr.cmd == "nop" {
				instr.cmd = "jmp"
			}
		}
		seen[pointer] = true

		if instr.cmd == "acc" {
			acc += instr.num
			pointer++
		} else if instr.cmd == "jmp" {
			pointer += instr.num

			if stacking {
				instr.acc = acc
				stack = append(stack, instr)
			}
		} else if instr.cmd == "nop" {
			pointer++

			if stacking {
				instr.acc = acc
				stack = append(stack, instr)
			}
		}
	}

	fmt.Println(acc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
