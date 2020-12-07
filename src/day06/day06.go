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
	absPath, _ := filepath.Abs("src/day06/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bagsByBag := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		mainSplit := strings.Split(line, "contain")
		if mainSplit[1] == " no other bags." {
			continue
		}
		container := strings.TrimSpace(mainSplit[0])
		container = container[:len(container)-5]

		for _, str := range strings.Split(mainSplit[1], ",") {
			split := strings.Split(str, " ")
			bag := strings.TrimSpace(split[2] + " " + split[3])
			bagsByBag[bag] = append(bagsByBag[bag], container)
		}
	}

	// look through
	queue := []string{"shiny gold"}
	checked := make(map[string]bool)

	for len(queue) > 0 {
		bag := queue[0]
		if !checked[bag] {
			checked[bag] = true
			for _, outerBag := range bagsByBag[bag] {
				queue = append(queue, outerBag)
			}
		}
		queue = queue[1:] // Dequeue
	}

	fmt.Println(len(checked) - 1)

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

	bagsByBag := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		mainSplit := strings.Split(line, "contain")
		if mainSplit[1] == " no other bags." {
			continue
		}
		container := strings.TrimSpace(mainSplit[0])
		container = container[:len(container)-5]

		for _, str := range strings.Split(mainSplit[1], ",") {
			split := strings.Split(str, " ")
			bag := strings.TrimSpace(split[2] + " " + split[3])
			num, _ := strconv.Atoi(split[1])
			for i := 0; i < num; i++ {
				bagsByBag[container] = append(bagsByBag[container], bag)
			}
		}
	}

	// look through
	queue := []string{"shiny gold"}
	num := -1

	for len(queue) > 0 {
		bag := queue[0]
		num++
		for _, outerBag := range bagsByBag[bag] {
			queue = append(queue, outerBag)
		}
		queue = queue[1:] // Dequeue
	}

	fmt.Println(num)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
