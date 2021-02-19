package _0171_Excel_Sheet_Column_Number

import (
	"fmt"
	"testing"
)

type input struct {
	X string

	Answer int
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := titleToNumber(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      "A",
		Answer: 1,
	})
	Run(&input{
		X:      "AB",
		Answer: 28,
	})
	Run(&input{
		X:      "ZY",
		Answer: 701,
	})
	Run(&input{
		X:      "AZ",
		Answer: 52,
	})
	Run(&input{
		X:      "AAA",
		Answer: 703,
	})
}
