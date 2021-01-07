package _0021_Merge_Two_Sorted_Lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	res := new(ListNode)
	head := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			res.Next = l2
			break
		}
		if l2 == nil {
			res.Next = l1
			break
		}
		if l1.Val <= l2.Val {
			res.Next = l1
			l1 = l1.Next
		} else {
			res.Next = l2
			l2 = l2.Next
		}
		res = res.Next
	}
	return head.Next
}
