package main

import (
	"fmt"

	"github.com/rustylampshade/advent-of-go/shared"
)

/*
For Day 8 I had added a shared.Grid class that can parse things that look like 2-D maps
and give some utility functions for walking around that map in four directions. I want to
reuse that here.

The shared.Grid class has a unique "constructor" called shared.NewGrid([]string) where I pass
in the lines that need to be parsed into the grid.

I first tried to just create `hill = shared.NewGrid(lines)` but then my `hill` variable is of
type shared.Grid and I can ONLY use the methods that already exist on the Grid class, defined in
the shared package. I attempted to define `func (g *shared.Grid) getStart()` and got an error that
"cannot define new methods on non-local type shared.Grid". Meaning I am not allowed to adjust the
methods of a type that is defined in a new package.

So the internet recommends that I extend the shared.Grid type into a type that IS defined in this
local package, so here I have `type Hill struct { shared.Grid }`. Neat, now I can make Hills as kind
of a special-case of Grids, and I can at least declare new methods that act on Hills (but wouldn't
generically work on Grids). Good good.

But the last problem -- How do I initialize `hill` now?! `hill := shared.NewGrid(lines)` still sets
`hill` to be a shared.Grid, and I can't figure out casts to force it into being a Hill. Internet snippets
show examples like

	myVar := ExtendedType{blah}

But that seems to expect that my base type uses the default initialization. Halp.
*/
type Hill struct {
	shared.Grid
}
type Coord struct {
	x, y int
}

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	hill := shared.NewGrid(shared.Splitlines("./input.txt"))

	hill.getStart()
	hill.getGoal()
	fmt.Println(hill)
	return
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
