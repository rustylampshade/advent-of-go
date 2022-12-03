package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/rustylampshade/advent-of-go/shared"
)

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	lines := shared.Splitlines("./input.txt")

	part1_priority := 0
	for _, rucksack := range lines {
		items := strings.Split(rucksack, "")
		compartmentLength := len(items) / 2
		compartment1 := items[:compartmentLength]
		compartment2 := items[compartmentLength:]
		common := findCommonElementTwo(compartment1, compartment2)
		part1_priority += priority(common)
	}

	part2_priority := 0
	for i := 0; i < len(lines); i += 3 {
		elf1 := strings.Split(lines[i], "")
		elf2 := strings.Split(lines[i+1], "")
		elf3 := strings.Split(lines[i+2], "")
		common := findCommonElementThree(elf1, elf2, elf3)
		part2_priority += priority(common)
	}

	return fmt.Sprint(part1_priority), fmt.Sprint(part2_priority)
}

func sortString(unsorted []string) []string {
	sort.Slice(unsorted, func(i int, j int) bool {
		return unsorted[i] < unsorted[j]
	})
	return unsorted
}

func findCommonElementTwo(c1 []string, c2 []string) string {
	c1 = sortString(c1)
	c2 = sortString(c2)
	for i, j := 0, 0; i < len(c1) && j < len(c2); {
		a := c1[i]
		b := c2[j]
		if a == b {
			return a
		}
		if a < b {
			i++
		}
		if b < a {
			j++
		}
	}
	return ""
}

func findCommonElementThree(c1 []string, c2 []string, c3 []string) string {
	c1 = sortString(c1)
	c2 = sortString(c2)
	c3 = sortString(c3)
	for i, j, k := 0, 0, 0; i < len(c1) && j < len(c2) && k < len(c3); {
		a := c1[i]
		b := c2[j]
		c := c3[k]
		if a == b && b == c {
			return a
		}
		if a <= b && a <= c {
			i++
		}
		if b <= a && b <= c {
			j++
		}
		if c <= a && c <= b {
			k++
		}
	}
	return ""
}

func priority(letter string) int {
	r := rune(letter[0])
	if unicode.IsUpper(r) {
		return int(r) - 38
	} else {
		return int(r) - 96
	}
}
