package leetcode

import "math"

/*type MinStack struct {
	stack []int
	min   int
	size  int
}

func Constructor() MinStack {
	return MinStack{min: math.MaxInt}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.size++
	if val < this.min {
		this.min = val
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

func (this *MinStack) GetMin() int {
	return this.min
}*/

type MinStack struct {
	stack, minStack []int
	size            int
}

func Constructor() MinStack {
	return MinStack{minStack: []int{math.MaxInt}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if val < this.GetMin() {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.GetMin())
	}
	this.size++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.size-1]
	this.minStack = this.minStack[:this.size]
	this.size--
}

func (this *MinStack) Top() int {
	return this.stack[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.size]
}
