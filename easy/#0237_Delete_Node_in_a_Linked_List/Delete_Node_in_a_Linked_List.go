package _0237_Delete_Node_in_a_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

// 注意：题目不提供head，只提供head其中的一个节点
func deleteNode(node *ListNode) {
	// 不可以这样写，这样只是改变node变量名指向的地址，并不会改变node本来的内容
	// node = node.Next

	// 正确是这样写
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
