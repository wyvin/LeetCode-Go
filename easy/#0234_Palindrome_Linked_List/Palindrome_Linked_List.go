package _0234_Palindrome_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：把链上的数取出来放在数组中，再用双指针判断是否回文
// 时间复杂度O(3/2n) 空间复杂度O(n)

// 方法二：
// 1.找到前半部分链表的尾节点。（快慢指针，同时开始，慢走一步快走两步）
// 2.反转后半部分链表。（改变箭头方向）
//   1->2->2->1  =>  1->2->2<-1
//   1->2->3->2->1  =>  1->2->3->2<-1
//                      ^           ^
//                      |           |
//                    头指针     反转后的指针
// 3.判断是否回文。（头指针和慢指针同时开始循环）
// 4.恢复链表。（题目不需要，但平时工作上注意恢复）

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	// 1
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 2
	var pre *ListNode
	cur := slow.Next
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	// 3
	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}
