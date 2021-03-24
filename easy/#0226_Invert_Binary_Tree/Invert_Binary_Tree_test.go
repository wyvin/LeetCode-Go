package _0226_Invert_Binary_Tree

import (
	"fmt"
	"testing"
)

type input struct {
	X *TreeNode

	Answer []int
}

var output []int

func iterator(a *TreeNode) *TreeNode {
	if a == nil {
		return nil
	}
	output = append(output, a.Val)
	iterator(a.Left)
	iterator(a.Right)
	return a
}

func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	tmp := invertTree(input.X)
	output = make([]int, 0)
	iterator(tmp)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   7,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 9},
			},
		},
		Answer: []int{4, 7, 9, 6, 2, 3, 1},
	})
}
