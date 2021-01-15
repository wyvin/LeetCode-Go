package _0007_Reverse_Integer

import (
	"fmt"
	"testing"
)

type input struct {
	X int
	Answer
}
type Answer int

var output int

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = reverse(input.X)
	fmt.Printf("output: %v\n", output)
	fmt.Printf("answer: %v\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      123,
		Answer: 321,
	})

	Run(&input{
		X:      -123,
		Answer: -321,
	})

	Run(&input{
		X:      -120,
		Answer: -21,
	})

	Run(&input{
		X:      1563847412,
		Answer: 0,
	})
}
