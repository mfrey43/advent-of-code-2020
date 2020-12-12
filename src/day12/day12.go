package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	one()
	two()
}

func one() {
	absPath, _ := filepath.Abs("src/day12/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := 0
	y := 0
	dir := 0 // 0 => E
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cmd := line[0]
		num, _ := strconv.Atoi(line[1:])

		if cmd == 'F' {
			switch dir {
			case 0:
				cmd = 'E'
			case 1:
				cmd = 'S'
			case 2:
				cmd = 'W'
			case 3:
				cmd = 'N'
			}
		}

		switch cmd {
		case 'N':
			y -= num
		case 'E':
			x += num
		case 'S':
			y += num
		case 'W':
			x -= num
		case 'R':
			dir = (dir + num/90) % 4
		case 'L':
			dir = (dir - num/90) % 4
			if dir < 0 {
				dir += 4
			}
		}
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day12/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	xS := 0
	yS := 0
	xW := 10
	yW := -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cmd := line[0]
		num, _ := strconv.Atoi(line[1:])

		switch cmd {
		case 'N':
			yW -= num
		case 'E':
			xW += num
		case 'S':
			yW += num
		case 'W':
			xW -= num
		case 'F':
			xS += num * xW
			yS += num * yW
		case 'R':
			for i := 0; i < num/90; i++ {
				tmp := xW
				xW = -yW
				yW = tmp
			}
		case 'L':
			for i := 0; i < num/90; i++ {
				tmp := xW
				xW = yW
				yW = -tmp
			}
		}
	}

	fmt.Println(math.Abs(float64(xS)) + math.Abs(float64(yS)))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
