package __Palindrome_Number

import (
	"fmt"
	"testing"
)

type input struct {
	X int
	Answer
}
type Answer bool

var output bool

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = isPalindrome(input.X)
	fmt.Printf("output: %v\n", output)
	fmt.Printf("answer: %v\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      121,
		Answer: true,
	})

	Run(&input{
		X:      -121,
		Answer: false,
	})

	Run(&input{
		X:      10,
		Answer: false,
	})

	Run(&input{
		X:      1221,
		Answer: true,
	})

	Run(&input{
		X:      0,
		Answer: true,
	})
}
