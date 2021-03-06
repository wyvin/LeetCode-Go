package _0242_Valid_Anagram

import (
	"fmt"
	"testing"
)

type input struct {
	X string
	Y string

	Answer bool
}

func Test(t *testing.T) {
	Run(&input{
		X:      "anagram",
		Y:      "nagaram",
		Answer: true,
	})
}

func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	output := isAnagram(input.X, input.Y)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}
