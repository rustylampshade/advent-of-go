package main

import (
	"fmt"

	"github.com/rustylampshade/advent-of-go/shared"
)

func main() {
	part1, part2 := solve()
	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}
func solve() (part1 string, part2 string) {
	input := shared.ReadIntFromLine("./input.txt")

	return fmt.Sprint(p1(input)), fmt.Sprint(p2(input))
}

func p1(input []int) (result int) {
	var increased_depth int
	for i, depth := range input {
		if i == 0 {
			continue
		}
		if depth > input[i-1] {
			increased_depth += 1
		}
	}
	return increased_depth
}

func p2(input []int) (result int) {
	var increased_depth int
	for i := 3; i < len(input); i++ {
		if input[i] > input[i-3] {
			increased_depth += 1
		}
	}
	return increased_depth
}
