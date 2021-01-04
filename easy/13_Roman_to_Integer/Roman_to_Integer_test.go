package _3_Roman_to_Integer

import (
	"fmt"
	"testing"
)

type input struct {
	X string
	Answer
}
type Answer int

var output int

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = romanToInt(input.X)
	fmt.Printf("output: %v\n", output)
	fmt.Printf("answer: %v\n\n", input.Answer)
}

func TestReverse(t *testing.T) {
	Run(&input{
		X:      "III",
		Answer: 3,
	})

	Run(&input{
		X:      "IV",
		Answer: 4,
	})

	Run(&input{
		X:      "IX",
		Answer: 9,
	})

	Run(&input{
		X:      "LVIII",
		Answer: 58,
	})

	Run(&input{
		X:      "MCMXCIV",
		Answer: 1994,
	})
}
