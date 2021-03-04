package _0200_Number_of_Islands

var (
	xl = 0
	yl = 0
)

// 深度优先
// 其实就是循环，从第一个为'1'的坐标开始深度搜索，把它改成'0'，
// 然后判断上下左右4个位置是否为'1'，是则递归把整个大陆都改成0，
// 下次循环时就不会遇到了
func numIslands(grid [][]byte) int {
	ans := 0
	xl = len(grid) - 1
	yl = len(grid[0]) - 1
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == '1' {
				ans++
				helper(grid, i, j)
			}
		}
	}
	return ans
}

func helper(grid [][]byte, i, j int) {

	grid[i][j] = '0'

	if i+1 <= xl && grid[i+1][j] == '1' {
		helper(grid, i+1, j)
	}
	if i-1 >= 0 && grid[i-1][j] == '1' {
		helper(grid, i-1, j)
	}
	if j+1 <= yl && grid[i][j+1] == '1' {
		helper(grid, i, j+1)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		helper(grid, i, j-1)
	}
}

// 广度优先
// 和深度优先类似，不同的是检查边界的方法
// 有点难理解
// 就是用一个队列，把每个'1'的坐标的上下左右都放进去，直到队列为空则继续主循环
// 可以理解成按对角线的方向做广度优先
