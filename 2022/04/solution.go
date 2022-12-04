package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	lines := shared.Splitlines("./input.txt")

	numCompletelyOverlapping, numPartiallyOverlapping := 0, 0
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		a1_min, a1_max := getRangeMinMax(assignments[0])
		a2_min, a2_max := getRangeMinMax(assignments[1])

		if (a1_min <= a2_min && a1_max >= a2_max) || (a2_min <= a1_min && a2_max >= a1_max) {
			numCompletelyOverlapping++
		}
		// Testing for completely disjoint sets is easy, so negate that test to find any partial overlaps.
		if !((a1_max < a2_min) || (a2_max < a1_min)) {
			numPartiallyOverlapping++
		}
	}
	return fmt.Sprint(numCompletelyOverlapping), fmt.Sprint(numPartiallyOverlapping)
}

func getRangeMinMax(s string) (int, int) {
	stringNums := strings.Split(s, "-")
	return toInt(stringNums[0]), toInt(stringNums[1])
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Not an int: " + s)
	}
	return n
}
