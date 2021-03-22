package _0225_Implement_Stack_using_Queues

// 类似汉诺塔的思路
type MyStack struct {
	ListIn  []int
	ListOut []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		ListIn:  []int{},
		ListOut: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.ListIn = append(this.ListIn, x)
	if len(this.ListOut) != 0 {
		this.ListIn = append(this.ListIn, this.ListOut...)
		this.ListOut = []int{}
	}
	this.ListIn, this.ListOut = this.ListOut, this.ListIn
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ans := this.ListOut[0]
	this.ListOut = this.ListOut[1:]
	return ans
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.ListOut[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.ListIn) == 0 && len(this.ListOut) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
