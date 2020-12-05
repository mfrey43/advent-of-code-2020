package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	one()
	two()
}

func one() {
	absPath, _ := filepath.Abs("src/day03/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	trees := 0

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	width := len(scanner.Text())
	nextPos := 3
	for scanner.Scan() {
		line := scanner.Text()
		if line[nextPos] == '#' {
			trees++
		}
		nextPos = (nextPos + 3) % width
	}

	fmt.Println(trees)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day03/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	trees1 := 0
	trees2 := 0
	trees3 := 0
	trees4 := 0
	trees5 := 0

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	width := len(scanner.Text())
	nextPos1 := 1
	nextPos2 := 3
	nextPos3 := 5
	nextPos4 := 7
	nextPos5 := 1
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line[nextPos1] == '#' {
			trees1++
		}
		nextPos1 = (nextPos1 + 1) % width
		if line[nextPos2] == '#' {
			trees2++
		}
		nextPos2 = (nextPos2 + 3) % width
		if line[nextPos3] == '#' {
			trees3++
		}
		nextPos3 = (nextPos3 + 5) % width
		if line[nextPos4] == '#' {
			trees4++
		}
		nextPos4 = (nextPos4 + 7) % width

		if lineNum%2 == 0 {
			if line[nextPos5] == '#' {
				trees5++
			}
			nextPos5 = (nextPos5 + 1) % width
		}
		lineNum++
	}

	fmt.Println(trees1 * trees2 * trees3 * trees4 * trees5)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
