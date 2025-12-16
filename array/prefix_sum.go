package array

// Pattern: prefix-sum
// Source: https://leetcode.com/problems/two-sum/description/
func TwoSum(nums []int, target int) []int {
	prefix := map[int]int{}

	for i, v := range nums {
		if index, ok := prefix[v]; ok {
			return []int{index, i}
		}

		prefix[target-v] = i
	}

	return []int{}
}

// Pattern: prefix-sum
// Source: LeetCode 303 - https://leetcode.com/problems/range-sum-query-immutable/description/
func RangeSumQuery(nums []int, left, right int) int {
	prefix := make([]int, len(nums))

	var sum int

	for i, v := range nums {
		sum += v

		prefix[i] = sum
	}

	if left == 0 {
		return prefix[right]
	}

	return prefix[right] - prefix[left-1]
}

// Source: LeetCode 724 - https://leetcode.com/problems/find-pivot-index/description/
func FindPivotIndex(nums []int) int {
	var left, right int

	n := len(nums) - 1

	for _, v := range nums {
		right += v
	}

	if right == nums[n] {
		return n
	}

	for i, v := range nums {
		right -= v

		if left == right {
			return i
		}

		left += v
	}

	return -1
}

// Pattern: prefix-sum
// Source: LeetCode 1413 - https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/description/
func MinimumValueToGetPositiveSum(nums []int) int {
	var min, sum int

	for _, v := range nums {
		sum += v

		if sum < min {
			min = sum
		}
	}

	return 1 - min
}

// Pattern: prefix-sum
// Source: LeetCode 1422 - https://leetcode.com/problems/maximum-score-after-splitting-a-string/description/
func MaximumScoreAfterSplittingString(s string) int {
	var max, zeros, ones int

	for _, v := range s {
		if v == '1' {
			ones++
		}
	}

	for _, v := range s {
		if v == '0' {
			zeros++
		} else {
			ones--
		}

		count := zeros + ones
		if count > max {
			max = count
		}
	}

	return max
}

// this belongs somewhere else
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
	const end = 2050
	const start = 1950

	year := make([]int, end+1)

	for _, val := range logs {
		year[val[0]]++
		year[val[1]]--
	}

	res := 0

	for i := start; i < end; i++ {
		year[i] += year[i-1]

		if year[i] > year[res] {
			res = i
		}
	}

	return res
}

// Pattern: prefix-sum, carry-on
// Source: LeetCode 1893 - https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/description/
func CheckIfAllIntegersInRange(nums [][]int, left, right int) bool {
	prefix := make([]int, 50)

	for _, v := range nums {
		start, end := v[0], v[1]+1

		prefix[start]++
		prefix[end]--
	}

	for i := 1; i < 50; i++ {
		prefix[i] += prefix[i-1]
	}

	for i := left; i <= right; i++ {
		if prefix[i] < 1 {
			return false
		}
	}

	return true
}

// Pattern: prefix-sum
// Source: LeetCode 1893 - https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered/description/
func FindMiddleIndex(nums []int) int {
	var rightSum int

	for _, v := range nums {
		rightSum += v
	}

	var leftSum int

	for i, v := range nums {
		rightSum -= v

		if leftSum == rightSum {
			return i
		}

		leftSum += v
	}

	return -1
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
	var count, max int
	prefix := map[int]int{0: -1}

	for i, v := range nums {
		if v == 0 {
			count--
		} else {
			count++
		}

		if index, ok := prefix[count]; ok {
			contiguous := i - index

			if contiguous > max {
				max = contiguous
			}
		} else {
			prefix[count] = i
		}
	}

	return max
}

// Pattern: prefix-sum, map
// Source: LeetCode 560 - https://leetcode.com/problems/subarray-sum-equals-k/description/
func SubarraySumEqualsK(nums []int, k int) int {
	var count, sum int
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

	prefix := map[int]int{0: -1}

	for i, v := range nums {
		sum += v
		remainder := sum % k

		if index, ok := prefix[remainder]; ok {
			if i-index >= 2 {
				return true
			}
		} else {
			prefix[remainder] = i
		}
	}

	return false
}

// Pattern: prefix-sum, map
// Source: LeetCode 974 - https://leetcode.com/problems/subarray-sums-divisible-by-k/description/
func SubarraySumsDivisibleByK(nums []int, k int) int {
	var count, sum int

	prefix := map[int]int{0: 1}

	for _, v := range nums {
		sum += v
		remainder := sum % k

		if remainder < 0 {
			remainder += k
		}

		count += prefix[remainder]
		prefix[remainder]++
	}

	return count
}

// Pattern: prefix-sum, map, maximum subarray
// Source: LeetCode 1658 - https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/description/
func MinimumOperationsToReduceXtoZero(nums []int, k int) int {
	var sum, sumK, max int

	prefix := map[int]int{0: -1}

	for _, v := range nums {
		sumK += v
	}

	sumK -= k

	for i, v := range nums {
		sum += v

		if index, ok := prefix[sum-sumK]; ok {
			sub := i - index

			if sub > max {
				max = sub
			}
		}

		prefix[sum] = i
	}

	if max > 0 {
		return len(nums) - max
	}

	return -1
}

func MinSubArrayLen(nums []int, target int) int {
	return 0
}

func ProductExceptSelf(nums []int) []int {
	return []int{}
}

func KRadiusSubarrayAverages(nums []int, k int) []int {
	res := make([]int, len(nums))

	var sum, left int

	for right, v := range nums {
		sum += v

		if right < k {
			res[right] = -1
			continue
		}

		if len(nums)-right <= k {
			res[right] = -1
		}

		index := right - k

		if left == right-(k*2) {
			res[index] = sum / (k*2 + 1)

			sum -= nums[left]
			left++
		}
	}

	return res
}

func MaxAverage(list [][]int) int {
	sumIndex := map[int]int{}
	countIndex := map[int]int{}

	for _, classGrade := range list {
		sumIndex[classGrade[0]] += classGrade[1]
		countIndex[classGrade[0]]++
	}

	var maxAvg float64
	var sol int

	for subject, scoreSum := range sumIndex {
		count := float64(countIndex[subject])
		avg := float64(scoreSum) / count

		if maxAvg == avg && sol < subject {
			sol = subject
		} else if avg > maxAvg {
			sol = subject
			maxAvg = avg
		}
	}

	return sol
}
