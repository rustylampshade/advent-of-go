package shared

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadIntFromLine(filename string) (result []int) {
	content, _ := ioutil.ReadFile(filename)

	for _, line := range strings.Split(string(content), "\n") {
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
