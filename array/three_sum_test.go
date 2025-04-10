package array_test

import (
	"fmt"
	"testing"

	"github.com/edgarSucre/challenge/array"
	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			},
		},
		{
			[]int{0, 1, 1},
			[][]int{},
		},
		{
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		desc := fmt.Sprintf("%+v", tt.nums)

		t.Run(desc, func(t *testing.T) {
			actual := array.ThreeSum(tt.nums)

			assert.ElementsMatch(t, tt.expected, actual)
		})
	}
}
