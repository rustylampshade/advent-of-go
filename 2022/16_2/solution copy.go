package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/rustylampshade/advent-of-go/shared"
)

type Valve struct {
	label     string
	flowRate  int
	neighbors []string
}

type Path struct {
	currentLocations []string // Length of this slice indicates how many creatures are opening valves.
	openValveNames   []string
	closedValveRates []int
	timeUsed         int
	pressure         int
}

var graph map[string]Valve
var bestPath Path
var timeLimit int

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	lines := shared.Splitlines("./sample.txt")
	graph = make(map[string]Valve)
	rates := make([]int, 0)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		label := tokens[1]
		rate := shared.Atoi(strings.Trim(strings.Split(tokens[4], "=")[1], ";"))
		if rate != 0 {
			rates = append(rates, rate)
		}
		neighbors := strings.Split(strings.Join(tokens[9:], ""), ",")
		graph[label] = Valve{label: label, flowRate: rate, neighbors: neighbors}
	}
	sort.Slice(rates, func(i, j int) bool { return rates[i] > rates[j] })
	for _, v := range graph {
		sort.Slice(v.neighbors, func(i, j int) bool { return graph[v.neighbors[i]].flowRate > graph[v.neighbors[j]].flowRate })
	}

	// Part 1
	start := time.Now()
	timeLimit = 30
	startingPath := Path{currentLocations: []string{"AA"}, timeUsed: 1, closedValveRates: rates[:]}
	bestPath = Path{}
	dfs(startingPath)
	part1 = fmt.Sprint(bestPath.pressure)
	duration := time.Since(start)
	fmt.Println(duration)

	/*
		// Part 2
		start = time.Now()
		timeLimit = 26
		startingPath = Path{currentLocations: []string{"AA", "AA"}, timeUsed: 1, closedValveRates: rates[:]}
		bestPath = Path{}
		dfs(startingPath)
		part2 = fmt.Sprint(bestPath.pressure)
		duration = time.Since(start)
		fmt.Println(duration)
	*/

	return
}

func dfs(path Path) {
	if path.timeUsed >= timeLimit {
		if path.pressure > bestPath.pressure {
			fmt.Printf("New best path found with %v\n", path.pressure)
			bestPath = path
		}
		return
	}

	if len(path.closedValveRates) == 0 {
		//fmt.Println("Optimized from closed valve number")
		return
	}
	maxAdditionalPressure := path.MaximumImprovement()
	if path.pressure+maxAdditionalPressure <= bestPath.pressure {
		//fmt.Println("Optimized from max possible pressure")
		return
	}

	myValve := graph[path.currentLocations[0]]
	for i := -1; i < len(myValve.neighbors); i++ {
		var justMe Path = path.copy()
		var err error
		if i == -1 {
			err = justMe.open(myValve)
		} else {
			err = justMe.visit(myValve, graph[myValve.neighbors[i]])
		}
		if err != nil {
			continue
		}

		if len(path.currentLocations) == 1 {
			justMe.timeUsed++
			dfs(justMe)
		} else {
			elephantValve := graph[path.currentLocations[1]]
			for j := len(elephantValve.neighbors); j >= 0; j-- {
				var bothPath Path = justMe.copy()
				if j == len(elephantValve.neighbors) {
					err = bothPath.open(elephantValve)
				} else {
					err = bothPath.visit(elephantValve, graph[elephantValve.neighbors[j]])
				}
				if err != nil {
					continue
				}

				bothPath.timeUsed++
				dfs(bothPath)
			}
		}
	}

	return
}

func (p *Path) copy() (copiedPath Path) {
	copiedPath.currentLocations = append(copiedPath.currentLocations, p.currentLocations...)
	copiedPath.openValveNames = append(copiedPath.openValveNames, p.openValveNames...)
	copiedPath.closedValveRates = append(copiedPath.closedValveRates, p.closedValveRates...)
	copiedPath.timeUsed = p.timeUsed
	copiedPath.pressure = p.pressure
	return
}

func (p *Path) MaximumImprovement() (maxAdditionalPressure int) {
	for round, valvesOpened := 0, 0; p.timeUsed+round*2 <= timeLimit; round++ {
		for worker := 0; worker < len(p.currentLocations); worker++ {
			if valvesOpened >= len(p.closedValveRates) {
				return
			}
			maxAdditionalPressure += p.closedValveRates[valvesOpened] * (timeLimit - p.timeUsed - 2*round)
			valvesOpened++
		}
	}
	return
}

func (p *Path) visit(from, to Valve) (err error) {
	idx, err := shared.FindFirst(p.currentLocations, from.label)
	if err != nil {
		panic("wtf visit")
	}
	p.currentLocations = append(shared.RemoveIndex(p.currentLocations, idx), to.label)
	return err
}

func (p *Path) open(v Valve) (err error) {
	// The valve will be opened if it was not already opened by a previous visit to this valve and will release steam
	if shared.TestIn(p.openValveNames, v.label) || v.flowRate == 0 {
		err = errors.New("invalid valve opening")
		return
	}
	p.openValveNames = append(p.openValveNames, v.label)
	idx, err := shared.FindFirst(p.closedValveRates, v.flowRate)
	if err != nil {
		panic("wtf open")
	}
	p.closedValveRates = shared.RemoveIndex(p.closedValveRates, idx)
	p.pressure = p.pressure + (timeLimit-p.timeUsed)*v.flowRate
	return
}
