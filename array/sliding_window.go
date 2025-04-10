package array

// Pattern: sliding window
// Source: LeetCode 713 - https://leetcode.com/problems/subarray-product-less-than-k/description/
func SubarrayProductLessThanK(nums []int, k int) int {
	var res, left int

	product := 1

	for right, v := range nums {
		product *= v

		for left <= right && product >= k {
			product /= nums[left]
			left++
		}

		res += right - left + 1
	}

	return res
}

// Pattern: sliding window
// Source: LeetCode 209 - https://leetcode.com/problems/minimum-size-subarray-sum/description/
func MinimumSizeSubArraySum(nums []int, k int) int {
	var sum, left int

	min := len(nums) + 1

	for right, v := range nums {
		sum += v

		for left <= right && sum >= k {
			min = right - left + 1

			sum -= nums[left]
			left++
		}
	}

	if min == len(nums)+1 {
		return 0
	}

	return min
}
