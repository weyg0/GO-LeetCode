package leetcode

type MyQueue struct {
	inStack, outStack []int
	inLen, outLen     int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
	this.inLen++
}

func (this *MyQueue) Pop() int {
	if this.outLen == 0 {
		this.convert()
	}
	res := this.outStack[this.outLen-1]
	this.outStack = this.outStack[:this.outLen-1]
	this.outLen--
	return res
}

func (this *MyQueue) Peek() int {
	if this.outLen == 0 {
		this.convert()
	}
	return this.outStack[this.outLen-1]
}

func (this *MyQueue) Empty() bool {
	return this.inLen == 0 && this.outLen == 0
}

func (this *MyQueue) convert() {
	for this.inLen > 0 {
		tail := this.inStack[this.inLen-1]
		this.outStack = append(this.outStack, tail)
		this.inStack = this.inStack[:this.inLen-1]
		this.inLen--
		this.outLen++
	}
}
