package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

var monkeys map[string]string

func solve() (part1 string, part2 string) {

	monkeys = make(map[string]string)
	for _, line := range shared.Splitlines("./input.txt") {
		tokens := strings.Split(line, " ")
		name := strings.Trim(tokens[0], ":")
		monkeys[name] = strings.Join(tokens[1:], " ")
	}

	part1 = fmt.Sprint(yellNumber("root"))

	var target, operand int
	treeWithHuman, target := getCloser("root")
	for treeWithHuman != "humn" {
		var temp string
		temp, operand = getCloser(treeWithHuman)
		target = reverseOperation(treeWithHuman, target, operand)
		treeWithHuman = temp
	}
	part2 = fmt.Sprint(target)
	return
}

func getCloser(monkey string) (human string, target int) {
	tokens := strings.Split(monkeys[monkey], " ")
	left, right := tokens[0], tokens[2]
	if containsHuman(left) {
		target = yellNumber(right)
		return left, target
	} else {
		target = yellNumber(left)
		return right, target
	}
}

func containsHuman(name string) bool {
	// fmt.Printf("Checking if %v contains the human\n", name)
	if name == "humn" {
		return true
	}
	job := monkeys[name]
	tokens := strings.Split(job, " ")
	if len(tokens) == 1 {
		return false
	}
	monkeyA, monkeyB := tokens[0], tokens[2]
	if monkeyA == "humn" || monkeyB == "humn" {
		return true
	} else {
		return containsHuman(monkeyA) || containsHuman(monkeyB)
	}
}

func yellNumber(name string) (n int) {
	job := monkeys[name]
	n, err := strconv.Atoi(job)
	if err == nil {
		return
	}

	tokens := strings.Split(job, " ")
	monkeyA, operation, monkeyB := tokens[0], tokens[1], tokens[2]
	switch operation {
	case "+":
		n = yellNumber(monkeyA) + yellNumber(monkeyB)
	case "-":
		n = yellNumber(monkeyA) - yellNumber(monkeyB)
	case "*":
		n = yellNumber(monkeyA) * yellNumber(monkeyB)
	case "/":
		n = yellNumber(monkeyA) / yellNumber(monkeyB)
	default:
		panic(job)
	}

	return
}

func reverseOperation(tree string, target int, operand int) int {
	tokens := strings.Split(monkeys[tree], " ")
	left, operation := tokens[0], tokens[1]
	switch operation {
	case "+":
		return target - operand
	case "*":
		return target / operand
	case "-":
		if containsHuman(left) {
			return target + operand
		} else {
			return operand - target
		}
	case "/":
		if containsHuman(left) {
			return target * operand
		} else {
			return operand / target
		}
	}
	panic("?")
}
