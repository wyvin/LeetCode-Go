package _0237_Delete_Node_in_a_Linked_List

import (
	"fmt"
	"testing"
)

type input struct {
	X []int
	Y int

	Answer []int
}

func Run(input *input) {
	fmt.Printf("input: %v, %d\n", input.X, input.Y)
	// 根据[]int生成链表
	head := &ListNode{Val: input.X[0]}
	node := head
	for i := 1; i < len(input.X); i++ {
		node.Next = &ListNode{Val: input.X[i]}
		node = node.Next
	}
	// 删除节点
	node = head
	for node != nil {
		if node.Val == input.Y {
			deleteNode(node)
			break
		}
		node = node.Next
	}
	// 链表转[]int
	output := make([]int, 0)
	for head != nil {
		output = append(output, head.Val)
		head = head.Next
	}

	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      []int{4, 5, 1, 9},
		Y:      5,
		Answer: []int{4, 1, 9},
	})
}
