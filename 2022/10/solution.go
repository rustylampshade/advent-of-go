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

type register int
type cycle int
type change struct {
	t cycle
	x register
}

var x register
var t cycle
var history []change

func solve() (part1 string, part2 string) {
	x, t = 1, 1
	history = append(history, change{t, x})

	signalStrength := 0
	for _, line := range shared.Splitlines("./input.txt") {
		if line == "noop" {
			t++
		} else {
			if t == 19 || (t-20)%40 == 39 {
				fmt.Printf("Adding signal at t == %v\n", t+1)
				signalStrength += int(t+1) * int(x)
			}
			t += 2
			x += register(shared.Atoi(strings.Split(line, " ")[1]))
		}
		if t == 20 || (t-20)%40 == 0 {
			fmt.Printf("Adding signal at t == %v\n", t)
			signalStrength += int(t) * int(x)
		}
		history = append(history, change{t, x})
	}
	var pixelTime cycle = 1
	for j := 1; j <= 6; j++ {
		for i := 0; i <= 39; i++ {
			var pixelRegister register
			for _, c := range history {
				if c.t <= pixelTime {
					pixelRegister = c.x
				} else {
					break
				}
			}
			if abs(int(pixelRegister)-i) <= 1 {
				fmt.Print("##")
			} else {
				fmt.Print("  ")
			}
			pixelTime++
		}
		fmt.Print("\n")
	}
	fmt.Println(signalStrength)
	return
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
