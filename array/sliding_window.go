package array

import (
	"math"
	"sort"
)

func MinimumPositiveSumSubarray(nums []int, left, right int) int {
	var sum int

	min := math.MaxInt

	prefixSum := make([]int, len(nums))

	for i, v := range nums {
		sum += v
		prefixSum[i] = sum
	}

	for i := 0; i+left < len(nums); i++ {
		for j := i + left - 1; j <= i+right-1 && j < len(nums); j++ {
			if i == 0 {
				sum = prefixSum[j]
			} else {
				sum = prefixSum[j] - prefixSum[i-1]
			}

			if sum > 0 && sum < min {
				min = sum
			}
		}
	}

	if min < math.MaxInt {
		return min
	}

	return -1
}

func LongestSubstringWithoutRepeatingCharacters(s string) int {
	index := map[byte]struct{}{}

	longest := 0

	i, j := 0, 0
	for i < len(s) && j < len(s) {
		if _, ok := index[s[j]]; ok {
			delete(index, s[i])

			i++
			continue
		}

		index[s[j]] = struct{}{}
		w := j - i + 1

		if w > longest {
			longest = w
		}

		j++
	}

	return longest
}

func MaximumAverageSubarray(nums []int, k int) float64 {
	maxAverage := float64(0)

	prefixSum := make([]int, len(nums))
	sum := 0

	for i, v := range nums {
		sum += v
		prefixSum[i] = sum
	}

	for left, right := 0, k-1; right < len(nums); left, right = left+1, right+1 {
		if left == 0 {
			sum = prefixSum[right]
		} else {
			sum = prefixSum[right] - prefixSum[left-1]
		}

		avg := float64(sum) / float64(k)

		if avg > maxAverage {
			maxAverage = avg
		}
	}

	return maxAverage
}

// O(nlog(n))
func MinimumDifferenceBetweenHighestAndLowestKScores(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}

	sort.Ints(nums)

	minDiff := math.MaxInt

	for i, j := 0, k-1; j < len(nums); i, j = i+1, j+1 {
		diff := nums[j] - nums[i]

		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

// Pattern: sliding window
// Source: LeetCode 713 - https://leetcode.com/problems/subarray-product-less-than-k/description/
func SubarrayProductLessThanK(nums []int, k int) int {
	var left, result int

	product := 1

	for right, v := range nums {
		product *= v

		for left <= right && product >= k {
			product /= nums[left]

			left++
		}

		result += right - left + 1
	}

	return result
}

// Pattern: sliding window
// Source: LeetCode 209 - https://leetcode.com/problems/minimum-size-subarray-sum/description/
func MinimumSizeSubArraySum(nums []int, k int) int {
	var left, min, sum int

	for right, v := range nums {
		sum += v

		for left <= right && sum >= k {
			min = right - left + 1
			sum -= nums[left]
			left++
		}
	}

	return min
}
