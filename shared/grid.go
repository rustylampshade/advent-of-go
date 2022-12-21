package shared

import (
	"errors"
	"fmt"
	"strings"
)

type Grid struct {
	x, y       int
	LenX, LenY int
	content    [][]string
}

func (g *Grid) Move(i int, j int) {
	g.x, g.y = i, j
}

func NewGrid(inputLines []string) *Grid {
	p := new(Grid)
	for j, line := range inputLines {
		p.content = append(p.content, make([]string, 0))
		p.content[j] = strings.Split(line, "")
	}
	p.LenX = len(p.content[0])
	p.LenY = len(p.content)
	return p
}

func (g *Grid) Print() {
	for _, line := range g.content {
		fmt.Println(strings.Join(line, ""))
	}
}

func (g *Grid) SetAt(x int, y int, v string) {
	g.content[y][x] = v
}

func (g *Grid) Set(v string) {
	g.content[g.y][g.x] = v
}

func (g *Grid) ValAt(x int, y int) string {
	return g.content[y][x]
}

func (g *Grid) Val() string {
	return g.content[g.y][g.x]
}

func (g *Grid) SliceLeft() (ret []string) {
	if g.x == 0 {
		return []string{}
	} else {
		return g.content[g.y][:g.x]
	}
}

func (g *Grid) SliceRight() (ret []string) {
	return g.content[g.y][g.x+1:]
}

func (g *Grid) SliceAbove() (ret []string) {
	for j := 0; j < g.y; j++ {
		ret = append(ret, g.content[j][g.x])
	}
	return ret
}

func (g *Grid) SliceBelow() (ret []string) {
	for j := g.y + 1; j < g.LenY; j++ {
		ret = append(ret, g.content[j][g.x])
	}
	return ret
}

func (g *Grid) ValLeft() (val string, x int, y int, err error) {
	x, y = g.x-1, g.y
	if g.x == 0 {
		return val, x, y, errors.New("no value in this direction")
	} else {
		return g.content[y][x], x, y, nil
	}
}

func (g *Grid) ValRight() (val string, x int, y int, err error) {
	x, y = g.x+1, g.y
	if g.x == g.LenX-1 {
		return val, x, y, errors.New("no value in this direction")
	} else {
		return g.content[y][x], x, y, nil
	}
}

func (g *Grid) ValAbove() (val string, x int, y int, err error) {
	x, y = g.x, g.y-1
	if g.y == 0 {
		return val, x, y, errors.New("no value in this direction")
	} else {
		return g.content[y][x], x, y, nil
	}
}

func (g *Grid) ValBelow() (val string, x int, y int, err error) {
	x, y = g.x, g.y+1
	if g.y == g.LenY-1 {
		return val, x, y, errors.New("no value in this direction")
	} else {
		return g.content[y][x], x, y, nil
	}
}
