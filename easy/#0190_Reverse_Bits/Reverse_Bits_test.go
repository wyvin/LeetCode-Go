package _0190_Reverse_Bits

import (
	"fmt"
	"testing"
)

type input struct {
	X uint32

	Answer uint32
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := reverseBits2(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      43261596,
		Answer: 964176192,
	})
	Run(&input{
		X:      4294967293,
		Answer: 3221225471,
	})
}
