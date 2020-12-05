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

func step(key rune, min int, max int) (int, int) {
	middle := min + (max-min)/2

	if key == 'F' || key == 'L' {
		return min, middle
	}
	return middle + 1, max
}

func one() {
	absPath, _ := filepath.Abs("src/day05/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxId := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		minR := 0
		maxR := 127
		minC := 0
		maxC := 7
		line := scanner.Text()

		for pos, char := range line {
			if pos < 7 {
				minR, maxR = step(char, minR, maxR)
			} else {
				minC, maxC = step(char, minC, maxC)
			}
		}
		seatId := minR*8 + minC
		if seatId > maxId {
			maxId = seatId
		}

	}

	fmt.Println(maxId)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sumAll(min int, max int) int {
	return ((max - min) + 1) * (min + max) / 2
}

func two() {
	absPath, _ := filepath.Abs("src/day05/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	minId := 1000
	maxId := 0
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		minR := 0
		maxR := 127
		minC := 0
		maxC := 7
		line := scanner.Text()

		for pos, char := range line {
			if pos < 7 {
				minR, maxR = step(char, minR, maxR)
			} else {
				minC, maxC = step(char, minC, maxC)
			}
		}
		seatId := minR*8 + minC
		if seatId > maxId {
			maxId = seatId
		}
		if seatId < minId {
			minId = seatId
		}
		sum += seatId
	}

	fmt.Println(sumAll(minId, maxId) - sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
