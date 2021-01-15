package _0028_Implement_strStr__

import (
	"fmt"
	"testing"
)

type input struct {
	X      string
	Needle string

	Answer int
}

var output int

func Run(input *input) {
	fmt.Printf("input: %q, needle: %q \n", input.X, input.Needle)
	output = strStr(input.X, input.Needle)
	fmt.Printf("output: %v\n", output)
	fmt.Printf("answer: %v\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      "hello",
		Needle: "ll",
		Answer: 2,
	})
	Run(&input{
		X:      "aaaaa",
		Needle: "bba",
		Answer: -1,
	})
	Run(&input{
		X:      "a",
		Needle: "",
		Answer: 0,
	})
	Run(&input{
		X:      "",
		Needle: "a",
		Answer: -1,
	})
	Run(&input{
		X:      "",
		Needle: "",
		Answer: 0,
	})
	Run(&input{
		X:      "a",
		Needle: "a",
		Answer: 0,
	})
	Run(&input{
		X:      "a",
		Needle: "aa",
		Answer: -1,
	})
}
