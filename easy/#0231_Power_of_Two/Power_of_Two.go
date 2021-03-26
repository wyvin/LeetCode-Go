package _0231_Power_of_Two

// 1000 -> 100 -> 10 -> 1 -> 0 true
// 1100 -> 110 -> 11 -> 1 false
func isPowerOfTwo(n int) bool {
	for n > 0 {
		if n&1 == 1 {
			return n>>1 == 0
		}
		n >>= 1
	}
	return false
}

// 1000 & 111 (1000-1) = 000 true
// 1100 & 1011 (1100-1) = 1000 false
func isPowerOfTwo2(n int) bool {
	if n == 0 {
		return false
	}
	return (n & (n - 1)) == 0
}
