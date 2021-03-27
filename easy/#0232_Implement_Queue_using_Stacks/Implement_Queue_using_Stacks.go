package _0232_Implement_Queue_using_Stacks

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
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	tmp := this.Peek()
	this.OutStack = this.OutStack[:len(this.OutStack)-1]
	return tmp
}

// 检查outStack是否空，空的话就把inStack倒进outStack
/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.OutStack) == 0 {
		for len(this.InStack) > 0 {
			this.OutStack = append(this.OutStack, this.InStack[len(this.InStack)-1])
			this.InStack = this.InStack[:len(this.InStack)-1]
		}
	}
	return this.OutStack[len(this.OutStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.OutStack) == 0 && len(this.InStack) == 0
}
