package _0067_Add_Binary

import (
	"fmt"
	"testing"
)

type input struct {
	X string
	Y string

	Answer string
}

var output string

func Run(input *input) {
	fmt.Printf("input: a=%q  b=%q\n", input.X, input.Y)
	output = addBinary2(input.X, input.Y)
	fmt.Printf("output: %q\n", output)
	fmt.Printf("answer: %q\n\n", input.Answer)
}

func Test(t *testing.T) {
	Run(&input{
		X:      "11",
		Y:      "1",
		Answer: "100",
	})
	Run(&input{
		X:      "1010",
		Y:      "1011",
		Answer: "10101",
	})
}
