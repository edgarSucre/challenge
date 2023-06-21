package linklist_test

import (
	"testing"

	linklist "github.com/edgarSucre/challenge/linked-list"
)

func TestMiddleNode(t *testing.T) {
	testCases := []struct {
		in       []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			3,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			4,
		},
	}

	for _, tc := range testCases {
		h := linklist.ListFromslice(tc.in)
		m := linklist.MiddleNode(h)

		if m == nil {
			t.Errorf("Middle is nil\n")
			return
		}

		if m.Val != tc.expected {
			t.Errorf("Missmatch, expected: %v, actual: %v\n", tc.expected, m.Val)
		}
	}
}
