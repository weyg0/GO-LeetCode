package offers

import "math"

type MinStack struct {
	stack []int
	min   int
	size  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min: math.MaxInt}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.size++
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	res := this.stack[this.size-1]
	this.stack = this.stack[:this.size-1]
	this.size--
	if res == this.min {
		this.min = math.MaxInt
		for _, val := range this.stack {
			if val < this.min {
				this.min = val
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[this.size-1]
}

func (this *MinStack) Min() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
