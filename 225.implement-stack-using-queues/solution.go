package leetcode

type MyStack struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		[]int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	tmp := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return tmp
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.data[len(this.data)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
