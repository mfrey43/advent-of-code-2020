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
	var stack []instruction

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
	stacking := true
	for true {
		if pointer >= len(lines) {
			break
		}
		instr := lines[pointer]

		if seen[pointer] {
			if stacking {
				stacking = false
			}
			// backtrack to the last jmp or nop command
			for i := len(stack) - 1; i >= 0; i-- {
				seen[stack[i].line] = false
				if stack[i].cmd == "jmp" || stack[i].cmd == "nop" {
					pointer = stack[i].line
					instr = stack[i]
					acc = stack[i].acc
					if instr.cmd == "jmp" {
						instr.cmd = "nop"
					} else if instr.cmd == "nop" {
						instr.cmd = "jmp"
					}
					stack = stack[:i]
					break
				}
			}
		}
		seen[pointer] = true

		if instr.cmd == "acc" {
			acc += instr.num
			pointer++
		} else if instr.cmd == "jmp" {
			pointer += instr.num
		} else if instr.cmd == "nop" {
			pointer++
		}
		instr.acc = acc

		if stacking {
			stack = append(stack, instr)
		}
	}

	fmt.Println(acc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
