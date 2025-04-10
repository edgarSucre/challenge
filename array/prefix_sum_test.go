package array_test

import (
	"fmt"
	"testing"

	"github.com/edgarSucre/challenge/array"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tc := range testCases {
		actual := array.TwoSum(tc.nums, tc.target)

		assert.ElementsMatch(t, tc.expected, actual)
	}
}

func TestFindPivotIndex(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{1, 7, 3, 6, 5, 6},
			3,
		},
		{
			[]int{1, 2, 3},
			-1,
		},
		{
			[]int{2, 1, -1},
			0,
		},
		{
			[]int{1, -1, 99},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.FindPivotIndex(tt.nums)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinimumValueToGetPositiveSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{-3, 2, -3, 4, 2},
			5,
		},
		{
			[]int{1, 2},
			1,
		},
		{
			[]int{1, -2, -3},
			5,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.MinimumValueToGetPositiveSum(tt.nums)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestRangeSumQuery(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}

	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{left: 0, right: 2, expected: 1},
		{left: 2, right: 5, expected: -1},
		{left: 0, right: 5, expected: -3},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("from %v to %v", tt.left, tt.right), func(t *testing.T) {
			actual := array.RangeSumQuery(nums, tt.left, tt.right)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestContiguousArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{0, 1},
			2,
		},
		{
			[]int{0, 1, 0},
			2,
		},
		{
			[]int{0, 1, 1, 1, 1, 1, 0, 0, 0},
			6,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.ContiguousArray(tt.nums)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSubarraySumEqualsK(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1},
			2,
			8,
		},
		{
			[]int{1, 2, 3},
			3,
			2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.SubarraySumEqualsK(tt.nums, tt.k)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestContinuousSubArraySum(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{
			[]int{23, 2, 4, 6, 7},
			6,
			true,
		},
		{
			[]int{23, 2, 6, 4, 7},
			6,
			true,
		},
		{
			[]int{23, 2, 6, 4, 7},
			13,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.ContinuousSubArraySum(tt.nums, tt.k)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSubarraySumsDivisibleByK(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			[]int{4, 5, 0, -2, -3, 1},
			5,
			7,
		},
		{
			[]int{5},
			9,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.SubarraySumsDivisibleByK(tt.nums, tt.k)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinimumOperationsToReduceXtoZero(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			[]int{1, 1, 4, 2, 3},
			5,
			2,
		},
		{
			[]int{5, 6, 7, 8, 9},
			4,
			-1,
		},
		{
			[]int{3, 2, 20, 1, 1, 3},
			10,
			5,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.MinimumOperationsToReduceXtoZero(tt.nums, tt.k)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestKRadiusSubarrayAverages(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			[]int{7, 4, 3, 9, 1, 8, 5, 2, 6},
			3,
			[]int{-1, -1, -1, 5, 4, 4, -1, -1, -1},
		},
		{
			[]int{100000},
			0,
			[]int{100000},
		},
		{
			[]int{8},
			100000,
			[]int{-1},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.KRadiusSubarrayAverages(tt.nums, tt.k)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
