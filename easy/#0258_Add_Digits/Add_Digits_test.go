package _0258_Add_Digits

import (
	"fmt"
	"testing"
)

type input struct {
	X int

	Answer int
}

func Test(t *testing.T) {
	Run(&input{
		X:      38,
		Answer: 2,
	})
}
func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	output := addDigits2(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}
