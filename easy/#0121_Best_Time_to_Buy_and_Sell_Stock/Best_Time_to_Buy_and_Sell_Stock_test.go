package _0121_Best_Time_to_Buy_and_Sell_Stock

import (
	"fmt"
	"math"
	"testing"
)

type input struct {
	X []int

	Answer int
}

var output int

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output = maxProfit(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      []int{7, 1, 5, 3, 6, 4},
		Answer: 5,
	})
}
