package array_test

import (
	"testing"

	"github.com/edgarSucre/challenge/array"
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
		if !equal(tc.expected, actual) {
			t.Errorf("Missmatch, expected: %v, actual: %v\n", tc.expected, actual)
		}
	}
}

func equal(expected, actual []int) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i, v := range expected {
		if actual[i] != v {
			return false
		}
	}

	return true
}
