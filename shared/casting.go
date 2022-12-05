package shared

import (
	"fmt"
	"strconv"
)

func Atoi(c string) int {
	n, err := strconv.Atoi(c)
	if err != nil {
		panic(fmt.Sprintf("Could not cast %s to an integer", c))
	}
	return n
}
