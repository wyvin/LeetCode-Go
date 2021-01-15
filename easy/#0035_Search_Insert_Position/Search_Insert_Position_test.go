package _0035_Search_Insert_Position

import (
	"fmt"
	"testing"
)

type input struct {
	X      []int
	Target int

	Answer int
}

var output int

func Run(input *input) {
	fmt.Printf("input: %v, target: %d \n", input.X, input.Target)
	output = searchInsert2(input.X, input.Target)
	fmt.Printf("output: %d\n", output)
	fmt.Printf("answer: %d\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      []int{1, 3, 5, 6},
		Target: 5,
		Answer: 2,
	})
	Run(&input{
		X:      []int{1, 3, 5, 6},
		Target: 2,
		Answer: 1,
	})
	Run(&input{
		X:      []int{1, 3, 5, 6},
		Target: 7,
		Answer: 4,
	})
	Run(&input{
		X:      []int{1, 3, 5, 6},
		Target: 0,
		Answer: 0,
	})
}
