package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var input string

//go:embed test_input
var testInput string

func main() {
	fmt.Println(part1(testInput))
	fmt.Println(part1(input))

	fmt.Println(part2(testInput))
	fmt.Println(part2(input))
}

func part1(s string) int {
	passports := strings.Split(strings.TrimSpace(s), "\n")

	passports = append(passports, "\n")

	passport := map[string]string{}
	validCount := 0

	for _, line := range passports {
		line = strings.TrimSpace(line)

		if len(line) > 0 {
			fields := strings.Split(line, " ")

			for _, field := range fields {
				pair := strings.Split(field, ":")
				passport[pair[0]] = pair[1]
			}
			continue
		}

		if validPassport1(passport) {
			validCount += 1
		}

		passport = map[string]string{}
	}

	return validCount
}

func validPassport1(passport map[string]string) bool {
	attrs := []string{`byr`, `iyr`, `eyr`, `hgt`, `hcl`, `ecl`, `pid`}
	for _, attr := range attrs {
		_, ok := passport[attr]
		if !ok {
			return false
		}
	}

	return true
}

func part2(s string) int {
	passports := strings.Split(strings.TrimSpace(s), "\n")

	passports = append(passports, "\n")

	passport := map[string]string{}
	validCount := 0

	for _, line := range passports {
		line = strings.TrimSpace(line)

		if len(line) > 0 {
			fields := strings.Split(line, " ")

			for _, field := range fields {
				pair := strings.Split(field, ":")

				passport[pair[0]] = pair[1]
			}

			continue
		}

		if validPassport2(passport) {
			validCount += 1
		}

		passport = map[string]string{}
	}

	return validCount
}

func validPassport2(passport map[string]string) bool {
	if n, _ := strconv.Atoi(passport[`byr`]); n < 1920 || n > 2002 {
		return false
	}

	if n, _ := strconv.Atoi(passport[`iyr`]); n < 2010 || n > 2020 {
		return false
	}

	if n, _ := strconv.Atoi(passport[`eyr`]); n < 2020 || n > 2030 {
		return false
	}

	if hgt, _ := passport[`hgt`]; len(hgt) < 3 {
		return false
	} else {
		measure := hgt[len(hgt)-2:]
		if !slices.Contains([]string{"cm", "in"}, measure) {
			return false
		}

		n, _ := strconv.Atoi(hgt[0 : len(hgt)-2])
		if measure == "cm" && (n < 150 || n > 193) {
			return false
		}
		if measure == "in" && (n < 59 || n > 76) {
			return false
		}
	}

	if matched, _ := regexp.Match(`^#[0-9a-f]{6}$`, []byte(passport[`hcl`])); !matched {
		return false
	}

	eyeColours := []string{`amb`, `blu`, `brn`, `gry`, `grn`, `hzl`, `oth`}
	if !slices.Contains(eyeColours, passport[`ecl`]) {
		return false
	}

	if matched, _ := regexp.Match("^[0-9]{9}$", []byte(passport[`pid`])); !matched {
		return false
	}

	return true
}
