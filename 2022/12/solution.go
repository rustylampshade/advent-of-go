package main

import (
	"fmt"

	"github.com/rustylampshade/advent-of-go/shared"
)

type Hill struct {
	*shared.Grid
}
type Coord struct {
	x, y int
}

var bestPath []Coord
var bestLetter string

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	hill := Hill{shared.NewGrid(shared.Splitlines("./input.txt"))}

	start := hill.getStart()
	goal := hill.getGoal()
	bestPath = make([]Coord, 0)
	bestLetter = "a"

	hill.findAllPaths(start, goal, []Coord{})
	fmt.Println(len(bestPath))
	return
}

func (hill *Hill) findAllPaths(start Coord, goal Coord, path []Coord) {
	if len(path) > len(bestPath) && len(bestPath) > 0 {
		return
	}
	if start == goal {
		if len(path) < len(bestPath) || len(bestPath) == 0 {
			fmt.Printf("Found a new best path of length %v\n", len(path))
			bestPath = path
		}
		return
	}
	for _, move := range hill.getValidMoves(start) {
		if TestIn(path, move) {
			continue
		}
		hill.findAllPaths(move, goal, append(path, start))
	}
}

func TestIn(path []Coord, c Coord) bool {
	for _, seen := range path {
		if c == seen {
			return true
		}
	}
	return false
}

func (hill *Hill) getStart() Coord {
	for i := 0; i < hill.LenX; i++ {
		for j := 0; j < hill.LenY; j++ {
			if hill.ValAt(i, j) == "S" {
				hill.SetAt(i, j, "a")
				return Coord{i, j}
			}
		}
	}
	panic("Couldn't find a starting 'S'!")
}

func (hill *Hill) getGoal() Coord {
	for i := 0; i < hill.LenX; i++ {
		for j := 0; j < hill.LenY; j++ {
			if hill.ValAt(i, j) == "E" {
				hill.SetAt(i, j, "z")
				return Coord{i, j}
			}
		}
	}
	panic("Couldn't find a goal 'E'!")
}

func (hill *Hill) getValidMoves(target Coord) (moves []Coord) {
	hill.Move(target.x, target.y)
	current := hill.Val()
	if current > bestLetter {
		bestLetter = current
		fmt.Printf("New high! %v\n", bestLetter)
	}
	if height, x, y, err := hill.ValRight(); err == nil {
		if validStep(current, height) {
			moves = append(moves, Coord{x, y})
		}
	}
	if height, x, y, err := hill.ValAbove(); err == nil {
		if validStep(current, height) {
			moves = append(moves, Coord{x, y})
		}
	}
	if height, x, y, err := hill.ValBelow(); err == nil {
		if validStep(current, height) {
			moves = append(moves, Coord{x, y})
		}
	}
	if height, x, y, err := hill.ValLeft(); err == nil {
		if validStep(current, height) {
			moves = append(moves, Coord{x, y})
		}
	}
	return
}

func validStep(src string, tgt string) bool {
	return int(rune(tgt[0]))-int(rune(src[0])) <= 1
}
