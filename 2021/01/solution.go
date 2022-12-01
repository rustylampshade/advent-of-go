package main

import (
	"fmt"

	"github.com/rustylampshade/advent-of-go/shared"
)

func main() {
	input := shared.ReadIntFromLine("./input.txt")

	fmt.Println(p1(input))
	fmt.Println(p2(input))
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
