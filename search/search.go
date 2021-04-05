package search

// Simple search: loops through values in slice from and returns
// index of the match or -1 if not found
// linear time algorithm O(n)
// works for unsorted slices
func simple(value int, from []int) int {
	for i, v := range from {
		if v == value {
			return i
		}
	}
	return -1
}

// Binary search: Checks the value against middle point of indexes
// Elinimates half of the indexes after each iteration
// returns -1 if value not found
// logarithmic time algorithm O(log n)
// works only for sorted slices
func binary(value int, from []int) int {
	low, high := 0, len(from)-1

	for low <= high {
		mid := (low + high) / 2

		if from[mid] == value {
			return mid
		}

		if value > from[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
