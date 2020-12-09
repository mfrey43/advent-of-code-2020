package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	one()
	two()
}

func one() {
	absPath, _ := filepath.Abs("src/day09/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	last25 := make(map[int64]int)
	queue := make([]int64, 0)

	lineNum := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		if lineNum > 25 {
			match := false
			for _, prev := range queue {
				if last25[number-prev] > 0 {
					match = true
					break
				}
			}
			if !match {
				fmt.Println(number)
				break
			}

			out := queue[0]
			last25[out]--
			// Discard top element
			queue = queue[1:]
		}
		last25[number]++
		queue = append(queue, number)

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day09/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	last25 := make(map[int64]int)
	queue := make([]int64, 0)
	allNumbers := make([]int64, 0)
	var target int64

	lineNum := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		allNumbers = append(allNumbers, number)

		if lineNum > 25 {
			match := false
			for _, prev := range queue {
				if last25[number-prev] > 0 {
					match = true
					break
				}
			}
			if !match {
				target = number
				break
			}

			out := queue[0]
			last25[out]--
			// Discard top element
			queue = queue[1:]
		}
		last25[number]++
		queue = append(queue, number)

		lineNum++
	}

	// find range a to b
	a := 0
	b := 1
	sum := allNumbers[a] + allNumbers[b]
	for sum != target {
		if sum < target {
			b++
			sum += allNumbers[b]
		} else {
			sum -= allNumbers[a]
			a++
		}
	}

	// find min and max within range
	var max int64 = 0
	var min int64 = 9223372036854775807
	for i := a; i <= b; i++ {
		if allNumbers[i] < min {
			min = allNumbers[i]
		} else if allNumbers[i] > max {
			max = allNumbers[i]
		}
	}

	fmt.Println(min + max)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
