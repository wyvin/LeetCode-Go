package _0191_Number_of_1_Bits

import (
	"fmt"
	"testing"
)

type input struct {
	X uint32

	Answer int
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := hammingWeight(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      0b00000000000000000000000000001011,
		Answer: 3,
	})
	Run(&input{
		X:      0b00000000000000000000000010000000,
		Answer: 1,
	})
	Run(&input{
		X:      0b11111111111111111111111111111101,
		Answer: 31,
	})
}
