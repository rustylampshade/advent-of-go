package main

import (
	"fmt"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	line := shared.Splitlines("./input.txt")[0]
	letters := strings.Split(line, "")
	return fmt.Sprint(findUniqueSpan(letters, 4)), fmt.Sprint(findUniqueSpan(letters, 14))
}

func findUniqueSpan(s []string, span int) int {
	// Given a span of 4, we need to start with s[0:4] to have a 4-wide window.
	for i := span - 1; i < len(s); i++ {
		if shared.TestEntirelyUnique(s[i-span+1 : i+1]) {
			// Add one to convert from 0-index to AoC's counting.
			return i + 1
		}
	}
	return -1
}

func all(m map[string]int, val int) bool {
	for _, v := range m {
		if v != val {
			return false
		}
	}
	return true
}
