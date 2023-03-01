package Offer

type CQueue struct {
	inStack, outStack []int
	inLen, outLen     int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
	this.inLen++
}

func (this *CQueue) DeleteHead() int {
	if this.outLen == 0 {
		if this.inLen == 0 {
			return -1
		} else {
			for this.inLen > 0 {
				tail := this.inStack[this.inLen-1]
				this.outStack = append(this.outStack, tail)
				this.inStack = this.inStack[:this.inLen-1]
				this.inLen--
				this.outLen++
			}
		}
	}
	res := this.outStack[this.outLen-1]
	this.outStack = this.outStack[:this.outLen-1]
	this.outLen--
	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
