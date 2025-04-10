package array_test

import (
	"fmt"
	"testing"

	"github.com/edgarSucre/challenge/array"
	"github.com/stretchr/testify/assert"
)

func TestAlternateElements(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50}

	expected := []int{10, 30, 50}
	actual := array.AlternateElements(arr)

	assert.Equal(t, expected, actual)

	arr = []int{-5, 1, 4, 2, 12}
	expected = []int{-5, 4, 12}
	actual = array.AlternateElements(arr)

	assert.Equal(t, expected, actual)
}

func TestSecondLargest(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{
			[]int{12, 35, 1, 10, 34, 1},
			34,
		},
		{
			[]int{10, 5, 10},
			5,
		},
		{
			[]int{10, 10, 10},
			-1,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case #%v", i+1), func(t *testing.T) {
			actual := array.SecondLargest(tt.arr)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLargestThree(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			[]int{10, 4, 3, 50, 23, 90},
			[]int{90, 50, 23},
		},
		{
			[]int{10, 9, 9},
			[]int{10, 9},
		},
		{
			[]int{10, 10, 10},
			[]int{10},
		},
		{
			[]int{},
			[]int{},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case #%v", i+1), func(t *testing.T) {
			actual := array.LargestThree(tt.arr)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLeaders(t *testing.T) {
	arr := []int{16, 17, 4, 3, 5, 2}

	expected := []int{17, 5, 2}
	actual := array.Leaders(arr)

	assert.Equal(t, expected, actual)
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		arr      []int
		expected bool
	}{
		{
			[]int{20, 21, 45, 89, 89, 90},
			true,
		},
		{
			[]int{20, 20, 45, 89, 89, 90},
			true,
		},
		{
			[]int{20, 20, 78, 98, 99, 97},
			false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case #%v", i+1), func(t *testing.T) {
			actual := array.IsSorted(tt.arr)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			[]int{2, 2, 2, 2, 2},
			[]int{2},
		},
		{
			[]int{1, 2, 2, 3, 4, 4, 4, 5, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case #%v", i+1), func(t *testing.T) {
			actual := array.RemoveDuplicates(tt.arr)

			assert.Equal(t, tt.expected, actual)
		})
	}

}

func TestGenerateSubArrays(t *testing.T) {
	arr := []int{1, 2, 3}

	expected := [][]int{
		{1},
		{1, 2},
		{2},
		{1, 2, 3},
		{2, 3},
		{3},
	}

	actual := array.GenerateSubArrays(arr)

	assert.ElementsMatch(t, expected, actual)
}

func TestReverseArray(t *testing.T) {
	arr := []int{1, 4, 3, 2, 6, 5}

	expected := []int{5, 6, 2, 3, 4, 1}

	array.Reverse(arr)

	assert.Equal(t, expected, arr)
}

func TestRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}

	//array.Rotate(arr, 2)
	array.RotateV2(arr, 2)
	assert.Equal(t, []int{5, 6, 1, 2, 3, 4}, arr)
}

func TestZerosToTheEnd(t *testing.T) {
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}

	array.ZerosToTheEnd(arr)
	assert.Equal(t, []int{1, 2, 4, 3, 5, 0, 0, 0}, arr)
}

func TestMinimumIncrementForEqual(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected int
	}{
		{
			[]int{4, 7, 19, 16},
			3,
			10,
		},
		{
			[]int{4, 4, 4, 4},
			3,
			0,
		},
		{
			[]int{4, 2, 6, 8},
			3,
			-1,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case #%v", i+1), func(t *testing.T) {
			actual := array.MinimumIncrementForEqual(tt.arr, tt.k)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinimumCostToMakeSizeOne(t *testing.T) {
	arr := []int{4, 3, 2}

	expected := 4

	assert.Equal(t, expected, array.MinimumCostToMakeSizeOne(arr))
}

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	tests := []struct {
		target   int
		expected int
	}{
		{2, 0},
		{5, 1},
		{8, 2},
		{12, 3},
		{16, 4},
		{23, 5},
		{38, 6},
		{56, 7},
		{72, 8},
		{91, 9},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("recursive, target: %v", tt.target), func(t *testing.T) {
			actual := array.BinarySearchRecursive(arr, tt.target)

			assert.Equal(t, tt.expected, actual)
		})

		t.Run(fmt.Sprintf("iterative, target: %v", tt.target), func(t *testing.T) {
			actual := array.BinarySearchIterative(arr, tt.target)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{9, 4, 3, 8, 10, 2, 5}

	expected := []int{2, 3, 4, 5, 8, 9, 10}

	array.HeapSort(arr, false)

	assert.Equal(t, expected, arr)

	arr = []int{9, 4, 3, 8, 10, 2, 5}
	array.HeapSort(arr, true)

	expected = []int{10, 9, 8, 5, 4, 3, 2}
	assert.Equal(t, expected, arr)
}
