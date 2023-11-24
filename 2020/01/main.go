package main

import (
	"aoc/pkg/s"
	_ "embed"
	"fmt"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	numbers := s.Atois(input)

	lookup := s.Dict(numbers)

	for _, n := range numbers {
		m := 2020 - n
		if lookup[m] {
			return m * n
		}
	}

	return 0
}

func part2() int {
	numbers := s.Atois(input)

	lookup := s.Dict(numbers)

	for i, n := range numbers {
		for j, m := range numbers {
			if i != j {
				left := 2020 - m - n
				if lookup[left] {
					return m * n * left
				}
			}
		}
	}

	return 0
}
