type MyQueue struct {
    stackIn []int
    stackOut []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{
        stackIn: []int{},
        stackOut: []int{},
    }
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.stackIn = append(this.stackIn, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    // 出栈有数据这直接返回栈顶元素
    if len(this.stackOut) == 0 && len(this.stackIn) == 0 {
        return 0
    }
    if len(this.stackOut) > 0 {
        ret := this.stackOut[0]
        this.stackOut = this.stackOut[1:]
        return ret
    // 否则 将入栈中的所有元素弹出到出栈中
    } else if len(this.stackIn) > 0 {
        for _, i := range this.stackIn{
            this.stackOut = append(this.stackOut, i)
        }
        this.stackIn = []int{}
        ret := this.stackOut[0]
        this.stackOut = this.stackOut[1:]
        return ret
    } 
    return 0
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if len(this.stackOut) == 0 && len(this.stackIn) == 0 {
        return 0
    }
    if len(this.stackOut) > 0 {
        ret := this.stackOut[0]
        return ret
    // 否则 将入栈中的所有元素弹出到出栈中
    } else if len(this.stackIn) > 0 {
        for _, i := range this.stackIn{
            this.stackOut = append(this.stackOut, i)
        }
        this.stackIn = []int{}
        ret := this.stackOut[0]
        return ret
    } 
    return 0
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.stackOut) == 0 && len(this.stackIn) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */