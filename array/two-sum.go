package array

func TwoSum(nums []int, target int) []int {
	residuePosition := make(map[int]int)
	for i, v := range nums {
		if pos, ok := residuePosition[v]; ok {
			return []int{pos, i}
		}

		residuePosition[target-v] = i
	}

	return []int{}
}
