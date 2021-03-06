package _0204_Count_Primes

import (
	"fmt"
	"testing"
)

type input struct {
	X int

	Answer int
}

func Run(input *input) {
	fmt.Printf("input: %v\n", input.X)
	output := countPrimes(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      10,
		Answer: 4,
	})
	Run(&input{
		X:      0,
		Answer: 0,
	})
	Run(&input{
		X:      1,
		Answer: 0,
	})
}
