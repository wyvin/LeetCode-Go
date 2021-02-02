package _0125_Valid_Palindrome

import (
	"fmt"
	"testing"
)

type input struct {
	X string

	Answer bool
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := isPalindrome(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      "A man, a plan, a canal: Panama",
		Answer: true,
	})
	Run(&input{
		X:      "race a car",
		Answer: false,
	})
	Run(&input{
		X:      "0P",
		Answer: false,
	})
}
