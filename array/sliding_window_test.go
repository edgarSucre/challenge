package array_test

import (
	"fmt"
	"testing"

	"github.com/edgarSucre/challenge/array"
	"github.com/stretchr/testify/assert"
)

func TestMinimumPositiveSumSubarray(t *testing.T) {
	tests := []struct {
		nums     []int
		left     int
		right    int
		expected int
	}{
		{
			[]int{3, -2, 1, 4},
			2,
			3,
			1,
		},
		{
			[]int{-2, 2, -3, 1},
			2,
			3,
			-1,
		},
		{
			[]int{1, 2, 3, 4},
			2,
			4,
			3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.MinimumPositiveSumSubarray(tt.nums, tt.left, tt.right)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			actual := array.LongestSubstringWithoutRepeatingCharacters(tt.s)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMaximumAverageSubarray(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{
			[]int{1, 12, -5, -6, 50, 3},
			4,
			12.75000,
		},
		{
			[]int{5},
			1,
			5.00,
		},
	}
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.MaximumAverageSubarray(tt.nums, tt.k)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinimumDifferenceBetweenHighestAndLowestKScores(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			[]int{90},
			1,
			0,
		},
		{
			[]int{9, 4, 1, 7},
			2,
			2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.MinimumDifferenceBetweenHighestAndLowestKScores(tt.nums, tt.k)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

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
