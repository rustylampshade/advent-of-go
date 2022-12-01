package shared

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// Read the input file and return an array of strings representing each line. Only basic trimming done.
func Splitlines(filename string) (result []string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	string_content := strings.Split(string(content), "\n")
	string_content = Map(string_content, func(elem string) string { return strings.TrimSuffix(elem, "\r") })

	return string_content
}

// For files that consist purely of one integer per line, this will return an
// array of those integers.
func ReadIntFromLine(filename string) (result []int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(string(content), "\n") {
		line = strings.TrimSuffix(line, "\r")
		number, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		result = append(result, number)
	}

	if len(result) == 0 {
		return nil
	}

	return
}
