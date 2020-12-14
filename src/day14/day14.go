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
	absPath, _ := filepath.Abs("src/day14/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[int]int64)

	var baseMask int64
	var keepMask int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			baseMask = 0
			keepMask = 0

			maskString := line[7:]
			bit := int64(1)
			for c := len(maskString) - 1; c >= 0; c-- {
				if maskString[c] == 'X' {
					keepMask += bit
				} else if maskString[c] == '1' {
					baseMask += bit
				}
				bit <<= 1
			}
		} else {
			equalSplit := strings.Split(line, " = ")
			value, _ := strconv.ParseInt(equalSplit[1], 10, 64)
			key, _ := strconv.Atoi(equalSplit[0][4:(len(equalSplit[0]) - 1)])

			m[key] = baseMask + (value & keepMask)
		}
	}

	sum := int64(0)
	for _, value := range m {
		sum += value
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day14/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[int64]int64)

	var floatBitmasks []int64
	var orMask int64
	var clearMask int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			floatBitmasks = []int64{0}
			orMask = 0
			clearMask = 0

			maskString := line[7:]
			bit := int64(1)
			for c := len(maskString) - 1; c >= 0; c-- {
				if maskString[c] == 'X' {
					var newMasks []int64
					for _, mask := range floatBitmasks {
						newMasks = append(newMasks, mask)
						newMasks = append(newMasks, mask+bit)
					}
					floatBitmasks = newMasks
					clearMask += bit
				} else if maskString[c] == '1' {
					orMask += bit
				}
				bit <<= 1
			}
			clearMask = ^clearMask
		} else {
			equalSplit := strings.Split(line, " = ")
			value, _ := strconv.ParseInt(equalSplit[1], 10, 64)
			key, _ := strconv.ParseInt(equalSplit[0][4:(len(equalSplit[0])-1)], 10, 64)

			keyWithoutFloatingBits := (key & clearMask) | orMask
			for _, floatBitmask := range floatBitmasks {
				m[keyWithoutFloatingBits|floatBitmask] = value
			}
		}
	}

	sum := int64(0)
	for _, value := range m {
		sum += value
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
