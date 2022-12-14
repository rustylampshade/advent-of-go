package shared

import (
	"errors"
	"fmt"
)

func Reverse[V int | string](array []V) []V {
	backwards := make([]V, len(array))
	for i, j := 0, len(array)-1; i < len(array); i, j = i+1, j-1 {
		backwards[j] = array[i]
	}
	return backwards
}

// Return the maximum element (index and value) of this array of integers.
func Max[V int | string](array []V) (index int, value V) {
	max_i := -1
	var max_v V
	for i, v := range array {
		if max_v < v || i == 0 {
			max_i = i
			max_v = v
		}
	}
	return max_i, max_v
}

// Return the minimum element (index and value) of this array of integers.
func Min(array []int) (index int, value int) {
	var min_i int
	var min_v int
	for i, v := range array {
		if v < min_v || i == 0 {
			min_i = i
			min_v = v
		}
	}
	return min_i, min_v
}

// Return the sum of this array of integers.
func Sum(array []int) (result int) {
	result = 0
	for _, e := range array {
		result += e
	}
	return result
}

// Pop the final n elements off this collection.
func Pop[V int | string](array []V, n int) (popped []V, remaining []V) {
	if n > len(array) {
		panic(fmt.Sprintf("Cannot pop %d elements from an array shorter than that: len=%d", n, len(array)))
	}
	return array[len(array)-n:], array[:len(array)-n]
}

func RemoveIndex[V int | string](s []V, index int) []V {
	// Use make to avoid cloberring the underlying array beneath the slices.
	ret := make([]V, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// Test if the element `elem` is In this collection
func TestIn[V int | string](array []V, elem V) bool {
	for _, v := range array {
		if v == elem {
			return true
		}
	}
	return false
}

// Test if this collection entirely consists of unique elements. This could be
// implemented more generically with a Counts() call and then checking the frequencies
// to ensure all are 1, but directly implementing this here lets us short-circuit and
// return early if there are duplicates without needing to process the entire collection.
func TestEntirelyUnique[V int | string](array []V) bool {
	seen := map[V]bool{}
	for _, v := range array {
		if seen[v] {
			return false
		} else {
			seen[v] = true
		}
	}
	return true
}

// Return a map of the frequency each element occurs with in this collection.
func Counts[V int | string](array []V) (counts map[V]int) {
	counts = make(map[V]int)
	for _, v := range array {
		counts[v] += 1
	}
	return counts
}

// Return the indices of all occurences of `elem` in this list.
func FindAll[V int | string](array []V, elem V) []int {
	var locations []int
	for i, v := range array {
		if v == elem {
			locations = append(locations, i)
		}
	}
	return locations
}

// Return the index of the FIRST occurence of `elem` in this list
func FindFirst[V int | string](array []V, elem V) (idx int, err error) {
	for i, v := range array {
		if v == elem {
			idx = i
			return
		}
	}
	err = errors.New("unable to find element")
	return
}

// Apply the function f to each element of this array of integers or strings.
// Questionably useful right now, since it cannot be used for casting since that returns a different
// return type from the input array. It also can't be used when `f` needs multiple parameters.
func Map[V int | string](vs []V, f func(V) V) []V {
	vsm := make([]V, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
