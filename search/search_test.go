package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		value int
		from  []int
		want  int
	}{
		"found":     {value: 2, from: []int{1, 2, 3, 4}, want: 1},
		"not found": {value: 5, from: []int{1, 2, 3, 4}, want: -1},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Simple(tc.value, tc.from)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestBinary(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		value int
		from  []int
		want  int
	}{
		"found":     {value: 2, from: []int{1, 2, 3, 4}, want: 1},
		"not found": {value: 5, from: []int{1, 2, 3, 4}, want: -1},
	}
	for name := range tests {
		tc := tests[name]
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Binary(tc.value, tc.from)

			require.Equal(t, tc.want, got)
		})
	}
}
