package _0083_Remove_Duplicates_from_Sorted_List

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type input struct {
	X *ListNode

	Answer string
}

var output *ListNode

func Run(input *input) {
	var in []string
	ins := input.X
	for ins != nil {
		in = append(in, strconv.Itoa(ins.Val))
		ins = ins.Next
	}
	fmt.Printf("input: %s\n", strings.Join(in, "->"))
	output = deleteDuplicates(input.X)
	var out []string
	for output != nil {
		out = append(out, strconv.Itoa(output.Val))
		output = output.Next
	}
	fmt.Printf("output: %s\n", strings.Join(out, "->"))
	fmt.Printf("answer: %s\n\n", input.Answer)
}

func Test(t *testing.T) {
	Run(&input{
		X: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			},
		},
		Answer: "1->2",
	})
	Run(&input{
		X: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{Val: 3},
					},
				},
			},
		},
		Answer: "1->2->3",
	})
	Run(&input{
		X:      nil,
		Answer: "",
	})
}
