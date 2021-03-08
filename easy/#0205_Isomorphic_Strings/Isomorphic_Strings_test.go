package _0205_Isomorphic_Strings

import (
	"fmt"
	"testing"
)

type input struct {
	X string
	Y string

	Answer bool
}

func Run(input *input) {
	fmt.Printf("input: %v %v\n", input.X, input.Y)
	output := isIsomorphic(input.X, input.Y)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      "egg",
		Y:      "add",
		Answer: true,
	})
	Run(&input{
		X:      "foo",
		Y:      "bar",
		Answer: false,
	})
	Run(&input{
		X:      "paper",
		Y:      "title",
		Answer: true,
	})
}
