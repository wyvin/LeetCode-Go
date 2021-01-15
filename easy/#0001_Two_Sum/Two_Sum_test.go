package _0001_Two_Sum

import (
	"fmt"
	"testing"
)

type input struct {
	Nums   []int
	Target int
	Answer
}
type Answer []int

var output []int

func Run(input *input) {
	fmt.Printf("input: %+v\n", input)
	output = twoSum1(input.Nums, input.Target)
	fmt.Printf("output: %v\n", output)
	fmt.Printf("answer: %v\n\n\n", input.Answer)
}

func TestTwoSum(t *testing.T) {
	Run(&input{
		Nums:   []int{2, 7, 11, 15},
		Target: 9,
		Answer: []int{0, 1},
	})
}
