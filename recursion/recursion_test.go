package recursion

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFactorial(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input int
		want  int
	}{
		"0": {input: 0, want: 1},
		"1": {input: 1, want: 1},
		"2": {input: 2, want: 2},
		"3": {input: 3, want: 6},
		"4": {input: 4, want: 24},
		"5": {input: 5, want: 120},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fact(tc.input)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestSum(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input []int
		want  int
	}{
		"basic": {input: []int{1, 2, 3, 4}, want: 10},
		"one":   {input: []int{5}, want: 5},
		"empty": {input: []int{}, want: 0},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := sum(tc.input)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestCount(t *testing.T) {
	t.Parallel()
	tail := item{value: 1}

	tests := map[string]struct {
		input item
		want  int
	}{
		"one":   {input: tail, want: 1},
		"two":   {input: item{value: 2, next: &tail}, want: 2},
		"three": {input: item{value: 3, next: &item{value: 2, next: &tail}}, want: 3},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := count(tc.input)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestMax(t *testing.T) {
	t.Parallel()
	tail := item{value: 10}

	tests := map[string]struct {
		input item
		want  int
	}{
		"one":   {input: tail, want: 10},
		"two":   {input: item{value: 20, next: &tail}, want: 20},
		"three": {input: item{value: 13, next: &item{value: 20, next: &tail}}, want: 20},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := max(tc.input)

			require.Equal(t, tc.want, got)
		})
	}
}
