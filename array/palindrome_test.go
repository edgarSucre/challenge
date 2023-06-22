package array_test

import (
	"testing"

	"github.com/edgarSucre/challenge/array"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		in       int
		expected bool
	}{
		{
			121,
			true,
		},
		{
			-121,
			false,
		},
		{
			10,
			false,
		},
	}

	for _, tc := range testCases {
		actual := array.IsPalindrome(tc.in)
		if tc.expected != actual {
			t.Errorf("Missmatch, expected: %v, actual: %v\n", tc.expected, actual)
		}
	}
}
