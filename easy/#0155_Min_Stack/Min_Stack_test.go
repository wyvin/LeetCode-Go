package _0155_Min_Stack

import (
	"fmt"
	"testing"
)

func Run() {
	fmt.Println("MinStack minStack = new MinStack();")
	minStack := Constructor()
	fmt.Println("minStack.push(-2);")
	minStack.Push(-2)
	fmt.Println("minStack.push(0);")
	minStack.Push(0)
	fmt.Println("minStack.push(-3);")
	minStack.Push(-3)
	fmt.Printf("minStack.getMin();   --> 返回 %d.\n", minStack.GetMin())
	if minStack.GetMin() != -3 {
		panic("解答错误")
	}

	fmt.Println("minStack.pop();")
	minStack.Pop()
	fmt.Printf("minStack.top();      --> 返回 %d.\n", minStack.Top())
	if minStack.Top() != 0 {
		panic("解答错误")
	}
	fmt.Printf("minStack.getMin();   --> 返回 %d.\n", minStack.GetMin())
	if minStack.GetMin() != -2 {
		panic("解答错误")
	}
}

func Test(t *testing.T) {
	Run()
}
