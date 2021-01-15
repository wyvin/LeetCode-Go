package _0014_Longest_Common_Prefix

import (
	"fmt"
	"testing"
)

type input struct {
	X []string

	Answer string
}

var output string

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = longestCommonPrefix(input.X)
	fmt.Printf("output: %v\n", output)
	fmt.Printf("answer: %v\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      []string{"flower", "flow", "flight"},
		Answer: "fl",
	})

	Run(&input{
		X:      []string{"dog", "racecar", "car"},
		Answer: "",
	})

	Run(&input{
		X:      []string{"ab", "a"},
		Answer: "a",
	})
}
