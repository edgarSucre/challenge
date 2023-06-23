package array_test

import (
	"testing"

	"github.com/edgarSucre/challenge/array"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		in       string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tc := range testCases {
		actual := array.RomanToInt(tc.in)
		if actual != tc.expected {
			t.Errorf("Missmatch, expected: %v, actual: %v\n", tc.expected, actual)
		}
	}
}
