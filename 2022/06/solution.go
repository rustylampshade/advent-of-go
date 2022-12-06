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
	return fmt.Sprint(findUniqueSpan(line, 4)), fmt.Sprint(findUniqueSpan(line, 14))
}

func findUniqueSpan(s string, span int) int {
	for i := span - 1; i < len(s); i++ {
		counts := shared.Counts(strings.Split(s[i-span+1:i+1], ""))
		if all(counts, 1) {
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
