package main

import (
	"aoc/pkg/x"
	_ "embed"
	"fmt"
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
	lines := x.Lines(s)
	lines = append(lines, "")
	total := 0
	grp := map[rune]bool{}

	for _, l := range lines {
		l = x.Trim(l)
		if l == "" {
			total += len(grp)
			grp = map[rune]bool{}
			continue
		}

		for _, c := range l {
			grp[c] = true
		}
	}
	return total
}

func part2(s string) int {
	lines := x.Lines(s)
	lines = append(lines, "")
	total := 0
	grp, ppl := map[rune]int{}, 0

	for _, l := range lines {
		l = x.Trim(l)
		if l == "" {
			for _, c := range grp {
				if c == ppl {
					total += 1
				}
			}

			grp, ppl = map[rune]int{}, 0

			continue
		}
		ppl += 1

		for _, c := range l {
			grp[c] += 1
		}
	}

	return total
}
