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

func TestMaximumScoreAfterSplittingString(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"011101", 5},
		{"00111", 5},
		{"1111", 3},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			actual := array.MaximumScoreAfterSplittingString(tt.s)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSumAllOddLengthSubarrays(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{1, 4, 2, 5, 3},
			58,
		},
		{
			[]int{1, 2},
			3,
		},
		{
			[]int{10, 11, 12},
			66,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.SumAllOddLengthSubarrays(tt.nums)

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

func TestMaximumPopulationYear(t *testing.T) {
	tests := []struct {
		logs     [][]int
		expected int
	}{
		{
			[][]int{
				{1993, 1999},
				{2000, 2010},
			},
			1993,
		},
		{
			[][]int{
				{1950, 1961},
				{1960, 1971},
				{1970, 1981},
			},
			1960,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.logs), func(t *testing.T) {
			actual := array.MaximumPopulationYear(tt.logs)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCheckIfAllIntegersInRange(t *testing.T) {
	tests := []struct {
		nums     [][]int
		left     int
		right    int
		expected bool
	}{
		{
			[][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			2,
			5,
			true,
		},
		{
			[][]int{
				{1, 10},
				{10, 20},
			},
			21,
			21,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.CheckIfAllIntegersInRange(tt.nums, tt.left, tt.right)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFindMiddleIndex(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{2, 3, -1, 8, 4},
			3,
		},
		{
			[]int{1, -1, 4},
			2,
		},
		{
			[]int{2, 5},
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.FindMiddleIndex(tt.nums)

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

func TestMaxAverage(t *testing.T) {
	r := array.MaxAverage([][]int{
		{2, 4},
		{2, 10},
		{3, 5},
		{3, 7},
		{5, 4},
		{5, 8},
		{5, 9},
	})

	assert.Equal(t, 5, r)
}

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			[]int{2, 3, 1, 2, 4, 3},
			7,
			2,
		},
		{
			[]int{1, 4, 4},
			4,
			1,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1},
			11,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.MinSubArrayLen(tt.nums, tt.target)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestProductsExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			actual := array.ProductExceptSelf(tt.nums)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
