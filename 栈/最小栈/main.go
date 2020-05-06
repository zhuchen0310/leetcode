type MinStack struct {
	stack  []int
	helper []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:  []int{},
		helper: []int{},
	}
}

func (this *MinStack) Push(x int) {
	if len(this.helper) == 0 || x < this.helper[len(this.helper)-1] {
		this.helper = append(this.helper, x)
	} else {
		this.helper = append(this.helper, this.helper[len(this.helper)-1])
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		lastIndex := len(this.stack) - 1
		this.stack = this.stack[:lastIndex]
		this.helper = this.helper[:lastIndex]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		lastIndex := len(this.stack) - 1
		return this.stack[lastIndex]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.stack) > 0 {
		lastIndex := len(this.stack) - 1
		return this.helper[lastIndex]
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */