package _0228_Summary_Ranges

import (
	"fmt"
	"testing"
)

type input struct {
	X []int

	Answer []string
}

func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	output := summaryRanges2(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      []int{0, 1, 2, 4, 5, 7},
		Answer: []string{"0->2", "4->5", "7"},
	})
	Run(&input{
		X:      []int{0, 2, 3, 4, 6, 8, 9},
		Answer: []string{"0", "2->4", "6", "8->9"},
	})
	Run(&input{
		X:      []int{},
		Answer: []string{},
	})
	Run(&input{
		X:      []int{-1},
		Answer: []string{"-1"},
	})
	Run(&input{
		X:      []int{0},
		Answer: []string{"0"},
	})
}
