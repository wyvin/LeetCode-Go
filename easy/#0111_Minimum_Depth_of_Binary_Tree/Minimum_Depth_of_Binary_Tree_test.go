package _0111_Minimum_Depth_of_Binary_Tree

import (
	"fmt"
	"testing"
)

type input struct {
	X *TreeNode
	x string

	Answer int
}

var output int

func Run(input *input) {
	fmt.Printf("input: %v\n", input.x)
	output = minDepth(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
		x:      "[3,9,20,null,null,15,7]",
		Answer: 2,
	})
	Run(&input{
		X: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 2},
		},
		x:      "[1,2]",
		Answer: 2,
	})
	/*
			[2,null,3,null,4,null,5,null,6]
			         2
			      N     3
			          N   4
			            N   5
		                   N  6
	*/
	Run(&input{
		X: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 6},
					},
				},
			},
		},
		x:      "[2,null,3,null,4,null,5,null,6]",
		Answer: 5,
	})
}
