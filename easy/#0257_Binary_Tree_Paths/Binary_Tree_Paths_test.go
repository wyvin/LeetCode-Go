package _0257_Binary_Tree_Paths

import (
	"fmt"
	"testing"
)

type input struct {
	X *TreeNode

	Answer []string
}

func Test(t *testing.T) {
	Run(&input{
		X: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Answer: []string{"1->2->5", "1->3"},
	})
}

func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	output := binaryTreePaths(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}
