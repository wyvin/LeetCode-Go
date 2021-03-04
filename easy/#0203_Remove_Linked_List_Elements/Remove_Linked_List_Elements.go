package _0203_Remove_Linked_List_Elements

type ListNode struct {
	Val  int
	Next *ListNode
}

// 先求出头节点：遇到第一个不为目标值的节点
// 记录前节点和当前节点，循环替换
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	ans := head
	lastNode := head
	curNode := head.Next
	for curNode != nil {
		if curNode.Val != val {
			lastNode = curNode
		} else {
			lastNode.Next = curNode.Next
		}
		curNode = curNode.Next
	}
	return ans
}

// 另一个比较巧妙的处理方式就是在链表前增加一个哨兵（伪头）
// 哨兵 -> head
func removeElements2(head *ListNode, val int) *ListNode {
	greeter := &ListNode{
		Val:  0,
		Next: head,
	}

	lastNode := greeter
	curNode := greeter.Next
	for curNode != nil {
		if curNode.Val != val {
			lastNode = curNode
		} else {
			lastNode.Next = curNode.Next
		}
		curNode = curNode.Next
	}
	return greeter.Next
}
