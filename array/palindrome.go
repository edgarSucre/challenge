package array

func IsPalindrome(n int) bool {
	if n < 0 {
		return false
	}

	var reversed int
	x := n
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return reversed == n
}
