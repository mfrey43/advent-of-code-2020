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
	absPath, _ := filepath.Abs("src/day02/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numValid := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		mode := 0
		var min int
		var max int
		var c rune
		count := 0

		var sb strings.Builder

		for _, rune := range line {
			if mode == 0 {
				if rune == '-' {
					min, _ = strconv.Atoi(sb.String())
					sb.Reset()
					mode = 1
				} else {
					sb.WriteRune(rune)
				}
			} else if mode == 1 {
				if rune == ' ' {
					max, _ = strconv.Atoi(sb.String())
					mode = 2
				} else {
					sb.WriteRune(rune)
				}
			} else if mode == 2 {
				if rune == ' ' {
					mode = 3
				} else if rune != ':' {
					c = rune
				}
			} else if mode == 3 {
				if rune == c {
					count++
				}
			}
		}
		if count <= max && count >= min {
			numValid++
		}
	}

	fmt.Println(numValid)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day02/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numValid := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		mode := 0
		var startPos int
		var posA int
		var posB int
		var c rune
		valid := false

		var sb strings.Builder

		for pos, rune := range line {
			if mode == 0 {
				if rune == '-' {
					posA, _ = strconv.Atoi(sb.String())
					sb.Reset()
					mode = 1
				} else {
					sb.WriteRune(rune)
				}
			} else if mode == 1 {
				if rune == ' ' {
					posB, _ = strconv.Atoi(sb.String())
					mode = 2
				} else {
					sb.WriteRune(rune)
				}
			} else if mode == 2 {
				if rune == ' ' {
					startPos = pos
					mode = 3
				} else if rune != ':' {
					c = rune
				}
			} else if mode == 3 {
				sPos := pos - startPos
				if rune == c && (sPos == posA || sPos == posB) {
					valid = !valid
				}
			}
		}
		if valid {
			numValid++
		}
	}

	fmt.Println(numValid)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
