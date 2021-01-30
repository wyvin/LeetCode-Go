package _0119_Pascal_s_Triangle_II

import (
	"fmt"
	"testing"
)

type input struct {
	X int

	Answer []int
}

var output []int

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = getRow(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      3,
		Answer: []int{1, 3, 3, 1},
	})
}
