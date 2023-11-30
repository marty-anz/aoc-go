package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"

	"aoc/pkg/x"
)

//go:embed input.txt
var input string

//go:embed test_input.txt
var testInput string

//go:embed test_input2.txt
var testInput2 string

func main() {
	fmt.Println(part1(testInput))
	fmt.Println(part1(input))

	fmt.Println(part2(testInput2))
	fmt.Println(part2(input))
}

func containGold(color string, tree map[string]map[string]int) bool {
	if color == "shiny gold" {
		return false
	}

	if tree[color] == nil {
		return false
	}

	for k := range tree[color] {
		if k == "shiny gold" {
			return true
		}

		if containGold(k, tree) {
			return true
		}
	}

	return false
}

func part1(s string) int {
	lines := x.Lines(s)
	tree := map[string]map[string]int{}
	for _, line := range lines {
		line = x.Trim(line)
		parts := strings.Split(line, " contain ")

		holdColor := x.Trim(strings.Split(parts[0], " bags")[0])
		if holdColor == "" {
			continue
		}

		_, ok := tree[holdColor]
		if !ok {
			tree[holdColor] = map[string]int{}
		}

		bags := strings.Split(parts[1], ", ")
		r := regexp.MustCompile(`(\d+)([\w ]+)bags?`)
		for _, bag := range bags {
			matches := r.FindStringSubmatch(bag)

			if len(matches) == 3 {
				count, _ := strconv.Atoi(matches[1])
				name := x.Trim(matches[2])
				tree[holdColor][name] = count
			} else {
				tree[holdColor] = nil
			}
		}
	}

	count := 0
	for _, color := range maps.Keys(tree) {
		if containGold(color, tree) {
			count += 1
		}
	}

	return count
}

func countBags(color string, tree map[string]map[string]int) int {
	if tree[color] == nil {
		return 0
	}

	count := 0

	for c, n := range tree[color] {
		count += n

		count += n * countBags(c, tree)
	}

	return count
}

func part2(s string) int {
	lines := x.Lines(s)
	tree := map[string]map[string]int{}
	for _, line := range lines {
		line = x.Trim(line)
		parts := strings.Split(line, " contain ")

		holdColor := x.Trim(strings.Split(parts[0], " bags")[0])
		if holdColor == "" {
			continue
		}

		_, ok := tree[holdColor]
		if !ok {
			tree[holdColor] = map[string]int{}
		}

		bags := strings.Split(parts[1], ", ")
		r := regexp.MustCompile(`(\d+)([\w ]+)bags?`)
		for _, bag := range bags {
			matches := r.FindStringSubmatch(bag)

			if len(matches) == 3 {
				count, _ := strconv.Atoi(matches[1])
				name := x.Trim(matches[2])
				tree[holdColor][name] = count
			} else {
				tree[holdColor] = nil
			}
		}
	}

	return countBags("shiny gold", tree)
}
