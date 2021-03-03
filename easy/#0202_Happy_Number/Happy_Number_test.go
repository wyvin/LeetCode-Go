package _0202_Happy_Number

import (
	"fmt"
	"testing"
)

type input struct {
	X int

	Answer bool
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := isHappy(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      19,
		Answer: true,
	})
	Run(&input{
		X:      2,
		Answer: false,
	})
	Run(&input{
		X:      5,
		Answer: false,
	})
}
