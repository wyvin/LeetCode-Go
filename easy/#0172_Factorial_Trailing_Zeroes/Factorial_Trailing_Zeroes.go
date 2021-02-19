package _0172_Factorial_Trailing_Zeroes

// 偶数 * 5 肯定会产生一个0 （阶乘5前肯定是一个偶数）
// 变相求n里面有多少个5的倍数的数，要注意的是这个数是否又是5的倍数
func trailingZeroes(n int) int {
	s := 0
	for n != 0 {
		n /= 5
		s += n
	}
	return s
}
