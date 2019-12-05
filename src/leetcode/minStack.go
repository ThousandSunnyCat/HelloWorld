package leetcode

// 此题有惊艳解法

type MinStack struct {
	value []interface{}
	min []interface{}
}


/** initialize your data structure here. */
func MinConstructor() MinStack {
    return MinStack {
		value: make([]interface{}, 0),
		min: make([]interface{}, 0),
	}
}


func (this *MinStack) Push(x int) {
	if this.IsEmpty() || this.GetMin() >= x {
		this.min = append(this.min, x)
	}
	this.value = append(this.value, x)
}


func (this *MinStack) Pop() {
	if this.IsEmpty() {
		return
	}

	l := len(this.value)
	if this.GetMin() == this.value[l-1] {
		this.min = this.min[:len(this.min)-1]
	}
	
	this.value = this.value[:l-1]
}


func (this *MinStack) Top() int {
    if this.IsEmpty() {
		return 0
	}
	return this.value[len(this.value)-1].(int)
}


func (this *MinStack) GetMin() int {
    if this.IsEmpty() {
		return 0
	}
	
	return this.min[len(this.min)-1].(int)
}

func (this *MinStack) IsEmpty() bool {
	return len(this.value) < 1
}