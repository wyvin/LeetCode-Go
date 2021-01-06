package _0020_Valid_Parentheses

import (
	"fmt"
	"testing"
)

type input struct {
	X string

	Answer bool
}

var output bool

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = isValid(input.X)
	fmt.Printf("output: %v\n", output)
	fmt.Printf("answer: %v\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      "()",
		Answer: true,
	})
	Run(&input{
		X:      "()[]{}",
		Answer: true,
	})
	Run(&input{
		X:      "(]",
		Answer: false,
	})
	Run(&input{
		X:      "([)]",
		Answer: false,
	})
	Run(&input{
		X:      "{[]}",
		Answer: true,
	})
	Run(&input{
		X:      "{",
		Answer: false,
	})
	Run(&input{
		X:      "{{",
		Answer: false,
	})
	Run(&input{
		X:      ")(",
		Answer: false,
	})
}
