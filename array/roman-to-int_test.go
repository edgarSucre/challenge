package array_test

import (
	"fmt"
	"testing"

	"github.com/edgarSucre/challenge/array"
	"github.com/stretchr/testify/assert"
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

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		opts     int
		expected []string
	}{
		{
			3,
			[]string{"1", "2", "Fizz"},
		},
		{
			5,
			[]string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			15,
			[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.opts), func(t *testing.T) {
			actual := array.FizzBuzz(tt.opts)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			actual := array.Fibonacci(tt.n)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
