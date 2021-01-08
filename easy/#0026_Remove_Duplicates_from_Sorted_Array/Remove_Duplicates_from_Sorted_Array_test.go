package _0026_Remove_Duplicates_from_Sorted_Array

import (
	"fmt"
	"testing"
)

type input struct {
	X []int

	Answer []int
}

var output int

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = removeDuplicates(input.X)
	fmt.Printf("output: %v\n", input.X[:output])
	fmt.Printf("answer: %v\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      []int{1,1,2},
		Answer: []int{1,2},
	})
	Run(&input{
		X:      []int{0,0,1,1,1,2,2,3,3,4},
		Answer: []int{0,1,2,3,4},
	})
}
