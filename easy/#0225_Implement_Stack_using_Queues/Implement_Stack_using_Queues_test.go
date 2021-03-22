package _0225_Implement_Stack_using_Queues

import (
	"fmt"
	"testing"
)

type input struct {
	X []string
	Y [][]int

	Answer []interface{}
}

var myStack MyStack

func Run(input *input) {
	fmt.Printf("input: \n")
	fmt.Printf("%v \n", input.X)
	fmt.Printf("%v \n", input.Y)

	selector := func(a string, b []int) interface{} {
		switch a {
		case "MyStack":
			myStack = Constructor()
		case "push":
			for _, v := range b {
				myStack.Push(v)
			}
		case "top":
			return myStack.Top()
		case "pop":
			return myStack.Pop()
		case "empty":
			return myStack.Empty()
		}
		return nil
	}

	output := make([]interface{}, 0)
	for i, _ := range input.X {
		output = append(output, selector(input.X[i], input.Y[i]))
	}

	fmt.Printf("output: %+v\n", output)
	fmt.Printf("answer: %+v\n\n", input.Answer)
	if fmt.Sprintf("%+v", input.Answer) != fmt.Sprintf("%+v", output) {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run(&input{
		X:      []string{"MyStack", "push", "push", "top", "pop", "empty"},
		Y:      [][]int{{}, {1}, {2}, {}, {}, {}},
		Answer: []interface{}{nil, nil, nil, 2, 2, false},
	})
	Run(&input{
		X:      []string{"MyStack", "push", "pop", "empty"},
		Y:      [][]int{{}, {1}, {}, {}},
		Answer: []interface{}{nil, nil, 1, true},
	})
}
