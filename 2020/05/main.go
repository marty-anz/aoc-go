package main

import (
	_ "embed"
	"fmt"

	"aoc/pkg/x"
)

//go:embed input
var input string

//go:embed test_input
var testInput string

func main() {
	// fmt.Println(part1(testInput))
	fmt.Println(part1(input))

	// fmt.Println(part2(testInput))
	fmt.Println(part2(input))
}

func findSeat(id string) (row int, col int) {
	s, e := 0, 127

	for _, c := range id[0:7] {
		half := s + (e-s+1)/2
		if c == 'F' {
			e = half - 1
		} else {
			s = half
		}
	}

	row = s
	s, e = 0, 7
	for _, c := range id[7:] {
		half := s + (e-s+1)/2
		if c == 'L' {
			e = half - 1
		} else {
			s = half
		}
	}

	col = s

	return
}

func part1(s string) int {
	seats := x.Lines(s)
	maxSeatID := 0
	for _, seat := range seats {
		r, c := findSeat(seat)
		seatID := r*8 + c
		if maxSeatID < seatID {
			maxSeatID = seatID
		}
	}

	return maxSeatID
}

func part2(s string) int {
	seats := x.Lines(s)
	found := map[int]bool{}

	for _, seat := range seats {
		r, c := findSeat(seat)
		seatID := r*8 + c
		found[seatID] = true
	}

	for i := 0; i < 127*8+7; i++ {
		if !found[i] && found[i+1] && found[i-1] {
			return i
		}
	}

	return 0
}
