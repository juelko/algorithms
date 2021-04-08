package sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"simple": {input: []int{4, 2, 3, 1}, want: []int{1, 2, 3, 4}},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Simple(tc.input)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestQuick(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"simple": {input: []int{4, 3, 6, 2, 1, 7, 9, 1}, want: []int{1, 1, 2, 3, 4, 6, 7, 9}},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			Quick(tc.input)

			require.Equal(t, tc.want, tc.input)
		})
	}
}

func TestSmallestAt(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input []int
		want  int
	}{
		"simple": {input: []int{2, 2, 3, 1, 2}, want: 3},
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

func TestPop(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		data  []int
		index int
		want  []int
	}{
		"pop first":      {index: 0, data: []int{1, 2, 3, 4}, want: []int{2, 3, 4}},
		"pop second":     {index: 1, data: []int{1, 2, 3, 4}, want: []int{1, 3, 4}},
		"negative index": {index: -1, data: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
		"index overflow": {index: 4, data: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := pop(tc.index, tc.data)

			require.Equal(t, tc.want, got)
		})
	}
}
