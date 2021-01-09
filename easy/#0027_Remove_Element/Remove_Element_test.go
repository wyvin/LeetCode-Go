package _0027_Remove_Element

import (
	"fmt"
	"testing"
)

type input struct {
	X []int
	Val int

	Answer []int
}

var output int

func Run(input *input) {
	fmt.Printf("input: %v, val:%d \n", input.X,input.Val)
	output = removeElement(input.X, input.Val)
	fmt.Printf("output: %v\n", input.X[:output])
	fmt.Printf("answer: %v\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      []int{3,2,2,3},
		Val: 3,
		Answer: []int{2,2},
	})
	Run(&input{
		X:      []int{0,1,2,2,3,0,4,2},
		Val: 2,
		Answer: []int{0,1,3,0,4},
	})
	Run(&input{
		X:      []int{1},
		Val: 1,
		Answer: []int{},
	})
	Run(&input{
		X:      []int{},
		Val: 0,
		Answer: []int{},
	})
}
