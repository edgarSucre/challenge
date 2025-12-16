package array

func RomanToInt(s string) int {
	guide := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var n, previous int

	for i := len(s) - 1; i >= 0; i-- {
		current := guide[s[i]]
		if current < previous {
			n -= current
		} else {
			n += current
		}
		previous = current
	}

	return n
}

func FizzBuzz(n int) []string {
	sol := make([]string, 0)

	return sol
}

func Fibonacci(n int) int {
	return 0
}
