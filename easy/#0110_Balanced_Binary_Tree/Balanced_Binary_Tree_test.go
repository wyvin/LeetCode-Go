package _0110_Balanced_Binary_Tree

import (
	"fmt"
	"testing"
)

type input struct {
	X *TreeNode
	x string

	Answer bool
}

var output bool

func Run(input *input) {
	fmt.Printf("input: %v\n", input.x)
	output = isBalanced(input.X)
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
		Answer: true,
	})
	Run(&input{
		X: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 2},
		},
		x:      "[1,2,2,3,3,null,null,4,4]",
		Answer: false,
	})
	/*
		[1,2,2,3,null,null,3,4,null,null,4]
		         1
		      2     2
		    3  N  N  3
		  4 N       N 4
	*/
	Run(&input{
		X: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 4},
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
		},
		x:      "[1,2,2,3,null,null,3,4,null,null,4]",
		Answer: false,
	})
}
