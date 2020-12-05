package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	one()
	two()
}

func one() {
	absPath, _ := filepath.Abs("src/day04/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numValid := 0
	m := make(map[string]bool)
	skip := false
	var sb strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if m["byr"] && m["iyr"] && m["eyr"] && m["hgt"] && m["hcl"] && m["ecl"] && m["pid"] {
				numValid++
			}
			m = make(map[string]bool)
		} else {
			for _, rune := range line {
				if rune == ' ' {
					skip = false
				} else if rune == ':' {
					m[sb.String()] = true
					sb.Reset()
					skip = true
				} else if !skip {
					sb.WriteRune(rune)
				}
			}
		}
		skip = false
	}

	if m["byr"] && m["iyr"] && m["eyr"] && m["hgt"] && m["hcl"] && m["ecl"] && m["pid"] {
		numValid++
	}

	fmt.Println(numValid)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day04/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numValid := 0
	m := make(map[string]bool)
	skip := false
	var sb strings.Builder
	var sbVal strings.Builder
	var field string
	valid := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if m["byr"] && m["iyr"] && m["eyr"] && m["hgt"] && m["hcl"] && m["ecl"] && m["pid"] && valid {
				numValid++
			}
			m = make(map[string]bool)
			valid = true
		} else if valid {
			for _, rune := range line {
				if rune == ' ' {
					skip = false
					valid = validate(field, sbVal.String())
					if !valid {
						break
					}
					sbVal.Reset()
				} else if rune == ':' {
					field = sb.String()
					m[field] = true
					sb.Reset()
					skip = true
				} else if !skip {
					sb.WriteRune(rune)
				} else {
					sbVal.WriteRune(rune)
				}
			}
			valid = validate(field, sbVal.String())
			sbVal.Reset()
		}
		skip = false
	}

	if m["byr"] && m["iyr"] && m["eyr"] && m["hgt"] && m["hcl"] && m["ecl"] && m["pid"] && valid {
		numValid++
	}

	fmt.Println(numValid)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var rHCL, _ = regexp.Compile("#([a-f0-9]{6})")
var rPID, _ = regexp.Compile("^\\d{9}$")

func validate(field string, value string) bool {
	var valid bool
	switch field {
	case "byr":
		num, _ := strconv.Atoi(value)
		valid = num >= 1920 && num <= 2002
	case "iyr":
		num, _ := strconv.Atoi(value)
		valid = num >= 2010 && num <= 2020
	case "eyr":
		num, _ := strconv.Atoi(value)
		valid = num >= 2020 && num <= 2030
	case "hgt":
		num, _ := strconv.Atoi(value[0 : len(value)-2])
		m := value[len(value)-2:]
		if m == "cm" {
			valid = num >= 150 && num <= 193
		} else if m == "in" {
			valid = num >= 59 && num <= 76
		} else {
			valid = false
		}
	case "hcl":
		valid = rHCL.MatchString(value)
	case "ecl":
		valid = value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth"
	case "pid":
		valid = rPID.MatchString(value)
	case "cid":
		valid = true
	}

	return valid
}
