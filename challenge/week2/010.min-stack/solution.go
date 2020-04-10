package leetcode

type MinStack struct {
	stack []int
	min   []int
	size  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
		size:  0,
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if this.size > 0 {
		this.min = append(this.min, min(this.min[this.size-1], x))
	} else {
		this.min = append(this.min, x)
	}
	this.size++
}

func (this *MinStack) Pop() {
	if this.size > 0 {
		this.stack = this.stack[:this.size-1]
		this.min = this.min[:this.size-1]
		this.size--
	}
}

func (this *MinStack) Top() int {
	return this.stack[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.size-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
