package _0141_Linked_List_Cycle

import (
	"fmt"
	"testing"
)

type input struct {
	X *ListNode
	x string

	Answer bool
}

var output bool

func Run(input *input) {
	fmt.Printf("input: %v\n", input.x)
	output = hasCycle(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	a1 := &ListNode{Val: 3}
	a2 := &ListNode{Val: 2}
	a3 := &ListNode{Val: 0}
	a4 := &ListNode{Val: -4}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a2
	Run(&input{
		X:      a1,
		x:      "[3,2,0,-4]",
		Answer: true,
	})
}
