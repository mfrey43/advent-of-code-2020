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
	absPath, _ := filepath.Abs("src/day06/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	m := make(map[rune]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			sum += len(m)
			m = make(map[rune]bool)
		} else {
			for _, char := range line {
				m[char] = true
			}
		}
	}

	sum += len(m) // add last group
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day06/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	minMap := make(map[rune]bool) // questions everyone in the group answered yes to
	first := true                 // true = parsing first person in group

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			sum += len(minMap)
			minMap = make(map[rune]bool)
			first = true
		} else {
			if first {
				for _, char := range line {
					minMap[char] = true
				}
				first = false
			} else {
				m := make(map[rune]bool)
				for _, char := range line {
					if minMap[char] {
						m[char] = true
					}
				}
				minMap = m
			}
		}
	}

	sum += len(minMap) // add last group
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
