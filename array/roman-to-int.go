package array

import "fmt"

func RomanToInt(s string) int {
	return 0
}

func FizzBuzz(n int) []string {
	sol := make([]string, 0)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			sol = append(sol, "FizzBuzz")
		} else if i%3 == 0 {
			sol = append(sol, "Fizz")
		} else if i%5 == 0 {
			sol = append(sol, "Buzz")
		} else {
			sol = append(sol, fmt.Sprint(i))
		}
	}

	return sol
}

func Fibonacci(n int) int {
	memo := map[int]int{}

	return fib(n, memo)
}

func fib(n int, memo map[int]int) int {
	if f, ok := memo[n]; ok {
		return f
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	sol := fib(n-1, memo) + fib(n-2, memo)

	memo[n] = sol

	return sol
}
