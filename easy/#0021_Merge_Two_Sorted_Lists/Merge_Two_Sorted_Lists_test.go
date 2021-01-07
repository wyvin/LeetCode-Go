package _0021_Merge_Two_Sorted_Lists

import (
	"fmt"
	"testing"
)

type input struct {
	L1 *ListNode
	L2 *ListNode

	Answer []int
}

var output *ListNode

func Run(input *input) {
	l1 := make([]int, 0)
	L1 := new(ListNode)
	if input.L1 != nil {
		*L1 = *input.L1
	}
	for L1 != nil {
		l1 = append(l1, L1.Val)
		L1 = L1.Next
	}
	l2 := make([]int, 0)
	L2 := new(ListNode)
	if input.L2 != nil {
		*L2 = *input.L2
	}
	for L2 != nil {
		l2 = append(l2, L2.Val)
		L2 = L2.Next
	}
	fmt.Printf("input: l1=%v, l2=%v\n", l1, l2)
	output = mergeTwoLists(input.L1, input.L2)

	res := make([]int, 0)
	for output != nil {
		res = append(res, output.Val)
		output = output.Next
	}

	fmt.Printf("output: %v\n", res)
	fmt.Printf("answer: %v\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		L1:      &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		L2:      &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		Answer: []int{1,1,2,3,4,4},
	})

	Run(&input{
		L1:  nil,
		L2:     nil,
		Answer: []int{},
	})
	Run(&input{
		L1:     nil,
		L2:      &ListNode{
			Val:  0,
			Next: nil,
		},
		Answer: []int{0},
	})
}
