package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

type Pt struct {
	x, y, z int
}

type AirPocket struct {
	air  map[string]bool
	size int
}

var cubes map[string]bool
var knownAirPockets []AirPocket

func ParsePt(s string) (p Pt) {
	tokens := strings.Split(s, ",")
	return Pt{x: shared.Atoi(tokens[0]), y: shared.Atoi(tokens[1]), z: shared.Atoi(tokens[2])}
}
func (p *Pt) Str() string {
	return fmt.Sprintf("%v,%v,%v", p.x, p.y, p.z)
}
func (p *Pt) Up() Pt {
	return Pt{x: p.x, y: p.y + 1, z: p.z}
}
func (p *Pt) Down() Pt {
	return Pt{x: p.x, y: p.y - 1, z: p.z}
}
func (p *Pt) Left() Pt {
	return Pt{x: p.x - 1, y: p.y, z: p.z}
}
func (p *Pt) Right() Pt {
	return Pt{x: p.x + 1, y: p.y, z: p.z}
}
func (p *Pt) Front() Pt {
	return Pt{x: p.x, y: p.y, z: p.z + 1}
}
func (p *Pt) Back() Pt {
	return Pt{x: p.x, y: p.y, z: p.z - 1}
}

func (cell *Pt) ExpandAirPocket(pocket AirPocket) (err error) {
	//fmt.Printf("Expanding the pocket, cell %v and size %v\n", cell, pocket.size)
	if pocket.size >= 5_000 {
		// It's a bit lame to just pick a hardcoded value here, but there is also a limit to the size of air pockets.
		// Air pockets' volume is maximized for a given surface area when they're roughly cubelike (hahaha sure, spheres
		// are higher volume:surface but is that even true in integer coords? I think everything looks like a cube). A
		// cubic air pocket would have SA = 6*(side)^2. We already computed the absolute maximum surface area (interior +
		// exterior), so probably only half of that can be interior surface area. So we can probably do sqrt(part1/12) to
		// be about the upper bound on the side length? And then that cubed would be the max pocket size.
		// That gives me ~4,600 as the truly maximum bubble size.
		err = errors.New("advent can't be this mean right?")
		return
	}
	for _, adjacentCell := range []Pt{cell.Up(), cell.Down(), cell.Left(), cell.Right(), cell.Front(), cell.Back()} {
		//fmt.Printf("Expanding the pocket, cell %v to %v and size %v\n", cell, adjacentCell, pocket.size)
		s := adjacentCell.Str()
		_, solid := cubes[s]
		_, alreadyFound := pocket.air[s]
		if solid || alreadyFound {
			continue
		}
		pocket.air[s] = true
		pocket.size++
		err = adjacentCell.ExpandAirPocket(pocket)
		if err != nil {
			return
		}
	}
	return
}

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {

	cubes = make(map[string]bool)
	for _, line := range shared.Splitlines("./input.txt") {
		cubes[line] = true
	}

	surface := 0
	var possibleExterior []Pt = make([]Pt, 0)
	for k := range cubes {
		p := ParsePt(k)
		for _, adj := range []Pt{p.Up(), p.Down(), p.Left(), p.Right(), p.Front(), p.Back()} {
			_, exists := cubes[adj.Str()]
			if !exists {
				surface++
				if !TestIn(possibleExterior, adj) {
					possibleExterior = append(possibleExterior, adj)
				}
			}
		}
	}

	knownAirPockets = make([]AirPocket, 0)
	for _, cell := range possibleExterior {
		if CellInKnownPockets(cell) {
			continue
		}
		fmt.Printf("Starting a new cell exploration at %v\n", cell)
		pocket := AirPocket{air: map[string]bool{cell.Str(): true}, size: 1}
		err := cell.ExpandAirPocket(pocket)
		if err == nil {
			knownAirPockets = append(knownAirPockets, pocket)
		}
	}
	fmt.Println(len(knownAirPockets))
	allInteriorAir := make(map[string]bool)
	for _, pocket := range knownAirPockets {
		fmt.Println(pocket)
		for key := range pocket.air {
			allInteriorAir[key] = true
		}
	}
	exteriorSurface := 0
	for k := range cubes {
		p := ParsePt(k)
		for _, adj := range []Pt{p.Up(), p.Down(), p.Left(), p.Right(), p.Front(), p.Back()} {
			_, solid := cubes[adj.Str()]
			_, interior := allInteriorAir[adj.Str()]
			if !solid && !interior {
				exteriorSurface++
			}
		}
	}

	fmt.Println(surface)
	fmt.Println(exteriorSurface)

	return
}

// Check if this air cell is already categorized into an air pocket.
func CellInKnownPockets(cell Pt) bool {
	for _, pocket := range knownAirPockets {
		_, exists := pocket.air[cell.Str()]
		if exists {
			return true
		}
	}
	return false
}

func TestIn(array []Pt, elem Pt) bool {
	for _, v := range array {
		if v == elem {
			return true
		}
	}
	return false
}
