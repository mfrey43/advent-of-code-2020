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
	absPath, _ := filepath.Abs("src/day16/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[int]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		ranges := strings.Split(strings.Split(line, ": ")[1], " or ")
		firstRange := strings.Split(ranges[0], "-")
		secondRange := strings.Split(ranges[1], "-")
		startA, _ := strconv.Atoi(firstRange[0])
		endA, _ := strconv.Atoi(firstRange[1])
		startB, _ := strconv.Atoi(secondRange[0])
		endB, _ := strconv.Atoi(secondRange[1])

		for i := startA; i <= endA; i++ {
			m[i] = true
		}
		for i := startB; i <= endB; i++ {
			m[i] = true
		}
	}

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line == "nearby tickets:" {
			continue
		}

		for _, numStr := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(numStr)
			if !m[num] {
				sum += num
			}
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day16/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[int][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		split := strings.Split(line, ": ")
		fieldName := split[0]
		ranges := strings.Split(split[1], " or ")
		firstRange := strings.Split(ranges[0], "-")
		secondRange := strings.Split(ranges[1], "-")
		startA, _ := strconv.Atoi(firstRange[0])
		endA, _ := strconv.Atoi(firstRange[1])
		startB, _ := strconv.Atoi(secondRange[0])
		endB, _ := strconv.Atoi(secondRange[1])

		for i := startA; i <= endA; i++ {
			m[i] = append(m[i], fieldName)
		}
		for i := startB; i <= endB; i++ {
			m[i] = append(m[i], fieldName)
		}
	}

	var validTickets [][]int

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line == "nearby tickets:" {
			continue
		}

		valid := true
		var ticket []int
		for _, numStr := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(numStr)
			ticket = append(ticket, num)
			if len(m[num]) == 0 {
				valid = false
				break
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	for i := 0; i < len(validTickets[0]); i++ {

		possibleFieldTypes := sliceToStrBoolMap(m[validTickets[0][i]])
		for _, ticket := range validTickets {
			num := ticket[i]
			next := make(map[string]bool)

			for _, fieldType := range m[num] {
				if possibleFieldTypes[fieldType] {
					next[fieldType] = true
				}
			}
			possibleFieldTypes = next
		}

		// fmt.Println(strings.Join(mapToStringSlice(possibleFieldTypes), ","))
	}

	/*
		eliminate entries until you arrive at this list:
		class
		row
		departure time
		type
		price
		departure track
		zone
		wagon
		departure station
		arrival location
		departure location
		arrival station
		route
		arrival track
		departure date
		duration
		train
		arrival platform
		departure platform
		seat

	*/

	// 3, 6, 9, 11, 15, 19
	// 109,199,223,179,97,227,197,151,73,79,211,181,71,139,53,149,137,191,83,193

	fmt.Println(223 * 227 * 73 * 211 * 53 * 83)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sliceToStrBoolMap(elements []string) map[string]bool {
	elementMap := make(map[string]bool)
	for _, s := range elements {
		elementMap[s] = true
	}
	return elementMap
}

func mapToStringSlice(m map[string]bool) []string {
	v := make([]string, 0, len(m))

	for key := range m {
		v = append(v, key)
	}

	return v
}
