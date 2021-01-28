package _0112_Path_Sum

import (
	"fmt"
	"testing"
)

type input struct {
	X      *TreeNode
	Target int
	x      string

	Answer bool
}

var output bool

func Run(input *input) {
	fmt.Printf("input: %v\n", input.x)
	output = hasPathSum(input.X, input.Target)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
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
		Target: 5,
		x:      "[1,2,3]",
		Answer: false,
	})
	Run(&input{
		X: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 2},
		},
		Target: 0,
		x:      "[1,2]",
		Answer: false,
	})
}
