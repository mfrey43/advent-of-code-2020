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
	absPath, _ := filepath.Abs("src/day10/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[int]bool)

	max := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		m[num] = true

		if num > max {
			max = num
		}
	}

	lastAdapter := 0
	occurrences := make([]int, 3)

	jolt := 1
	for jolt <= max {
		if m[jolt] {
			occurrences[jolt-lastAdapter-1]++
			lastAdapter = jolt
		}
		jolt++
	}

	occurrences[2]++

	fmt.Println(occurrences[0] * occurrences[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func numWays(m map[int]bool, jolt int, max int, memo map[int]int) int {
	if memo[jolt] > 0 {
		return memo[jolt]
	}
	if jolt == max {
		return 1
	}
	if m[jolt] {
		memo[jolt] = numWays(m, jolt+1, max, memo) + numWays(m, jolt+2, max, memo) + numWays(m, jolt+3, max, memo)
		return memo[jolt]
	}
	return 0
}

func two() {
	absPath, _ := filepath.Abs("src/day10/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[int]bool)
	m[0] = true

	max := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		m[num] = true

		if num > max {
			max = num
		}
	}

	memo := make(map[int]int)
	fmt.Println(numWays(m, 0, max, memo))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
