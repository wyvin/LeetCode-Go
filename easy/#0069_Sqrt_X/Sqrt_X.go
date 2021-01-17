package _0069_Sqrt_X

// 2020-01-16
// 类似二分查找
func mySqrt(x int) int {
	switch x {
	case 0:
		return 0
	case 1, 2, 3:
		return 1
	default:
		i, j := 2, x/2
		for i <= j {
			mid := (i + j) / 2
			s := mid * mid
			if s == x {
				return mid
			}
			if s > x {
				j = mid - 1
			}
			if s < x {
				i = mid + 1
			}
		}
		return i - 1
	}
}

// 最直接最慢的方法
func mySqrt2(x int) int {
	ans := 0
	for ans < x {
		sqrt := ans * ans
		switch {
		case sqrt == x:
			return ans
		case sqrt > x:
			return ans - 1
		case sqrt < x:
			ans++
		}
	}
	return ans
}
