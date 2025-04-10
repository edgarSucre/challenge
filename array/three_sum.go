package array

import "sort"

// Pattern: sorting + two pointers
// Source: https://leetcode.com/problems/3sum/
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	results := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum < 0 {
				j++
				continue
			}

			if sum > 0 {
				k--
				continue
			}

			triplet := []int{nums[i], nums[j], nums[k]}
			results = append(results, triplet)

			j++
			k--

			for j < k && nums[j] == nums[j-1] {
				j++
			}
			for j < k && nums[k] == nums[k+1] {
				k--
			}
		}
	}

	return results
}
