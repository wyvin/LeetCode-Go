package _0168_Excel_Sheet_Column_Title

import (
	"fmt"
	"testing"
)

type input struct {
	X int

	Answer string
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := convertToTitle(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      1,
		Answer: "A",
	})
	Run(&input{
		X:      28,
		Answer: "AB",
	})
	Run(&input{
		X:      701,
		Answer: "ZY",
	})
	Run(&input{
		X:      52,
		Answer: "AZ",
	})
}
