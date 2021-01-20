package _0100_Same_Tree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2021-01-20
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
			return true
		}
	}
	return false
}
