package _0101_Symmetric_Tree

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
	output = isSymmetric(input.X)
	fmt.Printf("output: %t\n", output)
	fmt.Printf("answer: %t\n\n", input.Answer)
	if fmt.Sprintf("%v", input.Answer) != fmt.Sprintf("%v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 3},
			},
		},
		x:      "[1,2,2,3,4,4,3]",
		Answer: true,
	})
	Run(&input{
		X: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
		},
		x:      "[1,2,2,null,3,null,3]",
		Answer: false,
	})
}
