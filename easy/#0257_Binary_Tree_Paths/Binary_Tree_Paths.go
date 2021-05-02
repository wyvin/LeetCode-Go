package _0257_Binary_Tree_Paths

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	leftPaths := binaryTreePaths(root.Left)
	for i := range leftPaths {
		leftPaths[i] = fmt.Sprintf("%d->", root.Val) + leftPaths[i]
	}
	rightPaths := binaryTreePaths(root.Right)
	for i := range rightPaths {
		rightPaths[i] = fmt.Sprintf("%d->", root.Val) + rightPaths[i]
	}
	return append(leftPaths, rightPaths...)
}
