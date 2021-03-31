package _0234_Palindrome_Linked_List

import (
	"fmt"
	"testing"
)

type input struct {
	X *ListNode
	x []int

	Answer bool
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.x)
	output := isPalindrome(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
		x:      []int{1, 2},
		Answer: false,
	})
	Run(&input{
		X: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
		x:      []int{1, 2, 2, 1},
		Answer: true,
	})
	Run(&input{
		X: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		x:      []int{1, 2, 3, 4, 5},
		Answer: false,
	})
}
