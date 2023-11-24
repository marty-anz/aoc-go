package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	lines := strings.Split(input, "\n")

	total := 0

	for _, l := range lines {
		parts := strings.Split(l, " ")

		if len(parts) < 3 {
			continue
		}

		r := strings.Split(parts[0], "-")

		x, _ := strconv.Atoi(r[0])
		y, _ := strconv.Atoi(r[1])
		char := parts[1][0]
		pass := parts[2]

		valid := false
		count := 0

		for _, c := range pass {
			if byte(c) == char {
				count += 1
			}

			if count >= x && count <= y {
				valid = true
			}

			if count > y {
				valid = false
				break
			}
		}

		if valid {
			total += 1
		}
	}

	return total
}

func part2() int {
	lines := strings.Split(input, "\n")

	total := 0

	for _, l := range lines {
		parts := strings.Split(l, " ")

		if len(parts) < 3 {
			continue
		}

		r := strings.Split(parts[0], "-")

		x, _ := strconv.Atoi(r[0])
		y, _ := strconv.Atoi(r[1])
		char := parts[1][0]
		pass := parts[2]

		f, s := byte(pass[x-1]), byte(pass[y-1])

		if f != s && (f == char || s == char) {
			total += 1
		}
	}

	return total
}
