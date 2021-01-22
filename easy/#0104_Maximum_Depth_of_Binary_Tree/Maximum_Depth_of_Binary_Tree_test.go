package _0104_Maximum_Depth_of_Binary_Tree

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
	output = maxDepth(input.X)
	fmt.Printf("output: %d\n", output)
	fmt.Printf("answer: %d\n\n", input.Answer)
	if fmt.Sprintf("%v", input.Answer) != fmt.Sprintf("%v", output) {
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
		Answer: 3,
	})
}
