package sort

// return index of the smallest value in given input
func smallestAt(input []int) int {
	var ret int
	smallest := input[0]

	for i, v := range input[1:] {
		if v < smallest {
			smallest, ret = v, i
		}
	}
	return ret
}

// returns new array with input values sorted in ascending order
func selectioSort(input []int) []int {
	ret := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		ret = append(ret, input[smallestAt(input[i:])])
	}

	return ret
}
