package _0066_Plus_One

import (
	"fmt"
	"testing"
)

type input struct {
	X []int

	Answer []int
}

var output []int

func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	output = plusOne(input.X)
	fmt.Printf("output: %v\n", output)
	fmt.Printf("answer: %v\n\n", input.Answer)
}

func Test(t *testing.T) {
	Run(&input{
		X:      []int{1, 2, 3},
		Answer: []int{1, 2, 4},
	})
	Run(&input{
		X:      []int{4, 3, 2, 1},
		Answer: []int{4, 3, 2, 2},
	})
	Run(&input{
		X:      []int{0},
		Answer: []int{1},
	})
	Run(&input{
		X:      []int{9},
		Answer: []int{1, 0},
	})
}
