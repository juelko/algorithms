package recursion

// recursive factorial
func fact(x int) int {
	// base case
	if x <= 1 {
		return 1
	}

	// recursive case
	return x * fact(x-1)
}

// sums values in input
func sum(input []int) int {
	if len(input) == 0 {
		return 0
	}
	if len(input) == 1 {
		return input[0]
	}

	return input[0] + sum(input[1:])
}

// item of linked list
type item struct {
	value int
	next  *item
}

// counts number of items in linked list
func count(head item) int {
	if head.next == nil {
		// end of the list, base case
		return 1
	}

	return 1 + count(*head.next)
}

// finds max value in linked list
func max(head item) int {
	if head.next == nil {
		// end of the list, base case
		return head.value
	}

	if head.value < head.next.value {
		// next is larger, call max with next
		return max(*head.next)
	}
	// next is smaller, carry current value and skip next
	return max(item{value: head.value, next: head.next.next})
}
