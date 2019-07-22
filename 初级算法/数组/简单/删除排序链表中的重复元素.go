package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	currNode := head
	for currNode != nil && currNode.Next != nil {
		if currNode.Val == currNode.Next.Val {
			currNode.Next = currNode.Next.Next
		} else {
			currNode = currNode.Next
		}
	}
	return head
}

func main() {
	
}
