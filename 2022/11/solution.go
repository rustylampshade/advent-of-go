package main

import (
	"fmt"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) bool
	onTrue    int
	onFalse   int

	inspections int
}

var divisor int

func ParseMonkey(inputLines []string) *Monkey {
	m := new(Monkey)
	for _, line := range inputLines {
		trimmed := strings.Trim(line, " ")
		tokens := strings.Split(trimmed, " ")
		if strings.HasPrefix(trimmed, "Starting items:") {
			for _, item := range tokens[2:] {
				m.items = append(m.items, shared.Atoi(strings.Trim(item, ",")))
			}
		} else if strings.HasPrefix(trimmed, "Operation:") {
			if strings.Join(tokens[3:], " ") == "old * old" {
				m.operation = func(n int) int { return n * n }
			} else {
				switch tokens[4] {
				case "*":
					m.operation = func(n int) int { return n * shared.Atoi(tokens[5]) }
				case "+":
					m.operation = func(n int) int { return n + shared.Atoi(tokens[5]) }
				}
			}
		} else if strings.HasPrefix(trimmed, "Test:") {
			m.test = func(n int) bool { return n%shared.Atoi(tokens[3]) == 0 }
			divisor *= shared.Atoi(tokens[3])
		} else if strings.HasPrefix(trimmed, "If true:") {
			m.onTrue = shared.Atoi(tokens[5])
		} else if strings.HasPrefix(trimmed, "If false:") {
			m.onFalse = shared.Atoi(tokens[5])
		}
	}
	return m
}

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	lines := shared.Splitlines("./input.txt")
	var monkeys []Monkey
	endOfMonkeys := shared.FindAll(lines, "")
	endOfMonkeys = append(endOfMonkeys, len(lines))
	divisor = 1
	for monkeyNumber := 0; monkeyNumber < len(endOfMonkeys); monkeyNumber++ {
		var startOfMonkey, endOfMonkey int
		if monkeyNumber == 0 {
			startOfMonkey = 1
		} else {
			startOfMonkey = endOfMonkeys[monkeyNumber-1] + 2
		}
		endOfMonkey = endOfMonkeys[monkeyNumber]
		monkeys = append(monkeys, *ParseMonkey(lines[startOfMonkey:endOfMonkey]))
	}

	for round := 1; round <= 10_000; round++ {
		for i := 0; i < len(monkeys); i++ {
			for _, item := range monkeys[i].items {
				monkeys[i].inspections++
				worry := monkeys[i].operation(item) % divisor
				if monkeys[i].test(worry) {
					monkeys[monkeys[i].onTrue].items = append(monkeys[monkeys[i].onTrue].items, worry)
				} else {
					monkeys[monkeys[i].onFalse].items = append(monkeys[monkeys[i].onFalse].items, worry)
				}
			}
			monkeys[i].items = nil
		}
	}
	for _, m := range monkeys {
		fmt.Println(m.inspections)
		// Should find the top two, but there are only a few monkeys.
	}
	return
}
