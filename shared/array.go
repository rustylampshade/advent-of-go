package shared

// Return the maximum element (index and value) of this array of integers.
func Max(array []int) (index int, value int) {
	var max_i int
	var max_v int
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