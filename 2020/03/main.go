package main

import (
	_ "embed"
	"fmt"
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

func part1(input string) int {
	base := strings.Split(strings.TrimSpace(input), "\n")
	return track(base, 1, 3)
}

func track(base []string, r, c int) (trees int) {
	var col, row int

	cols := len(base[0])

	for row < len(base)-r {
		row += r
		col = (col + c) % cols
		if base[row][col] == '#' {
			trees += 1
		}
	}

	return trees
}

func part2(input string) int {
	base := strings.Split(strings.TrimSpace(input), "\n")

	return track(base, 1, 1) *
		track(base, 1, 3) *
		track(base, 1, 5) *
		track(base, 1, 7) *
		track(base, 2, 1)
}
