package array

import "math"

func LongestPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var prefix []byte

	for i := 0; i < math.MaxInt; i++ {
		var current byte

		for _, word := range strs {
			if len(word) <= i {
				return string(prefix)
			}

			if current == 0 {
				current = word[i]

				continue
			}

			if word[i] != current {
				return string(prefix)
			}
		}

		prefix = append(prefix, current)
	}

	return string(prefix)
}
