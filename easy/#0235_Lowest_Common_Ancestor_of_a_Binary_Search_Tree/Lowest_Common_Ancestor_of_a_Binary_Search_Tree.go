package _0235_Lowest_Common_Ancestor_of_a_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
