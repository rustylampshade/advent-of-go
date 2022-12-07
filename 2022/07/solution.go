package main

import (
	"fmt"
	"strings"

	"github.com/rustylampshade/advent-of-go/shared"
)

type Directory struct {
	name      string // Only the basename, not the full path.
	parent    *Directory
	children  []*Directory
	files     map[string]int
	totalSize int // Computed lazily
}
type Command struct {
	cmd         string
	args        []string
	outputStart int
	outputEnd   int
}

var filesystem Directory
var stack []*Directory

func cwd() *Directory {
	return stack[len(stack)-1]
}

// Create the root-based filesystem and init the stack pointer here.
func initialize() {
	filesystem = Directory{
		name:  "/",
		files: make(map[string]int),
	}
	stack = append(stack, &filesystem)
}

func main() {
	part1, part2 := solve()

	fmt.Println("Solution for part 1: " + part1)
	fmt.Println("Solution for part 2: " + part2)
}

func solve() (part1 string, part2 string) {
	initialize()

	lines := shared.Splitlines("./input.txt")
	for _, command := range getCommandsFromInput(lines) {
		command.run(&lines)
	}

	used_space := filesystem.computeSize()
	part1int := 0
	for _, dirsize := range filesystem.findPart1() {
		part1int += dirsize
	}

	required_to_delete := 30_000_000 - (70_000_000 - used_space)
	part2int := 70_000_000
	for _, dirsize := range filesystem.findPart2(required_to_delete) {
		if dirsize < part2int {
			part2int = dirsize
		}
	}

	return fmt.Sprint(part1int), fmt.Sprint(part2int)
}

// Process the input to chunk it out into Command objects, each of which knows what
// command/args/output is associated with it.
// Careful with the output index tracking.
func getCommandsFromInput(lines []string) (commands []Command) {
	var start, end int
	for i, line := range lines {
		tokens := strings.Split(line, " ")
		if tokens[0] == "$" {
			if len(commands) > 0 {
				commands[len(commands)-1].outputStart = start
				commands[len(commands)-1].outputEnd = end
			}
			commands = append(commands, Command{cmd: tokens[1], args: tokens[2:]})
			start, end = i+1, i+1
		} else {
			end += 1
		}
	}
	commands[len(commands)-1].outputStart = start
	commands[len(commands)-1].outputEnd = end
	return
}

// Run the given command.
func (command Command) run(lines *[]string) {
	switch command.cmd {

	case "cd":
		cwd().cd(command.args[0])

	case "ls":
		cwd().ls(lines, command.outputStart, command.outputEnd)
	}
}

// Change directory. Mutates the stack pointer and filesystem tree.
func (d *Directory) cd(target string) {
	switch target {

	case "/":
		stack = stack[0:1]

	case "..":
		stack = stack[0 : len(stack)-1]

	default:
		for _, child := range d.children {
			if child.name == target {
				stack = append(stack, child)
				return
			}
		}
		newdir := &Directory{name: target, parent: cwd(), files: make(map[string]int)}
		d.children = append(d.children, newdir)
		stack = append(stack, newdir)
	}
}

// List directory contents. Mutates the filesystem tree as more info is learned.
func (d *Directory) ls(sharedLines *[]string, start int, end int) {
	for _, outputLine := range (*sharedLines)[start:end] {
		tokens := strings.Split(outputLine, " ")
		if tokens[0] == "dir" {
			dirname := tokens[1]
			var existing_children []string
			for _, child := range d.children {
				existing_children = append(existing_children, child.name)
			}
			if !shared.TestIn(existing_children, dirname) {
				d.children = append(d.children, &Directory{name: dirname, parent: cwd(), files: make(map[string]int)})
			}
		} else {
			filesize, filename := shared.Atoi(tokens[0]), tokens[1]
			d.files[filename] = filesize
		}
	}
}

// Recurse through this Directory and its children recomputing the totalSize of each node.
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

// Find directories with totalSize >= a given cutoff.
func (d *Directory) findPart2(cutoff int) map[string]int {
	ret := make(map[string]int)
	if len(d.children) != 0 {
		for i := 0; i < len(d.children); i++ {
			for k, v := range d.children[i].findPart2(cutoff) {
				ret[k] = v
			}
		}
	}
	if d.totalSize >= cutoff {
		ret[d.name] = d.totalSize
	}
	return ret
}

// Find directories with totalSize <= 100_000
func (d *Directory) findPart1() map[string]int {
	ret := make(map[string]int)
	if len(d.children) != 0 {
		for i := 0; i < len(d.children); i++ {
			for k, v := range d.children[i].findPart1() {
				ret[k] = v
			}
		}
	}
	if d.totalSize <= 100_000 {
		ret[d.name] = d.totalSize
	}
	return ret
}
