package _0160_Intersection_of_Two_Linked_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for headA != nil {
		B := headB
		for B != nil {
			if headA == B {
				return B
			}
			B = B.Next
		}
		headA = headA.Next
	}
	return nil
}

// 哈希表法
// 循环一条链表，记录在map里，空间复杂度O(m)或O(n)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	hash := make(map[*ListNode]struct{})
	for headA != nil {
		hash[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := hash[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 双指针法（浪漫法?)
// A尾链接到B头，B尾链接到A头，AB指针分表从AB链表开始跑，相遇则存在交点。
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ap := headA
	bp := headB
	// 结束条件：同时为nil
	for ap != nil || bp != nil {
		if ap == nil {
			ap = headB
		}
		if bp == nil {
			bp = headA
		}
		if ap == bp {
			return ap
		}
		ap = ap.Next
		bp = bp.Next
	}
	return nil
}
