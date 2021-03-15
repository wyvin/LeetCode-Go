package _0219_Contains_Duplicate_II

import (
	"fmt"
	"testing"
)

type input struct {
	X []int
	K int

	Answer bool
}

func Run(input *input) {
	fmt.Printf("input: %v \n", input.X)
	output := containsNearbyDuplicate(input.X, input.K)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      []int{1, 2, 3, 1},
		K:      3,
		Answer: true,
	})
	Run(&input{
		X:      []int{1, 0, 1, 1},
		K:      1,
		Answer: true,
	})
	Run(&input{
		X:      []int{1, 2, 3, 1, 2, 3},
		K:      2,
		Answer: false,
	})
}
