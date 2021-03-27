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

var myStack MyQueue

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
		case "peek":
			return myStack.Peek()
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
		X:      []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
		Y:      [][]int{{}, {1}, {2}, {}, {}, {}},
		Answer: []interface{}{nil, nil, nil, 1, 1, false},
	})
}

type MyQueue struct {
	InStack  []int
	OutStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.InStack = append(this.InStack, x)
	for ol := len(this.OutStack) - 1; ol >= 0; ol-- {
		this.InStack = append(this.InStack, this.OutStack[ol])
	}
	this.OutStack = this.InStack
	this.InStack = []int{}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ol := len(this.OutStack)
	tmp := this.OutStack[ol-1]
	this.OutStack = this.OutStack[:ol-1]
	return tmp
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.OutStack[len(this.OutStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.OutStack) == 0
}
