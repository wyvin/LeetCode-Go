package _0217_Contains_Duplicate

import (
	"fmt"
	"testing"
)

type input struct {
	X []int

	Answer bool
}

func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	output := containsDuplicate2(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      []int{1, 2, 3, 1},
		Answer: true,
	})
	Run(&input{
		X:      []int{1, 2, 3, 4},
		Answer: false,
	})
	Run(&input{
		X:      []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		Answer: true,
	})
}
