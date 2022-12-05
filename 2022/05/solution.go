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

func initialize(startingConfig []string) map[int][]string {
	stacks := make(map[int][]string)
	for stackHeight := len(startingConfig) - 1; stackHeight >= 0; stackHeight-- {
		for offset, stackNum := 1, 1; offset < len(startingConfig[stackHeight]); offset, stackNum = offset+4, stackNum+1 {
			crate := string(startingConfig[stackHeight][offset])
			if crate != " " {
				stacks[stackNum] = append(stacks[stackNum], crate)
			}
		}
	}
	return stacks
}

func solve() (part1 string, part2 string) {
	lines := shared.Splitlines("./input.txt")
	divider := shared.FindFirst(lines, "")
	startingConfig := lines[:divider-1]
	moveInstructions := lines[divider+1:]

	stacks := initialize(startingConfig)
	for _, instruction := range moveInstructions {
		words := strings.Split(instruction, " ")
		count := shared.Atoi(words[1])
		source := shared.Atoi(words[3])
		target := shared.Atoi(words[5])

		for ; count > 0; count-- {
			var crates []string
			crates, stacks[source] = shared.Pop(stacks[source], 1)
			stacks[target] = append(stacks[target], crates...)
		}
	}
	for i := 1; i <= len(stacks); i++ {
		part1 += stacks[i][len(stacks[i])-1]
	}

	stacks = initialize(startingConfig)
	for _, instruction := range moveInstructions {
		words := strings.Split(instruction, " ")
		count := shared.Atoi(words[1])
		source := shared.Atoi(words[3])
		target := shared.Atoi(words[5])

		var crates []string
		crates, stacks[source] = shared.Pop(stacks[source], count)
		stacks[target] = append(stacks[target], crates...)
	}
	for i := 1; i <= len(stacks); i++ {
		part2 += stacks[i][len(stacks[i])-1]
	}
	return
}
