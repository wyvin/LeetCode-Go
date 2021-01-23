package _0107_Binary_Tree_Level_Order_Traversal_II

import (
	"fmt"
	"testing"
)

type input struct {
	X *TreeNode
	x string

	Answer [][]int
}

var output [][]int

func Run(input *input) {
	fmt.Printf("input: %v\n", input.x)
	output = levelOrderBottom1(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
		x:      "[3,9,20,null,null,15,7]",
		Answer: [][]int{{15, 7}, {9, 20}, {3}},
	})
}
