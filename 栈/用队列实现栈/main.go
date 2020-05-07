type MyStack struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		stack: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.stack) > 0 {
		lastIndex := len(this.stack) - 1
		ret := this.stack[lastIndex]
		this.stack = this.stack[:lastIndex]
		return ret
	}
	return 0
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.stack) > 0 {
		lastIndex := len(this.stack) - 1
		return this.stack[lastIndex]
	}
	return 0
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */