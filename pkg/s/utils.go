package s

import (
	"strconv"
	"strings"
)

// text to int slice
func Atois(input string) []int {
	lines := strings.Split(input, "\n")

	ns := []int{}

	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		ns = append(ns, n)
	}

	return ns
}

// look up dictionary
func Dict[K comparable](list []K) map[K]bool {
	m := map[K]bool{}

	for _, l := range list {
		m[l] = true
	}

	return m
}
