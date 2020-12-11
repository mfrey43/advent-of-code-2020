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
	absPath, _ := filepath.Abs("src/day11/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var m [][]rune

	scanner := bufio.NewScanner(file)

	r := 0
	for scanner.Scan() {
		m = append(m, make([]rune, 0))
		line := scanner.Text()
		c := 0
		for _, char := range line {
			m[r] = append(m[r], char)
			c++
		}
		r++
	}

	m2 := make([][]rune, len(m))
	for r = 0; r < len(m); r++ {
		m2[r] = make([]rune, len(m[r]))
		for c := 0; c < len(m[r]); c++ {
			m2[r][c] = m[r][c]
		}
	}

	from := m
	to := m2
	hasChanged := true
	for hasChanged {
		hasChanged = false
		for r = 0; r < len(from); r++ {
			for c := 0; c < len(from[r]); c++ {
				// fmt.Print(string(from[r][c]))
				if from[r][c] == '.' {
					continue
				}
				to[r][c] = from[r][c]
				numAdjacent := getNumAdjacent(from, r, c)
				if numAdjacent == 0 {
					to[r][c] = '#'
				} else if numAdjacent >= 4 {
					to[r][c] = 'L'
				}
				if from[r][c] != to[r][c] {
					hasChanged = true
				}
			}
			// fmt.Println()
		}
		tmp := from
		from = to
		to = tmp
	}

	count := 0
	for r = 0; r < len(from); r++ {
		for c := 0; c < len(from[r]); c++ {
			if to[r][c] == '#' {
				count++
			}
		}
	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func safeIsTaken(m [][]rune, r int, c int) int {
	if r >= 0 && r < len(m) && c >= 0 && c < len(m[r]) && m[r][c] == '#' {
		return 1
	}
	return 0
}

func getNumAdjacent(m [][]rune, r int, c int) int {
	return safeIsTaken(m, r-1, c-1) + safeIsTaken(m, r-1, c) + safeIsTaken(m, r-1, c+1) +
		safeIsTaken(m, r, c-1) + safeIsTaken(m, r, c+1) +
		safeIsTaken(m, r+1, c-1) + safeIsTaken(m, r+1, c) + safeIsTaken(m, r+1, c+1)
}

func two() {
	absPath, _ := filepath.Abs("src/day11/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var m [][]rune

	scanner := bufio.NewScanner(file)

	r := 0
	for scanner.Scan() {
		m = append(m, make([]rune, 0))
		line := scanner.Text()
		c := 0
		for _, char := range line {
			m[r] = append(m[r], char)
			c++
		}
		r++
	}

	m2 := make([][]rune, len(m))
	for r = 0; r < len(m); r++ {
		m2[r] = make([]rune, len(m[r]))
		for c := 0; c < len(m[r]); c++ {
			m2[r][c] = m[r][c]
		}
	}

	from := m
	to := m2
	hasChanged := true
	for hasChanged {
		hasChanged = false
		for r = 0; r < len(from); r++ {
			for c := 0; c < len(from[r]); c++ {
				// fmt.Print(string(from[r][c]))
				if from[r][c] == '.' {
					continue
				}
				to[r][c] = from[r][c]
				numAdjacent := getOccSeatsSeen(from, r, c)
				if numAdjacent == 0 {
					to[r][c] = '#'
				} else if numAdjacent >= 5 {
					to[r][c] = 'L'
				}
				if from[r][c] != to[r][c] {
					hasChanged = true
				}
			}
			// fmt.Println()
		}
		tmp := from
		from = to
		to = tmp
	}

	count := 0
	for r = 0; r < len(from); r++ {
		for c := 0; c < len(from[r]); c++ {
			if to[r][c] == '#' {
				count++
			}
		}
	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func safeIsTaken2(m [][]rune, r int, c int) int {
	if r >= 0 && r < len(m) && c >= 0 && c < len(m[r]) {
		if m[r][c] == '#' {
			return 1
		} else if m[r][c] == 'L' {
			return -1
		}
		return 0 // floor
	}
	return -1 // no more seats
}

func getOccSeatsSeen(m [][]rune, r int, c int) int {
	total := 0
	dist := 1
	for true {
		res := safeIsTaken2(m, r-dist, c-dist)
		if res == 1 {
			total++
		}
		if res != 0 {
			break
		}
		dist++
	}
	dist = 1
	for true {
		res := safeIsTaken2(m, r-dist, c)
		if res == 1 {
			total++
		}
		if res != 0 {
			break
		}
		dist++
	}
	dist = 1
	for true {
		res := safeIsTaken2(m, r-dist, c+dist)
		if res == 1 {
			total++
		}
		if res != 0 {
			break
		}
		dist++
	}
	dist = 1
	for true {
		res := safeIsTaken2(m, r, c-dist)
		if res == 1 {
			total++
		}
		if res != 0 {
			break
		}
		dist++
	}
	dist = 1
	for true {
		res := safeIsTaken2(m, r, c+dist)
		if res == 1 {
			total++
		}
		if res != 0 {
			break
		}
		dist++
	}
	dist = 1
	for true {
		res := safeIsTaken2(m, r+dist, c-dist)
		if res == 1 {
			total++
		}
		if res != 0 {
			break
		}
		dist++
	}
	dist = 1
	for true {
		res := safeIsTaken2(m, r+dist, c)
		if res == 1 {
			total++
		}
		if res != 0 {
			break
		}
		dist++
	}
	dist = 1
	for true {
		res := safeIsTaken2(m, r+dist, c+dist)
		if res == 1 {
			total++
		}
		if res != 0 {
			break
		}
		dist++
	}
	return total
}
