package _0009_Palindrome_Number

// 2020-01-03 利用整数反转，也可以用字符串实现
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x > 10 && x%10 == 0 {
		return false
	}

	var (
		y   = x
		pop int
		res int
	)
	for y != 0 {
		pop = y % 10
		y /= 10
		res = res*10 + pop
	}
	return res == x
}
