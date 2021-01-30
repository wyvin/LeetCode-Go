package _0119_Pascal_s_Triangle_II

// 2021-01-30
// 每次循环从中间开始往前计算，镜像赋值
func getRow(rowIndex int) []int {
	switch rowIndex {
	case 0:
		return []int{1}
	case 1:
		return []int{1, 1}
	default:
		ans := make([]int, rowIndex+1)
		ans[0] = 1
		ans[1] = 1
		l := 3
		for rowIndex > 1 {
			mid := l / 2
			if l%2 == 0 {
				mid--
			}
			for i := mid; i >= 1; i-- {
				ans[i] += ans[i-1]
				ans[l-i-1] = ans[i]
			}
			ans[l-1] = 1
			l++
			rowIndex--
		}
		return ans
	}
}
