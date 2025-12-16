package array

import "math"

// Pattern: prefix-sum
// Source: https://leetcode.com/problems/two-sum/description/
func TwoSum(nums []int, target int) []int {
	sumPrefix := map[int]int{}

	// i + j == target
	//target - i == j

	for i, v := range nums {
		if j, ok := sumPrefix[v]; ok {
			return []int{i, j}
		}

		sumPrefix[target-v] = i
	}

	return []int{}
}

// Pattern: prefix-sum
// Source: LeetCode 303 - https://leetcode.com/problems/range-sum-query-immutable/description/
func RangeSumQuery(nums []int, left, right int) int {
	sumPrefix := make([]int, len(nums))

	var sum int
	for i, v := range nums {
		sum += v
		sumPrefix[i] = sum
	}

	if left == 0 {
		return sumPrefix[right]
	}

	return sumPrefix[right] - sumPrefix[left-1]
}

// Source: LeetCode 724 - https://leetcode.com/problems/find-pivot-index/description/
func FindPivotIndex(nums []int) int {
	sumPrefix := make([]int, len(nums))

	var sum int
	for i, num := range nums {
		sum += num
		sumPrefix[i] = sum
	}

	if sum-nums[0] == 0 {
		return 0
	}

	n := len(nums) - 1

	if sum-nums[n] == 0 {
		return n
	}

	for i, v := range nums {
		left := sumPrefix[i] - v
		right := sum - sumPrefix[i]

		if left == right {
			return i
		}
	}

	return -1
}

// Pattern: prefix-sum
// Source: LeetCode 1413 - https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/description/
func MinimumValueToGetPositiveSum(nums []int) int {
	var sum int

	min := math.MaxInt

	for _, v := range nums {
		sum += v

		if sum < min {
			min = sum
		}
	}

	if min > 0 {
		return min
	}

	return min*(-1) + 1
}

// Pattern: prefix-sum
// Source: LeetCode 1422 - https://leetcode.com/problems/maximum-score-after-splitting-a-string/description/
func MaximumScoreAfterSplittingString(s string) int {
	var zeroes, ones, sol int

	for _, v := range s {
		if v == '1' {
			ones++
		}
	}

	for _, v := range s {
		if v == '0' {
			ones++
		} else {
			zeroes--
		}

		sol = max(sol, zeroes+ones)
	}

	return sol
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func SumAllOddLengthSubarrays(nums []int) int {
	var sum int

	n := len(nums)

	for i, v := range nums {
		count := (((i+1)*(n-i) + 1) / 2)
		sum += v * count
	}

	return sum
}

// Pattern: prefix-sum, carry-on
// Source: LeetCode 1854 - https://leetcode.com/problems/maximum-population-year/description/
func MaximumPopulationYear(logs [][]int) int {
	sumPrefix := map[int]int{}

	for _, birthDeath := range logs {
		birth, death := birthDeath[0], birthDeath[1]
		sumPrefix[birth]++
		sumPrefix[death]--
	}

	var sum, max, year int
	for i := 1950; i <= 2050; i++ {
		sum = sum + sumPrefix[i]

		if sum > max {
			max = sum
			year = i
		}
	}

	return year
}

// Pattern: prefix-sum, carry-on
// Source: LeetCode 1893 - https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/description/
func CheckIfAllIntegersInRange(nums [][]int, left, right int) bool {
	prefixSum := make([]int, 52)

	for _, numRange := range nums {
		left, right := numRange[0], numRange[1]+1

		prefixSum[left]++
		prefixSum[right]--
	}

	for i := 1; i <= 50; i++ {
		prefixSum[i] += prefixSum[i-1]
	}

	for i := left; i <= right; i++ {
		if prefixSum[i] < 1 {
			return false
		}
	}

	return true
}

func FindMiddleIndex(nums []int) int {
	return 0
}

/*
Similar problem statements:
Find the max length subarray which sums to a desired sum.
Find the max length subarray which sums to a 0.
Find the max length subarray which has equal numbers of odd and even numbers.
Find the max length subarray which has equal numbers of 0 and 1s

Pattern: prefix-sum, map
Source: LeetCode 525 - https://leetcode.com/problems/contiguous-array/description/
*/
func ContiguousArray(nums []int) int {
	var count, m int

	prefixSum := map[int]int{0: -1}

	for i, v := range nums {
		if v == 0 {
			count--
		} else {
			count++
		}

		if index, ok := prefixSum[count]; ok {
			contiguous := i - index

			m = max(m, contiguous)
		} else {
			prefixSum[count] = i
		}
	}

	return m
}

// Pattern: prefix-sum, map
// Source: LeetCode 560 - https://leetcode.com/problems/subarray-sum-equals-k/description/
func SubarraySumEqualsK(nums []int, k int) int {
	var sum, count int

	prefix := map[int]int{0: 1}

	for _, v := range nums {
		sum += v

		count += prefix[sum-k]

		prefix[sum]++
	}

	return count
}

// Pattern: prefix-sum, map
// Source: LeetCode 523 - https://leetcode.com/problems/continuous-subarray-sum/description/
func ContinuousSubArraySum(nums []int, k int) bool {
	var sum int

	prefixSum := map[int]int{0: -1}

	for i, v := range nums {
		sum += v

		residue := sum % k

		if index, ok := prefixSum[residue]; ok {
			if i-index >= 2 {
				return true
			}
		} else {
			prefixSum[residue] = i
		}
	}

	return false
}

// Pattern: prefix-sum, map
// Source: LeetCode 974 - https://leetcode.com/problems/subarray-sums-divisible-by-k/description/
func SubarraySumsDivisibleByK(nums []int, k int) int {
	var sum, count int

	prefixSum := map[int]int{0: 1}

	for _, v := range nums {
		sum += v

		residue := sum % k

		if residue < 0 {
			residue += k
		}

		count += prefixSum[residue]

		prefixSum[residue]++
	}

	return count
}

// Pattern: prefix-sum, map, maximum subarray
// Source: LeetCode 1658 - https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/description/
func MinimumOperationsToReduceXtoZero(nums []int, k int) int {
	var sum, sumK, max int

	for _, v := range nums {
		sumK += v
	}

	sumK -= k

	prefixSum := map[int]int{0: -1}

	for i, v := range nums {
		sum += v

		if index, ok := prefixSum[sum-sumK]; ok {
			sub := i - index

			if sub > max {
				max = sub
			}
		}

		prefixSum[sum] = i
	}

	if max > 0 {
		return len(nums) - max
	}

	return -1
}

func MinSubArrayLen(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	min := math.MaxInt

	for left := 0; left < len(nums); left++ {
		if nums[left] == target {
			return 1
		}

		sum := nums[left]
		count := 1
		for right := left + 1; right < len(nums); right++ {
			sum += nums[right]
			count++

			if sum >= target {
				if count < min {
					min = count
				}

				break
			}
		}
	}

	if min == math.MaxInt {
		return 0
	}

	return min
}

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= right
		right *= nums[i]
	}

	return ans
}

func KRadiusSubarrayAverages(nums []int, k int) []int {
	return []int{}
}

func MaxAverage(list [][]int) int {
	var sol int

	return sol
}
