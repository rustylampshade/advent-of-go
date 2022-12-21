package main

import (
	"fmt"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

type Blueprint struct {
	num                                         int
	oreRobotOreCost                             int
	clayRobotOreCost                            int
	obsidianRobotOreCost, obsidianRobotClayCost int
	geodeRobotOreCost, geodeRobotObsidianCost   int
}

type State struct {
	time                                           int // Full minutes elapsed
	bp                                             *Blueprint
	oreRobot, clayRobot, obsidianRobot, geodeRobot int
	ore, clay, obsidian, geode                     int
	buildSequence, declined                        []string
	mostRecentRobot                                string
}

var copyCount, simulationCount int
var mostGeodes int

func (s *State) Copy() (ns *State) {
	copyCount++
	var newState State
	newState.time = s.time
	newState.bp = s.bp
	newState.oreRobot = s.oreRobot
	newState.clayRobot = s.clayRobot
	newState.obsidianRobot = s.obsidianRobot
	newState.geodeRobot = s.geodeRobot
	newState.ore = s.ore
	newState.clay = s.clay
	newState.obsidian = s.obsidian
	newState.geode = s.geode
	newState.buildSequence = s.buildSequence[:]
	newState.declined = s.declined[:]
	newState.mostRecentRobot = s.mostRecentRobot
	return &newState
}

func (s *State) SpendMaterials(robotToBuild string) (underConstruction string) {
	switch robotToBuild {
	case "geode":
		if s.ore >= s.bp.geodeRobotOreCost && s.obsidian >= s.bp.geodeRobotObsidianCost {
			s.ore -= s.bp.geodeRobotOreCost
			s.obsidian -= s.bp.geodeRobotObsidianCost
			s.mostRecentRobot = "geode"
			s.declined = make([]string, 0)
			return "geode"
		}
	case "obsidian":
		if s.ore >= s.bp.obsidianRobotOreCost && s.clay >= s.bp.obsidianRobotClayCost {
			s.ore -= s.bp.obsidianRobotOreCost
			s.clay -= s.bp.obsidianRobotClayCost
			s.mostRecentRobot = "obsidian"
			s.declined = make([]string, 0)
			return "obsidian"
		}
	case "clay":
		if s.ore >= s.bp.clayRobotOreCost {
			s.ore -= s.bp.clayRobotOreCost
			s.mostRecentRobot = "clay"
			s.declined = make([]string, 0)
			return "clay"
		}
	case "ore":
		if s.ore >= s.bp.oreRobotOreCost {
			s.ore -= s.bp.oreRobotOreCost
			s.mostRecentRobot = "ore"
			s.declined = make([]string, 0)
			return "ore"
		}
	case "nothing":
		if s.ore >= s.bp.obsidianRobotOreCost && s.clay >= s.bp.obsidianRobotClayCost {
			s.declined = append(s.declined, "obsidian")
		}
		if s.ore >= s.bp.clayRobotOreCost {
			s.declined = append(s.declined, "clay")
		}
		if s.ore >= s.bp.oreRobotOreCost {
			s.declined = append(s.declined, "ore")
		}
	}
	return
}

func (s *State) Options() (options []string) {
	if s.ore >= s.bp.geodeRobotOreCost && s.obsidian >= s.bp.geodeRobotObsidianCost {
		options = append(options, "geode")
		// Return early if we can make a geode.
		return
	}
	if s.ore >= s.bp.obsidianRobotOreCost && s.clay >= s.bp.obsidianRobotClayCost && !shared.TestIn(s.declined, "obsidian") {
		options = append(options, "obsidian")
	}
	if s.ore >= s.bp.clayRobotOreCost && !shared.TestIn(s.declined, "clay") {
		options = append(options, "clay")
	}
	if s.ore >= s.bp.oreRobotOreCost && !shared.TestIn(s.declined, "ore") {
		options = append(options, "ore")
	}
	if len(options) < 3 {
		options = append(options, "nothing")
	}
	return
}

func (s *State) GatherMaterials() {
	s.ore += s.oreRobot
	s.clay += s.clayRobot
	s.obsidian += s.obsidianRobot
	s.geode += s.geodeRobot
}

func (s *State) AddRobot(robot string) {
	switch robot {
	case "geode":
		s.geodeRobot++
	case "obsidian":
		s.obsidianRobot++
	case "clay":
		s.clayRobot++
	case "ore":
		s.oreRobot++
	}
}

func (s *State) simulateAll(maxRounds int) {
	if s.time == maxRounds {
		if s.geode > mostGeodes {
			mostGeodes = s.geode
			// fmt.Println(s)
		}
		return
	}

	// s.time is [0..maxRounds), so timeRemaining is the number of complete rounds left to simulate.
	timeRemaining := maxRounds - s.time
	// The most geodes I can possibly have at the end of this simulation is
	// most = s.geodes + (n)*(n-1)/2 + s.geodeRobots
	// Representing how many I currently have (no effect with time), how many I could
	// produce over the upcoming n turns if I produce a new geode robot per turn, and
	// a factor for how many robots I already have right now.
	if s.geode+timeRemaining*(timeRemaining-1)/2+(s.geodeRobot*timeRemaining) < mostGeodes {
		return
	}

	simulationCount++
	s.time++
	options := s.Options()
	for _, robotToBuild := range options {
		var newState *State
		if len(options) > 1 {
			newState = s.Copy()
		} else {
			newState = s
		}
		newState.buildSequence = append(newState.buildSequence, robotToBuild)
		underConstruction := newState.SpendMaterials(robotToBuild)
		newState.GatherMaterials()
		newState.AddRobot(underConstruction)
		newState.simulateAll(maxRounds)
	}
}

func main() {
	part1, part2 := solve()

	fmt.Println(simulationCount, copyCount)
	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {

	quality, mult := 0, 1
	for i, line := range shared.Splitlines("./input.txt") {
		tokens := strings.Split(line, " ")
		blueprint := Blueprint{
			num:                    i + 1,
			oreRobotOreCost:        shared.Atoi(tokens[6]),
			clayRobotOreCost:       shared.Atoi(tokens[12]),
			obsidianRobotOreCost:   shared.Atoi(tokens[18]),
			obsidianRobotClayCost:  shared.Atoi(tokens[21]),
			geodeRobotOreCost:      shared.Atoi(tokens[27]),
			geodeRobotObsidianCost: shared.Atoi(tokens[30])}

		p1 := State{oreRobot: 1, bp: &blueprint}
		mostGeodes = 0
		p1.simulateAll(24)
		quality += (mostGeodes * blueprint.num)

		p2 := State{oreRobot: 1, bp: &blueprint}
		mostGeodes = 0
		if blueprint.num <= 3 {
			p2.simulateAll(32)
			mult *= mostGeodes
		}

	}
	return fmt.Sprint(quality), fmt.Sprint(mult)
}
