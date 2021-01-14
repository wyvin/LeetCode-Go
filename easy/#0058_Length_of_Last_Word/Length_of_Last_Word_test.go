package _0058_Length_of_Last_Word

import (
	"fmt"
	"testing"
)

type input struct {
	X string

	Answer int
}

var output int

func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	output = lengthOfLastWord(input.X)
	fmt.Printf("output: %d\n", output)
	fmt.Printf("answer: %d\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      "Hello World",
		Answer: 5,
	})
	Run(&input{
		X:      "  ",
		Answer: 0,
	})
	Run(&input{
		X:      " Hello World ",
		Answer: 5,
	})
}
