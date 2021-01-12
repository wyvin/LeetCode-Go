package _0038_Count_and_Say

import (
	"fmt"
	"testing"
)

type input struct {
	X int

	Answer string
}

var output string

func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	output = countAndSay2(input.X)
	fmt.Printf("output: %q\n", output)
	fmt.Printf("answer: %q\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      1,
		Answer: "1",
	})
	Run(&input{
		X:      2,
		Answer: "11",
	})
	Run(&input{
		X:      3,
		Answer: "21",
	})
	Run(&input{
		X:      4,
		Answer: "1211",
	})
	Run(&input{
		X:      5,
		Answer: "111221",
	})
}
