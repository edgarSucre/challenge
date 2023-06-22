package hmap_test

import (
	"testing"

	hmap "github.com/edgarSucre/challenge/hash-map"
)

func Test_RansomNote(t *testing.T) {
	testcases := []struct {
		note     string
		mag      string
		expected bool
	}{
		{
			note:     "aa",
			mag:      "ab",
			expected: false,
		},
		{
			note:     "aa",
			mag:      "aab",
			expected: true,
		},
	}

	for _, tc := range testcases {
		actual := hmap.RansomNote(tc.note, tc.mag)
		if tc.expected != actual {
			t.Errorf("Missmatch, expected: %v, actual: %v\n", tc.expected, tc.expected)
		}
	}
}
