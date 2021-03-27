package _0232_Implement_Queue_using_Stacks

import (
	"fmt"
	"testing"
)

type input struct {
	X []string
	Y [][]int

	Answer []interface{}
}

var myQueue MyQueue

func Run(input *input) {
	fmt.Printf("input: \n")
	fmt.Printf("%v \n", input.X)
	fmt.Printf("%v \n", input.Y)

	selector := func(a string, b []int) interface{} {
		switch a {
		case "MyQueue":
			myQueue = Constructor()
		case "push":
			for _, v := range b {
				myQueue.Push(v)
			}
		case "peek":
			return myQueue.Peek()
		case "pop":
			return myQueue.Pop()
		case "empty":
			return myQueue.Empty()
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
		X:      []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
		Y:      [][]int{{}, {1}, {2}, {}, {}, {}},
		Answer: []interface{}{nil, nil, nil, 1, 1, false},
	})
	Run(&input{
		X:      []string{"MyQueue", "push", "push", "push", "push", "pop", "push", "pop", "pop", "pop", "pop"},
		Y:      [][]int{{}, {1}, {2}, {3}, {4}, {}, {5}, {}, {}, {}, {}},
		Answer: []interface{}{nil, nil, nil, nil, nil, 1, nil, 2, 3, 4, 5},
	})
}
