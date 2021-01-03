package __Reverse_Integer

import "math"

// 2021-01-02
func reverse(x int) int {
	var b int
	var res int
	for x!=0 {
		b = x % 10
		x /= 10
		if res >math.MaxInt32/10 || (res==math.MaxInt32/10 && b>7) {
			return 0
		}
		if res < math.MinInt32/10 || (res==math.MinInt32/10 && b< -8) {
			return 0
		}
		res = res * 10 + b

	}
	return res
}
