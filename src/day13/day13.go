package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	absPath, _ := filepath.Abs("src/day13/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	arrival, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	var buses []int
	for _, id := range strings.Split(scanner.Text(), ",") {
		if id != "x" {
			num, _ := strconv.Atoi(id)
			buses = append(buses, num)
		}
	}

	minId := 0
	min := math.MaxInt32
	for _, busId := range buses {
		waitTime := busId - (arrival % busId)
		if waitTime < min {
			min = waitTime
			minId = busId
		}
	}

	fmt.Println(min * minId)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day13/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	var buses []int
	var diffs []int
	for pos, id := range strings.Split(scanner.Text(), ",") {
		if id != "x" {
			num, _ := strconv.Atoi(id)
			buses = append(buses, num)
			diffs = append(diffs, pos)
		}
	}

	lcm := buses[0]
	result := buses[0]

	for i := 1; i < len(buses); i++ {
		busId := buses[i]
		diffs[i] = diffs[i] % busId
		diff := diffs[i]
		for true {
			if busId-(result%busId) == diff {
				break
			}
			result += lcm
		}
		lcm = LCM(lcm, busId)
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
