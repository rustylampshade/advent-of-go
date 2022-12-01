package shared

import (
	"log"
	"os"
	"strconv"
	"strings"
)

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
