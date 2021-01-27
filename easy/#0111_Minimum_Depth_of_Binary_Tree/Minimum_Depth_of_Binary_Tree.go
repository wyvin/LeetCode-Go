package _0111_Minimum_Depth_of_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2021-01-27
// 深度优先，遍历所有的节点
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	nodeDepth := 100000
	if root.Left != nil {
		nodeDepth = min(minDepth(root.Left), nodeDepth)
	}
	if root.Right != nil {
		nodeDepth = min(minDepth(root.Right), nodeDepth)
	}
	return nodeDepth + 1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 广度优先，找到第一个叶节点就返回
