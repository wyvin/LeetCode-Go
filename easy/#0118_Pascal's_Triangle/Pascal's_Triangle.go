package _0118_Pascal_s_Triangle

// 2021-01-29
func generate(numRows int) [][]int {
	switch numRows {
	case 0:
		return [][]int{}
	case 1:
		return [][]int{{1}}
	case 2:
		return [][]int{{1}, {1, 1}}
	default:
		ans := make([][]int, numRows)
		ans[0] = []int{1}
		ans[1] = []int{1, 1}
		for i := 2; numRows-i > 0; i++ {
			ans[i] = make([]int, i+1)
			ans[i][0] = 1
			ans[i][i] = 1
			// 只循环一半
			for x := 1; x <= i/2+i%2; x++ {
				ans[i][x] = ans[i-1][x] + ans[i-1][x-1]
				ans[i][i-x] = ans[i][x]
			}
		}
		return ans
	}
}
