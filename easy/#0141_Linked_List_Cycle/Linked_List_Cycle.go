package _0141_Linked_List_Cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	方法一：
	可以使用map记录曾经访问过的节点
	当下一个节点在map里已经有里，则代表有环
*/
func hasCycle(head *ListNode) bool {
	cache := make(map[*ListNode]struct{})
	for head != nil {
		cache[head] = struct{}{}
		if head.Next == nil {
			return false
		}
		if _, ok := cache[head.Next]; ok {
			return true
		}
		head = head.Next
	}
	return false
}

/*
	方法二：
	节点数范围数最大10000，能跑完则代表有环
*/
func hasCycle2(head *ListNode) bool {
	x := 0
	for head != nil {
		x++
		if x >= 10000 {
			return true
		}
		head = head.Next
	}
	return false
}

/*
	方法三：
	由于Node.val的值最大为100000
	可以利用最大值来判断
*/
func hasCycle3(head *ListNode) bool {
	v := 100001
	for head != nil {
		head.Val = v
		if head.Next == nil {
			return false
		}
		if head.Next.Val == v {
			return true
		}
		head = head.Next
	}
	return false
}

/*
	方法四：
	快慢指针
	快指针能追上慢指针则代表有环
*/

func hasCycle4(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head.Next, head
	for fast != nil {
		if fast == slow {
			return true
		}
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
