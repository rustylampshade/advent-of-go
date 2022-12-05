package main

import (
	"fmt"
	"sort"

	"github.com/rustylampshade/advent-of-go/shared"
)

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	lines := shared.Splitlines("./input.txt")

	var start = 0
	var calories_held []int
	for _, end := range append(shared.FindAll(lines, ""), len(lines)) {
		total_elf_calories := 0
		for _, snack := range lines[start:end] {
			calories := shared.Atoi(snack)
			total_elf_calories += calories
		}
		calories_held = append(calories_held, total_elf_calories)
		start = end + 1
	}
	sort.Ints(calories_held)
	top1_calories := shared.Sum(calories_held[len(calories_held)-1:])
	top3_calories := shared.Sum(calories_held[len(calories_held)-3:])

	return fmt.Sprint(top1_calories), fmt.Sprint(top3_calories)
}
