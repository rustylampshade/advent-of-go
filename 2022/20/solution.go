package main

import (
	"fmt"

	"github.com/rustylampshade/advent-of-go/shared"
)

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

type Number struct {
	n             int
	originalIndex int
}

var encryptedFile []Number
var fileLength int

func solve() (part1 string, part2 string) {

	for i, line := range shared.Splitlines("./input.txt") {
		encryptedFile = append(encryptedFile, Number{n: shared.Atoi(line), originalIndex: i})
		fileLength++
	}
	for i := 0; i < fileLength; i++ {
		var numToMix Number
		var currentIndex int
		for currentIndex, numToMix = range encryptedFile {
			if numToMix.originalIndex == i {
				break
			}
		}
		Mix(currentIndex, numToMix.n)
	}
	indexOfZero := 0
	for i, num := range encryptedFile {
		if num.n == 0 {
			indexOfZero = i
			break
		}
	}
	p1coordinates := encryptedFile[(indexOfZero+1000)%fileLength].n + encryptedFile[(indexOfZero+2000)%fileLength].n + encryptedFile[(indexOfZero+3000)%fileLength].n

	encryptedFile = make([]Number, 0)
	decryptionKey := 811589153
	for i, line := range shared.Splitlines("./input.txt") {
		encryptedFile = append(encryptedFile, Number{n: shared.Atoi(line) * decryptionKey, originalIndex: i})
	}
	for iteration := 0; iteration < 10; iteration++ {
		for i := 0; i < fileLength; i++ {
			var numToMix Number
			var currentIndex int
			for currentIndex, numToMix = range encryptedFile {
				if numToMix.originalIndex == i {
					break
				}
			}
			Mix(currentIndex, numToMix.n)
		}
	}
	for i, num := range encryptedFile {
		if num.n == 0 {
			indexOfZero = i
			break
		}
	}
	p2coordinates := encryptedFile[(indexOfZero+1000)%fileLength].n + encryptedFile[(indexOfZero+2000)%fileLength].n + encryptedFile[(indexOfZero+3000)%fileLength].n

	return fmt.Sprint(p1coordinates), fmt.Sprint(p2coordinates)
}

func Mix(current, shift int) {
	if shift == 0 {
		return
	}

	withMovingElementRemoved := make([]Number, 0)
	withMovingElementRemoved = append(withMovingElementRemoved, encryptedFile[0:current]...)
	if current != len(encryptedFile)-1 {
		withMovingElementRemoved = append(withMovingElementRemoved, encryptedFile[current+1:]...)
	}
	if len(withMovingElementRemoved) != len(encryptedFile)-1 {
		panic("bad math")
	}

	target := current + shift
	for target < 0 {
		// Increment by a silly amount so the following modulo does the math quickly.
		target += 1_000_000_000 * len(withMovingElementRemoved)
	}
	if target >= fileLength {
		target = target % len(withMovingElementRemoved)
	}

	newOrder := make([]Number, 0)
	newOrder = append(newOrder, withMovingElementRemoved[0:target]...)
	newOrder = append(newOrder, encryptedFile[current])
	newOrder = append(newOrder, withMovingElementRemoved[target:]...)
	encryptedFile = newOrder
}
