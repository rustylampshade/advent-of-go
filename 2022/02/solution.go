package main

import (
	"fmt"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

var winning_play = [4]int{-1, 2, 3, 1}
var losing_play = [4]int{-1, 3, 1, 2}

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	lines := shared.Splitlines("./input.txt")
	part1_score := 0
	part2_score := 0
	for _, game := range lines {
		players := strings.Split(game, " ")
		if len(players) != 2 {
			panic("Exactly two players required for a tic-tac-toe game")
		}
		// Shift everything to integers. Rock=1, Paper=2, Scissors=3 for each player.
		elf_choice, my_choice := int(byte(players[0][0])-64), int(byte(players[1][0])-23-64)

		part1_score += my_choice + game_result(elf_choice, my_choice)

		// Elf was wrong, reinterpret the second column as the desired game result
		desired_result := my_choice
		switch desired_result {
		case 1:
			my_choice = losing_play[elf_choice]
		case 2:
			my_choice = elf_choice
		case 3:
			my_choice = winning_play[elf_choice]
		}
		part2_score += my_choice + game_result(elf_choice, my_choice)
	}

	return fmt.Sprint(part1_score), fmt.Sprint(part2_score)
}

func game_result(elf_choice int, my_choice int) (score int) {
	if elf_choice == my_choice {
		return 3
	}
	if my_choice == winning_play[elf_choice] {
		return 6
	}
	return 0
}
