package _0053_Maximum_Subarray

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
	fmt.Printf("input: %v \n", input.X)
	output = maxSubArray(input.X)
	fmt.Printf("output: %d\n", output)
	fmt.Printf("answer: %d\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		Answer: 6,
	})
	Run(&input{
		X:      []int{1},
		Answer: 1,
	})
	Run(&input{
		X:      []int{-1},
		Answer: -1,
	})
}
