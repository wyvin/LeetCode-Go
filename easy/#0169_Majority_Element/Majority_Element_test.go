package _0169_Majority_Element

import (
	"fmt"
	"testing"
)

type input struct {
	X []int

	Answer int
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := majorityElement(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      []int{3, 2, 3},
		Answer: 3,
	})
	Run(&input{
		X:      []int{2, 2, 1, 1, 1, 2, 2},
		Answer: 2,
	})
}
