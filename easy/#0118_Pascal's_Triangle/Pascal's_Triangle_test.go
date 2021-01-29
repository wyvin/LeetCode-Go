package _0118_Pascal_s_Triangle

import (
	"fmt"
	"testing"
)

type input struct {
	X int

	Answer [][]int
}

var output [][]int

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = generate(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X: 5,
		Answer: [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		},
	})
}
