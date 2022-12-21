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
	for _, line := range shared.Splitlines("./sample.txt") {
		chunks := strings.Split(line, " ")
		for path := 0; path < len(chunks)-2; path++ {
			if chunks[path] == "->" {
				continue
			}
			= strings.Split(chunks[path], ","
			chunks[path+2]
		}
	}
	return
}
