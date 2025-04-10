package array_test

import (
	"fmt"
	"testing"

	"github.com/edgarSucre/challenge/array"
	"github.com/stretchr/testify/assert"
)

func TestSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			[]int{10, 5, 2, 6},
			100,
			8,
		},
		{
			[]int{1, 2, 3},
			0,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.SubarrayProductLessThanK(tt.nums, tt.k)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinimumSizeSubArraySum(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			[]int{2, 3, 1, 2, 4, 3},
			7,
			2,
		},
		{
			[]int{1, 4, 4},
			4,
			1,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1},
			11,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.MinimumSizeSubArraySum(tt.nums, tt.k)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
