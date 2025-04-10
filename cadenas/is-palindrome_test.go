package cadenas_test

import (
	"testing"

	"github.com/edgarSucre/challenge/cadenas"
)

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		in       string
		expected bool
	}{
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
		{
			" ",
			true,
		},
	}

	for _, tc := range testCases {
		actual := cadenas.IsPalindrome(tc.in)
		if actual != tc.expected {
			t.Errorf("Missmatch, expected: %v, actual: %v\n", tc.expected, actual)
		}
	}
}
