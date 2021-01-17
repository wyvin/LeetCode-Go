package _0069_Sqrt_X

import (
	"fmt"
	"testing"
)

type input struct {
	X int

	Answer int
}

var output int

func Run(input *input) {
	fmt.Printf("input: a=%d\n", input.X)
	output = mySqrt(input.X)
	fmt.Printf("output: %d\n", output)
	fmt.Printf("answer: %d\n\n", input.Answer)
}

func Test(t *testing.T) {
	Run(&input{
		X:      4,
		Answer: 2,
	})
	Run(&input{
		X:      8,
		Answer: 2,
	})
}
