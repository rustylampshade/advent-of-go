package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

type Valve struct {
	label     string
	flowRate  int
	neighbors []string
}

type Path struct {
	visited    []string
	openValves []string
	timeUsed   int
	pressure   int
}

var graph map[string]Valve
var bestPath Path
var maxFlow int
var timeLimit int

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	lines := shared.Splitlines("./input.txt")
	graph = make(map[string]Valve)
	maxFlow = 0
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		label := tokens[1]
		rate := shared.Atoi(strings.Trim(strings.Split(tokens[4], "=")[1], ";"))
		if rate > maxFlow {
			maxFlow = rate
		}
		neighbors := strings.Split(strings.Join(tokens[9:], ""), ",")
		graph[label] = Valve{label: label, flowRate: rate, neighbors: neighbors}
	}

	// Part 1
	timeLimit = 30
	startingPath := Path{visited: []string{"AA"}, timeUsed: 1}
	bestPath = Path{visited: []string{"AA"}, timeUsed: 1}
	dfs(startingPath, graph["AA"])
	part1 = fmt.Sprint(bestPath.pressure)

	timeLimit = 26
	bestPath = Path{visited: []string{"AA"}, timeUsed: 1}
	dfs(startingPath, graph["AA"])
	part2 = fmt.Sprint(bestPath.pressure)

	return
}

func dfs(path Path, v Valve) {
	if path.timeUsed >= timeLimit {
		if path.pressure > bestPath.pressure {
			bestPath = path
		}
		return
	}

	// Hypothetically what's the best we can do? Pretend that the max flow rate valve is next door. If
	// even such a rosy scenario doesn't mean we do better than our best-so-far pressure, we're on a dead
	// path and can return early.
	var bestCaseAdditionalPressure int
	for i := timeLimit - (path.timeUsed + 2) + 1; i > 0; i -= 2 {
		bestCaseAdditionalPressure += i * maxFlow
	}
	if path.pressure+bestCaseAdditionalPressure < bestPath.pressure {
		return
	}

	for _, destinationValve := range v.neighbors {
		pathIfWeOpen, err := path.visit(destinationValve, true)
		if err == nil {
			dfs(pathIfWeOpen, graph[destinationValve])
		}

		pathIfWePass, err := path.visit(destinationValve, false)
		if err == nil {
			dfs(pathIfWePass, graph[destinationValve])
		}
	}

	return
}

func (p *Path) visit(dest string, shouldOpenValve bool) (newPath Path, err error) {
	destValve := graph[dest]
	// It always takes one minute to travel to the next valve.
	newPath.timeUsed = p.timeUsed + 1
	newPath.pressure = p.pressure
	newPath.openValves = p.openValves[:]

	if shouldOpenValve {

		// The valve will be opened if it was a) requested to be opened, b) was not already opened by a previous
		// visit to this valve, c) will release steam, and d) we have enough time remaining that opening matters.
		if shared.TestIn(p.openValves, dest) || destValve.flowRate == 0 || newPath.timeUsed >= timeLimit {
			err = errors.New("invalid state")
			return
		}
		newPath.openValves = append(p.openValves, dest)
		newPath.pressure += (timeLimit - newPath.timeUsed) * destValve.flowRate
		newPath.timeUsed++
	}
	newPath.visited = append(p.visited, dest)
	return newPath, err
}
