package array_test

import (
	"testing"

	"github.com/edgarSucre/challenge/array"
)

func TestLongestPrefix(t *testing.T) {
	testCases := []struct {
		in       []string
		expected string
	}{
		{
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			[]string{"dog", "racecar", "car"},
			"",
		},
	}

	for _, tc := range testCases {
		actual := array.LongestPrefix(tc.in)
		if actual != tc.expected {
			t.Errorf("Missmatch, expected: %#v, actual: %#v\n", tc.expected, actual)
		}
	}
}
