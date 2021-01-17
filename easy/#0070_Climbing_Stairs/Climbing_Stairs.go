package _0070_Climbing_Stairs

// 经典的动态规划题，官网给出了6种解法，时间复杂度最低为O(log n)，膜拜
// 动态规划
func climbStairs(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	default:
		x, y := 2, 3
		for i := 3; i < n; i++ {
			x, y = y, x+y
		}
		return y
	}
}

// 递归法 最直接最慢
func climbStairs2(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	default:
		return climbStairs2(n-1) + climbStairs2(n-2)
	}
}
