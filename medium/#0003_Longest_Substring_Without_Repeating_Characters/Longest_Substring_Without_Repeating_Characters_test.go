package _0003_Longest_Substring_Without_Repeating_Characters

import (
	"fmt"
	"strings"
	"testing"
)

type input struct {
	X string

	Answer int
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := lengthOfLongestSubstring2(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      "dvdf",
		Answer: 3,
	})
	Run(&input{
		X:      "abcabcbb",
		Answer: 3,
	})
}
