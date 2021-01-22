package _0104_Maximum_Depth_of_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2021-01-22
// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

// 2021-01-22
// 广度优先搜索
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		// 每循环一层计算器加一
		ans++
		// 每层有多少个节点
		tmp := len(queue)
		for i := 0; i < tmp; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		// 替换为下层的节点
		queue = queue[tmp:]
	}
	return ans
}
