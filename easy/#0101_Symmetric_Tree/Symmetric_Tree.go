package _0101_Symmetric_Tree

/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2021-01-21 递归
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameNode(root.Left, root.Right)
}

func isSameNode(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l != nil && r != nil {
		if l.Val == r.Val {
			return isSameNode(l.Left, r.Right) && isSameNode(l.Right, r.Left)
		}
	}
	return false
}
