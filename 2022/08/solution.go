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
	forest := shared.NewGrid(shared.Splitlines("./input.txt"))

	totalVisible := 0
	for i := 0; i < forest.LenX; i++ {
		for j := 0; j < forest.LenY; j++ {
			forest.Move(i, j)
			visible := false
			for _, trees := range [][]string{forest.SliceAbove(), forest.SliceBelow(), forest.SliceLeft(), forest.SliceRight()} {
				if _, max := shared.Max(trees); len(trees) == 0 || forest.Val() > max {
					visible = true
				}
			}
			if visible {
				totalVisible++
			}
		}
	}

	bestScenicScore, scenicScore := 0, 0
	for i := 0; i < forest.LenX; i++ {
		for j := 0; j < forest.LenY; j++ {
			forest.Move(i, j)
			var visibleInDirections [4]int
			for dir, trees := range [][]string{forest.SliceBelow(), forest.SliceRight(), shared.Reverse(forest.SliceAbove()), shared.Reverse(forest.SliceLeft())} {
				for _, tree := range trees {
					visibleInDirections[dir]++
					if tree >= forest.Val() {
						break
					}
				}
			}
			scenicScore = visibleInDirections[0] * visibleInDirections[1] * visibleInDirections[2] * visibleInDirections[3]
			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}

	return fmt.Sprint(totalVisible), fmt.Sprint(bestScenicScore)
}
