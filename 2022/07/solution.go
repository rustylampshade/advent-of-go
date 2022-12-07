package main

import (
	"fmt"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

type Directory struct {
	name      string
	children  []*Directory
	files     map[string]int
	totalSize int
}

var stack []*Directory

func cwd() *Directory {
	return stack[len(stack)-1]
}

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	lines := shared.Splitlines("./input.txt")

	filesystem := Directory{
		name:  "/",
		files: make(map[string]int),
	}
	stack = append(stack, &filesystem)

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if tokens[0] == "$" {
			runCommand(tokens[1], tokens[2:])
		} else {
			cwd().ls(line)
		}
	}

	used_space := filesystem.computeSize()
	part1int := 0
	for _, dirsize := range filesystem.traverseGetLTE(100_000) {
		part1int += dirsize
	}
	required_to_delete := 30_000_000 - (70_000_000 - used_space)
	part2int := 70_000_000
	for _, dirsize := range filesystem.traverseGetGTE(required_to_delete) {
		if dirsize < part2int {
			part2int = dirsize
		}
	}

	return fmt.Sprint(part1int), fmt.Sprint(part2int)
}

func runCommand(cmd string, args []string) {
	if cmd == "cd" {
		target := args[0]
		switch target {
		case "/":
			stack = stack[0:1]
		case "..":
			stack = stack[0 : len(stack)-1]
		default:
			for i := 0; i < len(cwd().children); i++ {
				if cwd().children[i].name == target {
					stack = append(stack, cwd().children[i])
					break
				}
			}
			// If we didn't find it in our children, need to create
			// LOL this never happens in the input, don't need to bother.
		}
	}
}

func (d *Directory) ls(line string) {
	tokens := strings.Split(line, " ")
	if tokens[0] == "dir" {
		dirname := tokens[1]
		for i := 0; i < len(d.children); i++ {
			if d.children[i].name == dirname {
				return
			}
		}
		d.children = append(d.children, &Directory{name: dirname, files: make(map[string]int)})
	} else {
		filesize, filename := shared.Atoi(tokens[0]), tokens[1]
		d.files[filename] = filesize
	}
}

func (d *Directory) computeSize() int {
	size := 0
	for _, v := range d.files {
		size += v
	}
	if len(d.children) != 0 {
		for i := 0; i < len(d.children); i++ {
			size += d.children[i].computeSize()
		}
	}
	d.totalSize = size
	return size
}

func (d *Directory) traverseGetGTE(cutoff int) map[string]int {
	ret := make(map[string]int)
	if len(d.children) != 0 {
		for i := 0; i < len(d.children); i++ {
			for k, v := range d.children[i].traverseGetGTE(cutoff) {
				ret[k] = v
			}
		}
	}
	if d.totalSize >= cutoff {
		ret[d.name] = d.totalSize
	}
	return ret
}

func (d *Directory) traverseGetLTE(cutoff int) map[string]int {
	ret := make(map[string]int)
	if len(d.children) != 0 {
		for i := 0; i < len(d.children); i++ {
			for k, v := range d.children[i].traverseGetLTE(cutoff) {
				ret[k] = v
			}
		}
	}
	if d.totalSize <= cutoff {
		ret[d.name] = d.totalSize
	}
	return ret
}
