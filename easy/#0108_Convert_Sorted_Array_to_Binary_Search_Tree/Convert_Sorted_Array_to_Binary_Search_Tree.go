package _0108_Convert_Sorted_Array_to_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2021-01-24
// 递归
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}

// 2021-01-25
// 迭代
func sortedArrayToBST1(nums []int) *TreeNode {
	switch len(nums) {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: nums[0]}
	case 2:
		return &TreeNode{
			Val:  nums[1],
			Left: &TreeNode{Val: nums[0]},
		}
	case 3:
		return &TreeNode{
			Val:   nums[1],
			Left:  &TreeNode{Val: nums[0]},
			Right: &TreeNode{Val: nums[2]},
		}
	default:
		left, right := 0, len(nums)-1
		queue := [][]int{{left, right}}
		root := new(TreeNode)
		nodeQueue := []*TreeNode{root}
		for len(queue) > 0 {
			l := len(queue)
			for i := 0; i < l; i++ {
				r := queue[i]
				left, right = r[0], r[1]
				mid := (left + right) / 2
				nodeQueue[i].Val = nums[mid]
				if left < mid {
					queue = append(queue, []int{left, mid - 1})
					nodeQueue[i].Left = &TreeNode{}
					nodeQueue = append(nodeQueue, nodeQueue[i].Left)
				}
				if right > mid {
					queue = append(queue, []int{mid + 1, right})
					nodeQueue[i].Right = &TreeNode{}
					nodeQueue = append(nodeQueue, nodeQueue[i].Right)
				}
			}
			queue = queue[l:]
			nodeQueue = nodeQueue[l:]
		}
		return root
	}
}
