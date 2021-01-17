package _0070_Climbing_Stairs

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
	fmt.Printf("input: %d\n", input.X)
	output = climbStairs(input.X)
	fmt.Printf("output: %d\n", output)
	fmt.Printf("answer: %d\n\n", input.Answer)
}

func Test(t *testing.T) {
	Run(&input{
		X:      2,
		Answer: 2,
	})
	Run(&input{
		X:      3,
		Answer: 3,
	})
}
