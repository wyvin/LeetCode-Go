package __Reverse_Integer

// 2021-01-02
func reverse(x int) int {
	var b int
	var res int
	for x!=0 {
		b = x % 10
		x /= 10
		res = res * 10 + b
	}
	if res > 2<<30-1 || res < -2<<30 {
		return 0
	}
	return res
}
