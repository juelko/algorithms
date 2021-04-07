package sort

// simple quick sort implementation
func quick(data []int, lo, hi int) {

	if hi-lo < 1 {

		return
	}
	if hi-lo == 1 {
		if data[lo] > data[hi] {
			data[lo], data[hi] = data[hi], data[lo]
			return
		}
		return
	}

	pivot := data[lo]

	a, b := lo+1, hi

	for a <= b {
		if data[a] <= pivot {
			data[a], data[a-1] = data[a-1], data[a]
			a++
		} else {
			data[b], data[a] = data[a], data[b]
			b--
		}
	}
	if lo < b-1 {
		quick(data, lo, b-1)
	}
	if a < hi {
		quick(data, a, hi)
	}
}
