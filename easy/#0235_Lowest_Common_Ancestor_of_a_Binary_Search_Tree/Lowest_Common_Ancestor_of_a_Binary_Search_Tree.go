package _0235_Lowest_Common_Ancestor_of_a_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 找出p，q到root节点的路径，然后判断
// 优化：由于题目是二叉搜索树，搜索路径时可以根据value大小判断路径方向
// 优化：根据二叉搜索树的性质，可以同时遍历p和q，且不用存储路径
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	listp := helper(root, p)
	listq := helper(root, q)
	lenp := len(listp) - 1
	lenq := len(listq) - 1
	var ans *TreeNode
	for lenp >= 0 && lenq >= 0 {
		if listp[lenp] == listq[lenq] {
			ans = listp[lenp]
			lenp--
			lenq--
		} else {
			break
		}
	}
	return ans
}

func helper(root, p *TreeNode) (res []*TreeNode) {
	if root == nil {
		return []*TreeNode{}
	}
	if root.Val == p.Val {
		res = append(res, root)
		return res
	}
	res = helper(root.Left, p)
	if len(res) > 0 {
		res = append(res, root)
		return
	}
	res = helper(root.Right, p)
	if len(res) > 0 {
		res = append(res, root)
		return
	}
	return
}

// 优化二
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	ans := root
	for {
		if p.Val > ans.Val && q.Val > ans.Val {
			ans = ans.Right
		} else if p.Val < ans.Val && q.Val < ans.Val {
			ans = ans.Left
		} else {
			break
		}
	}
	return ans
}
