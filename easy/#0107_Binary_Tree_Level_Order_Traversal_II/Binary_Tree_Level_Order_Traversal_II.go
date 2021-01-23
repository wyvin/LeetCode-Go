package _0107_Binary_Tree_Level_Order_Traversal_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2021-01-23
// 广度优先搜索
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ans [][]int
	tmpNode := []*TreeNode{root}
	for len(tmpNode) > 0 {
		l := len(tmpNode)
		tmp := make([]int, l)
		for i := 0; i < l; i++ {
			tmp[i] = tmpNode[i].Val
			if tmpNode[i].Left != nil {
				tmpNode = append(tmpNode, tmpNode[i].Left)
			}
			if tmpNode[i].Right != nil {
				tmpNode = append(tmpNode, tmpNode[i].Right)
			}
		}
		ans = append(ans, tmp)
		// 用以下插入的话会产生野数组
		// ans = append([][]int{tmp},ans...)
		tmpNode = tmpNode[l:]
	}
	for x := 0; x < len(ans)/2; x++ {
		ans[x], ans[len(ans)-x-1] = ans[len(ans)-x-1], ans[x]
	}
	return ans
}

// 2021-01-23
// 递归
func levelOrderBottom1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 1)
	ans = addQueue(root, 0, ans)
	for x := 0; x < len(ans)/2; x++ {
		ans[x], ans[len(ans)-x-1] = ans[len(ans)-x-1], ans[x]
	}
	return ans
}

func addQueue(node *TreeNode, level int, ans [][]int) [][]int {
	ans[level] = append(ans[level], node.Val)
	if node.Left == nil && node.Right == nil {
		return ans
	}
	if len(ans)-1 == level {
		ans = append(ans, []int{})
	}
	if node.Left != nil {
		ans = addQueue(node.Left, level+1, ans)
	}
	if node.Right != nil {
		ans = addQueue(node.Right, level+1, ans)
	}
	return ans
}
