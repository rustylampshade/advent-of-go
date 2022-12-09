package main

import (
	"fmt"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

type Pt struct {
	x int
	y int
}

var knots []Pt

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	movements := shared.Splitlines("./input.txt")

	p1cnt := len(ropePhysics(1, movements))
	p2cnt := len(ropePhysics(9, movements))
	return fmt.Sprint(p1cnt), fmt.Sprint(p2cnt)
}

func ropePhysics(numberOfTails int, movements []string) (visited map[string]bool) {
	knots = make([]Pt, 0)
	knots = append(knots, Pt{0, 0}) // Head
	for i := 1; i <= numberOfTails; i++ {
		knots = append(knots, Pt{0, 0})
	}
	visited = make(map[string]bool)
	for _, movement := range movements {
		direction := strings.Split(movement, " ")[0]
		count := strings.Split(movement, " ")[1]
		for i := 0; i < shared.Atoi(count); i++ {
			moveHead(direction)
			for knot := range knots {
				moveTail(knot)
			}
			visited[fmt.Sprintf("%v,%v", knots[len(knots)-1].x, knots[len(knots)-1].y)] = true
		}
	}
	return visited
}

func moveHead(direction string) {
	switch direction {
	case "R":
		knots[0].x++
	case "L":
		knots[0].x--
	case "U":
		knots[0].y++
	case "D":
		knots[0].y--
	}
}

func moveTail(tailNum int) {
	if tailNum == 0 {
		return
	}
	xDelta := knots[tailNum-1].x - knots[tailNum].x
	yDelta := knots[tailNum-1].y - knots[tailNum].y

	// "If the head and tail are touching..."
	if abs(xDelta) <= 1 && abs(yDelta) <= 1 {
		return
	}
	// "If the head is ever two steps directly up, down, left, or right from the tail, the tail
	// must also move one step in that direction."
	if abs(xDelta) == 2 && yDelta == 0 {
		knots[tailNum].x += xDelta / 2
	} else if abs(yDelta) == 2 && xDelta == 0 {
		knots[tailNum].y += yDelta / 2
	} else {
		// "Otherwise, the tail always moves one step diagonally"
		knots[tailNum].x += sgn(xDelta)
		knots[tailNum].y += sgn(yDelta)
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sgn(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}
