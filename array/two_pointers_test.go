package array_test

import (
	"testing"

	"github.com/edgarSucre/challenge/array"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeS(t *testing.T) {
	tests := []struct {
		s        string
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

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.expected, array.IsPalindromeS(tt.s))
		})
	}
}
