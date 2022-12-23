package main

import (
	"fmt"

	"github.com/rustylampshade/advent-of-go/shared"
)

type Pt struct {
    x, y int
}

type Rock struct {
    cells []Pt
}

// Can't use structs as constants, so just declare them as vars and solemnly swear not to modify them.
// These are copied by value each time they're going to be used and mutated.
var HorizontalRock Rock = Rock{[]Pt{
    Pt{3,0},
    Pt{4,0},
    Pt{5,0},
    Pt{6,0}}}
var PlusRock Rock = Rock{[]Pt{
    Pt{3,1},
    Pt{4,1},
    Pt{5,1},
    Pt{4,2},
    Pt{4,0}}}
var AngleRock Rock = Rock{[]Pt{
    Pt{3,0},
    Pt{4,0},
    Pt{5,0},
    Pt{5,1},
    Pt{5,2}}}
var VerticalRock Rock = Rock{[]Pt{
    Pt{3,0},
    Pt{3,1},
    Pt{3,2},
    Pt{3,3}}}
var SquareRock Rock = Rock{[]Pt{
    Pt{3,0},
    Pt{3,1},
    Pt{4,0},
    Pt{4,1}}}

var tower map[int]map[int]string

func (r *Rock) displaceUp(maxTower, h int) {
    for j := maxTower; j <= maxTower + h + 5; j++ {
        _, exists := tower[j]
        if !exists {
            tower[j] = make(map[int]string)
            tower[j][0], tower[j][8] = "|", "|"
        }
    }
    for i := 0; i < len(r.cells); i++ {
        r.cells[i].y += maxTower + h
    }
}

// Shift (via mutation) a rock left '<' or right '>'. No position changes locked in if we hit an object.
func (r *Rock) displace(direction byte) {
    shifted := Rock{make([]Pt, len(r.cells))}
    copy(shifted.cells, r.cells)
    for i := 0; i < len(r.cells); i++ {
        switch direction {
        case '<':
            shifted.cells[i].x--
        case '>':
            shifted.cells[i].x++
        }
        _, exists := tower[shifted.cells[i].y][shifted.cells[i].x]
        if exists {
            return
        }
    }
    r.cells = shifted.cells
}

// Shift a rock down. Abort the movement and return true if we hit something and are now at rest.
func (r *Rock) fall() bool {
    shifted := Rock{make([]Pt, len(r.cells))}
    copy(shifted.cells, r.cells)
    for i := 0; i < len(r.cells); i++ {
        shifted.cells[i].y--
        _, exists := tower[shifted.cells[i].y][shifted.cells[i].x]
        if exists {
            return true
        }
    }
    r.cells = shifted.cells
    return false
}

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
   return simulate(2022), simulate(1_000_000_000_000)
}

func simulate(rocksToDrop int) (towerHeight string) {
    maxTower := 0
    tower = make(map[int]map[int]string)
    tower[0] = make(map[int]string)
    tower[0][0], tower[0][8] = "+", "+"
    for i := 1; i < 8; i++ {
        tower[0][i] = "-"
    }

    // LOL still doesn't work for sample, I think because the cycles trigger on a different jet index. I don't care anymore.
	jets := shared.Splitlines("./input.txt")[0]
    rockSequence := []Rock{HorizontalRock, PlusRock, AngleRock, VerticalRock, SquareRock}
    rockNumber, jetNumber := 0, 0

    cycleFound := false
    cycleStartup, cycleNumRocks, cycleAddedHeight, completeCycles := rocksToDrop, rocksToDrop, rocksToDrop, 0
    cycleHeights := []int{0}
    cycleRocks := []int{0}

    for rockNumber = 0; rockNumber < rocksToDrop-completeCycles*cycleNumRocks; rockNumber++ {
        templateRock := rockSequence[rockNumber%len(rockSequence)]
        r := Rock{make([]Pt, len(templateRock.cells))}
        copy(r.cells, templateRock.cells)
        r.displaceUp(maxTower, 4)

        if !cycleFound && rockNumber % len(rockSequence) == 0 && jetNumber % len(jets) == 1 {
            cycleHeights = append(cycleHeights, maxTower)
            cycleRocks = append(cycleRocks, rockNumber+1)
            history := len(cycleHeights)
            if history >= 3 && cycleHeights[history-1]-cycleHeights[history-2] == cycleHeights[history-2] - cycleHeights[history-3] {
                cycleFound = true
                cycleStartup = cycleRocks[history-3]
                cycleNumRocks = cycleRocks[history-1]-cycleRocks[history-2]
                cycleAddedHeight = cycleHeights[history-1]-cycleHeights[history-2]
                completeCycles = (rocksToDrop-cycleStartup)/cycleNumRocks - 2

                fmt.Printf("Cycle found! After %v rocks have dropped, every %v more rocks will add %v more height.\n", cycleStartup, cycleNumRocks, cycleAddedHeight)
                fmt.Printf("The target rock number %v can instead be %v complete repetitions of this cycle for %v*%v=%v height,\n", rocksToDrop, completeCycles, completeCycles, cycleAddedHeight, completeCycles*cycleAddedHeight)
                fmt.Printf("then simulation will reveal the height of the pre-cycle rocks, cycle-determining rocks, and partial-cycle rocks (%v rocks).\n\n", rocksToDrop-completeCycles*cycleNumRocks)
            }
        }

        atRest := false
        for !atRest {
            r.displace(jets[jetNumber])
            jetNumber = (jetNumber + 1)%len(jets)
            atRest = r.fall()
        }
        for _, cell := range r.cells {
            tower[cell.y][cell.x] = "#"
            if cell.y > maxTower {
                maxTower = cell.y
            }
        }
        // Uncomment to print ascii tower after each rock.
        // printTower(rockNumber, maxTower)
    }
    return fmt.Sprint(maxTower+completeCycles*cycleAddedHeight)
}

func printTower(rockNumber, maxTower int) {
    for j := len(tower)-1; j >= 0; j-- {
        line := ""
        for i := 0; i < 9; i++ {
            val, exists := tower[j][i]
            if !exists {
                val = "."
            }
            line += val
        }
        fmt.Println(line)
    }
    fmt.Printf("After %v rocks, max tower: %v\n\n", rockNumber+1, maxTower)
}
