package array

import (
	"slices"
)

func AlternateElements(arr []int) []int {
	resp := make([]int, 0, len(arr))

	for i := 0; i < len(arr); i += 2 {
		resp = append(resp, arr[i])
	}

	return resp
}

func SecondLargest(arr []int) int {
	largest, secondLargest := -1, -1

	for _, v := range arr {
		if v > largest {
			secondLargest = largest
			largest = v
		} else if v > secondLargest && v < largest {
			secondLargest = v
		}
	}

	return secondLargest
}

func LargestThree(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	resp := make([]int, 3)

	for _, v := range arr {
		if v > resp[0] {
			resp[2] = resp[1]
			resp[1] = resp[0]
			resp[0] = v
		}

		if v > resp[1] && resp[0] > v {
			resp[2] = resp[1]
			resp[1] = v
		}

		if v > resp[2] && resp[1] > v {
			resp[2] = v
		}
	}

	if resp[2] == 0 {
		resp = resp[:2]
	}

	if resp[1] == 0 {
		resp = resp[:1]
	}

	if resp[0] == 0 {
		resp = []int{}
	}

	return resp
}

func Leaders(arr []int) []int {
	leaders := make([]int, 1, len(arr))

	n := len(arr) - 1
	max := arr[n]

	leaders[0] = max

	for i := n - 1; i >= 0; i-- {
		if arr[i] >= max {
			max = arr[i]
			leaders = append(leaders, max)
		}
	}

	slices.Reverse(leaders)

	return leaders
}

func IsSorted(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}

	return true
}

func RemoveDuplicates(arr []int) []int {
	result := make([]int, 0, len(arr))
	result = append(result, arr[0])

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			result = append(result, arr[i])
		}
	}

	return result
}

func GenerateSubArrays(arr []int) [][]int {
	res := make([][]int, 0)
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sub := make([]int, 0)
			for k := i; k <= j; k++ {
				sub = append(sub, arr[k])
			}

			res = append(res, sub)
		}
	}

	return res
}

func Reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func Rotate(arr []int, n int) {
	for i := 0; i < n; i++ {
		swap := arr[len(arr)-1]
		for j := 0; j < len(arr); j++ {
			arr[j], swap = swap, arr[j]
		}
	}
}

func RotateV2(arr []int, n int) {
	l := len(arr)

	temp := append(arr[(l-n):], arr[:(l-n)]...)

	for i := 0; i < l; i++ {
		arr[i] = temp[i]
	}
}

func ZerosToTheEnd(arr []int) {
	var count int

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			arr[i], arr[count] = arr[count], arr[i]
			count++
		}
	}
}

func MinimumIncrementForEqual(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	var res int

	for _, v := range arr {
		if (max-v)%k != 0 {
			return -1
		}

		res += (max - v) / k
	}

	return res
}

// https://www.geeksforgeeks.org/minimum-cost-make-array-size-1-removing-larger-pairs/
func MinimumCostToMakeSizeOne(arr []int) int {
	min := slices.Min(arr)

	return (len(arr) - 1) * min
}

func BinarySearchRecursive(arr []int, target int) int {
	return binarySearchRecursive(arr, 0, len(arr)-1, target)
}

func binarySearchRecursive(arr []int, min, max, target int) int {
	if min > max {
		return -1
	}

	mid := min + (max-min)/2

	if arr[mid] > target {
		return binarySearchRecursive(arr, 0, mid-1, target)
	} else if arr[min] < target {
		return binarySearchRecursive(arr, mid+1, max, target)
	} else {
		return mid
	}
}

func BinarySearchIterative(arr []int, target int) int {
	min := 0
	max := len(arr) - 1

	for min <= max {
		mid := min + (max-min)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	return -1
}

func parent(n int) int {
	return (n - 1) / 2
}

func left(root int) int {
	return root*2 + 1
}

func right(root int) int {
	return root*2 + 2
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type lessFn func(arr []int, i, j int) bool

func heapIfy(arr []int, last, root int, less lessFn) {
	left, right := left(root), right(root)

	largest := root

	if left < last && less(arr, left, largest) {
		largest = left
	}

	if right < last && less(arr, right, largest) {
		largest = right
	}

	if largest != root {
		swap(arr, root, largest)

		heapIfy(arr, last, largest, less)
	}
}

func max(arr []int, i, j int) bool {
	return arr[i] > arr[j]
}

func min(arr []int, i, j int) bool {
	return arr[i] < arr[j]
}

func HeapSort(arr []int, desc bool) {
	n := len(arr)

	less := max
	if desc {
		less = min
	}

	for i := parent(n); i >= 0; i-- {
		heapIfy(arr, n, i, less)
	}

	for i := n - 1; i >= 0; i-- {
		swap(arr, 0, i)

		heapIfy(arr, i, 0, less)
	}
}
