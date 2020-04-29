type CQueue struct {
	in  []int
	out []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.in) == 0 && len(this.out) == 0 {
		return -1
	}
	// 如果out为空则把 in中元素压入out
	if len(this.out) == 0 {
		for i := len(this.in) - 1; i >= 0; i-- {
			this.out = append(this.out, this.in[i])
		}
		this.in = []int{}
	}
	// 取出out 最后一个元素
	popIndex := len(this.out) - 1
	popValue := this.out[popIndex]
	this.out = this.out[:popIndex]
	return popValue
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */