package sort

// returns new slice with input values sorted in ascending order
func Simple(input []int) []int {
	l := len(input)
	ret := make([]int, l)

	for i := 0; i < l; i++ {
		s := smallestAt(input)
		ret[i] = input[s]
		input = pop(s, input)
	}

	return ret
}

// return index of the smallest value in given input
func smallestAt(input []int) int {
	ret := 0
	smallest := input[0]

	for i := 1; i < len(input); i++ {
		if input[i] < smallest {
			smallest, ret = input[i], i
		}
	}
	return ret
}

// pops index i from []int
func pop(i int, from []int) []int {
	if i < 0 || i >= len(from) {
		return from
	}

	ret := make([]int, 0, len(from)-1)

	if i == 0 {
		return append(ret, from[1:]...)
	}

	if i == len(from)-1 {
		return append(ret, from[:len(from)-1]...)
	}

	ret = append(ret, from[:i]...)
	ret = append(ret, from[i+1:]...)

	return ret
}
