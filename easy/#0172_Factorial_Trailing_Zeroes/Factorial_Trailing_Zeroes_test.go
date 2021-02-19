package _0172_Factorial_Trailing_Zeroes

import (
	"fmt"
	"testing"
)

type input struct {
	X int

	Answer int
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := trailingZeroes(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      3,
		Answer: 0,
	})
	Run(&input{
		X:      5,
		Answer: 1,
	})
}
