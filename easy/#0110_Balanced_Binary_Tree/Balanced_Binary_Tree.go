package _0110_Balanced_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2021-01-26
// 从下到上递归，任意一个节点不是平衡则返回-1
func isBalanced(root *TreeNode) bool {
	return high(root) >= 0
}

func high(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHigh := high(node.Left)
	if leftHigh == -1 {
		return -1
	}
	rightHigh := high(node.Right)
	if rightHigh == -1 {
		return -1
	}
	diff := leftHigh - rightHigh
	if diff >= 2 || diff <= -2 {
		return -1
	}
	if diff > 0 {
		return leftHigh + 1
	}
	return rightHigh + 1
}
