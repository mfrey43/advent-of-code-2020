package main

import "fmt"

func main() {
	one()
	two()
}

func one() {
	fmt.Println(xthNumberSpoken([]int{19, 0, 5, 1, 10, 13}, 2020))
}

func two() {
	fmt.Println(xthNumberSpoken([]int{19, 0, 5, 1, 10, 13}, 30000000))
}

func xthNumberSpoken(input []int, turn int) int {

	m := make(map[int][]int)

	for i := 0; i < len(input); i++ {
		m[input[i]] = []int{i + 1, 0}
	}

	last := input[len(input)-1]
	for i := len(input) + 1; i <= turn; i++ {
		speak := 0
		if m[last][0] > 0 && m[last][1] > 0 {
			speak = m[last][0] - m[last][1]
		}
		if len(m[speak]) == 0 {
			m[speak] = []int{0, 0}
		}
		if m[speak][0] > 0 {
			m[speak][1] = m[speak][0]
		}
		m[speak][0] = i
		last = speak
	}

	return last
}
