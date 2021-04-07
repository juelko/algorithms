package sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSmallestAt(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{2, 2, 3, 1, 2}, want: 3},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := smallestAt(tc.input)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestSelectionSort(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"1": {input: []int{4, 2, 3, 1}, want: []int{1, 2, 3, 4}},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := selection(tc.input)

			require.Equal(t, tc.want, got)
		})
	}
}
