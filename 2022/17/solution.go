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

func (r *Rock) displaceUp(maxTower, h int) {
    for j := maxTower; j <= maxTower + h + 5; j++ {
        // fmt.Printf("Creating %v\n", j)
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

func (r *Rock) displace(direction byte) {
    shifted := Rock{make([]Pt, len(r.cells))}
    copy(shifted.cells, r.cells)
    /*
    if direction == '<' {
        fmt.Println("Blowing left")
    } else {
        fmt.Println("Blowing right")
    }
    */
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

var tower map[int]map[int]string

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {


    tower = make(map[int]map[int]string)
    maxTower := 0
    tower[0] = make(map[int]string)
    tower[0][0], tower[0][8] = "+", "+"
    for i := 1; i < 8; i++ {
        tower[0][i] = "-"
    }

	jets := shared.Splitlines("./input.txt")[0]
    rockSequence := []Rock{HorizontalRock, PlusRock, AngleRock, VerticalRock, SquareRock}
    rockNumber, jetNumber := 0, 0

    // All these numbers change with sample.txt

    huge := 1_000_000_000_000
    startupRocks := 1711        // Fully dropped.
    cycleLength := 1725
    cycleHeight := 2694
    // prevNum, prevHeight := 0, 0
    for rockNumber = 0; rockNumber < startupRocks + (huge - startupRocks)%cycleLength; rockNumber++ {
        templateRock := rockSequence[rockNumber%len(rockSequence)]
        r := Rock{make([]Pt, len(templateRock.cells))}
        copy(r.cells, templateRock.cells)
        r.displaceUp(maxTower, 4)

        /*
        if rockNumber % len(rockSequence) == 0 && jetNumber % len(jets) == 1{
            fmt.Println(prevNum + rockNumber+1, maxTower-prevHeight)
            prevHeight = maxTower
        }
        */

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

        //printTower(rockNumber, maxTower)
    }
    fmt.Println(maxTower+((huge - startupRocks)/cycleLength)*cycleHeight)
    return
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
