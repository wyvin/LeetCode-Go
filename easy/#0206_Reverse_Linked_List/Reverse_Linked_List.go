package _0206_Reverse_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

// 用一个列表保存
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	vl := make([]int, 0)
	for head != nil {
		vl = append(vl, head.Val)
		head = head.Next
	}
	head = new(ListNode)
	ans := head

	for i := len(vl) - 1; i >= 0; i-- {
		head.Next = &ListNode{
			Val:  vl[i],
			Next: nil,
		}
		head = head.Next
	}
	return ans.Next
}

// 直接迭代
// 把next指向前一个元素
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	var curr *ListNode
	for head != nil {
		curr = head
		head = head.Next
		curr.Next = prev
		prev = curr
	}
	return curr
}
