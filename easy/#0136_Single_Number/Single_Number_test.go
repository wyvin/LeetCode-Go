package _0136_Single_Number

import (
	"fmt"
	"testing"
)

type input struct {
	X []int

	Answer int
}

var output int

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = singleNumber(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      []int{2, 2, 1},
		Answer: 1,
	})
	Run(&input{
		X:      []int{4, 1, 2, 1, 2},
		Answer: 4,
	})
}
