package array

func IsPalindrome(n int) bool {
	if n < 0 {
		return false
	}

	x := n

	var reversed int

	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return reversed == n
}
