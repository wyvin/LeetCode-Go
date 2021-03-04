package _0009_Palindrome_Number

// 2020-01-03 利用整数反转，也可以用字符串实现
func isPalindrome(x int) bool {
	// 小于0的反转后都不是回文数。
	// 大于0的且最后一位为0的都不是回文数
	if x < 0 || (x > 0 && x%10 == 0) {
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

// 思路优化：只反转一半的数字
func isPalindrome2(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	res := 0
	// 循环的结束条件是当 res > x 时就代表已经反转一半或以上
	for x > res {
		res = res*10 + x%10
		x /= 10
	}

	// 当长度为偶数时则 x 和 res 的位数相同
	// 当长度为奇数时 res 就会比 x 多一位数字，所以要除10
	return x == res || x == res/10
}
