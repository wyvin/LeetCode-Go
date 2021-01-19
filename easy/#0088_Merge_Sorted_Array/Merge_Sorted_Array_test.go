package _0088_Merge_Sorted_Array

import (
	"fmt"
	"testing"
)

type input struct {
	X []int
	M int
	Y []int
	N int

	Answer []int
}

//var output []int

func Run(input *input) {
	fmt.Printf("input: nums1=%v, m=%d, nums2=%v, n=%d\n", input.X, input.M, input.Y, input.N)
	merge(input.X, input.M, input.Y, input.N)
	fmt.Printf("output: %d\n", input.X)
	fmt.Printf("answer: %d\n\n", input.Answer)
	if fmt.Sprintf("%v", input.Answer) != fmt.Sprintf("%v", input.X) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      []int{1, 2, 3, 0, 0, 0},
		M:      3,
		Y:      []int{2, 5, 6},
		N:      3,
		Answer: []int{1, 2, 2, 3, 5, 6},
	})
	Run(&input{
		X:      []int{1},
		M:      1,
		Y:      []int{},
		N:      0,
		Answer: []int{1},
	})
}
