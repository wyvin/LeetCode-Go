package _0224_Basic_Calculator

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
	output := calculate(input.X)
	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      "1 + 1",
		Answer: 2,
	})
	Run(&input{
		X:      " 2-1 + 2 ",
		Answer: 3,
	})
	Run(&input{
		X:      "(1+(4+5+2)-3)+(6+8)",
		Answer: 23,
	})
}

// 官方
func calculate(s string) int {
	ans := 0
	// 用栈数据结构来记录括号前的运算符
	ops := []int{1}
	// sign 用于反转括号内的运算符： +( 时不影响； -( 时反转括号内的运算符
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}
