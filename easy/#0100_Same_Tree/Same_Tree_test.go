package _0100_Same_Tree

import (
	"fmt"
	"testing"
)

type input struct {
	X *TreeNode
	x string
	Y *TreeNode
	y string

	Answer bool
}

var output bool

func Run(input *input) {
	fmt.Printf("input: %v, %v\n", input.x, input.y)
	output = isSameTree(input.X, input.Y)
	fmt.Printf("output: %t\n", output)
	fmt.Printf("answer: %t\n\n", input.Answer)
	if fmt.Sprintf("%v", input.Answer) != fmt.Sprintf("%v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		},
		x: "[1,2,3]",
		Y: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		},
		y:      "[1,2,3]",
		Answer: true,
	})
	Run(&input{
		X: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 2},
		},
		x: "[1,2]",
		Y: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		y:      "[1,null,2]",
		Answer: false,
	})
	Run(&input{
		X: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 1},
		},
		x: "[1,2,1]",
		Y: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		y:      "[1,1,2]",
		Answer: false,
	})
}
