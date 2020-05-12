type MyQueue struct {
	inStack  []int
	outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 && len(this.inStack) == 0 {
		return 0
	}
	if len(this.outStack) == 0 {
		this.outStack = append(this.outStack, this.inStack...)
		this.inStack = []int{}
	}
	ret := this.outStack[0]
	this.outStack = this.outStack[1:]
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 && len(this.inStack) == 0 {
		return 0
	}
	if len(this.outStack) == 0 {
		this.outStack = append(this.outStack, this.inStack...)
		this.inStack = []int{}
	}
	return this.outStack[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */