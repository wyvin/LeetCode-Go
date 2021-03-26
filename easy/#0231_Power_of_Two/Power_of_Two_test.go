package _0231_Power_of_Two

import (
	"fmt"
	"testing"
)

type input struct {
	X int

	Answer bool
}

func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	output := isPowerOfTwo2(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      1,
		Answer: true,
	})
	Run(&input{
		X:      16,
		Answer: true,
	})
	Run(&input{
		X:      218,
		Answer: false,
	})
}
