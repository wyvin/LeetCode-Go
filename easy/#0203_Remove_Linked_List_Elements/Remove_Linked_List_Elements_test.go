package _0203_Remove_Linked_List_Elements

import (
	"fmt"
	"testing"
)

type input struct {
	X   *ListNode
	Val int

	Answer []int
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := removeElements2(input.X, input.Val)
	op := make([]int, 0)
	for output != nil {
		op = append(op, output.Val)
		output = output.Next
	}
	fmt.Printf("output: %+v\n", op)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", op) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: &ListNode{Val: 6},
							},
						},
					},
				},
			},
		},
		Val:    6,
		Answer: []int{1, 2, 3, 4, 5},
	})
}
