package _0155_Min_Stack

import (
	"fmt"
	"math"
	"testing"
)

type MinStack struct {
	items        []int
	minIndexList []int
	min          int
	length       int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		items:        []int{},
		length:       0,
		minIndexList: []int{},
		min:          math.MaxInt64,
	}
}

func (this *MinStack) Push(x int) {
	this.items = append(this.items, x)
	if x < this.min {
		this.min = x
		this.minIndexList = append(this.minIndexList, this.length)
	}
	this.length++
}

func (this *MinStack) Pop() {
	if this.length <= 0 {
		return
	}
	this.length--
	this.items = this.items[:this.length]
	if this.length == this.minIndexList[len(this.minIndexList)-1] {
		this.minIndexList = this.minIndexList[:len(this.minIndexList)-1]
		if len(this.minIndexList) == 0 {
			this.min = math.MaxInt64
		} else {
			this.min = this.items[this.minIndexList[len(this.minIndexList)-1]]
		}
	}
}

func (this *MinStack) Top() int {
	if this.length > 0 {
		return this.items[this.length-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.length > 0 {
		return this.min
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

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
