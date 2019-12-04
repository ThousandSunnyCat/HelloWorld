package leetcode

type MyCircularDeque struct {
	p, q int
	length int
    value []interface{}
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque {
		p: k,
		q: 0,
		length: k + 1,
		value: make([]interface{}, k + 1),
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
		return false
	}

	this.value[this.p] = value
	this.p = (this.p-1+this.length)%this.length
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
		return false
	}

	this.value[this.q] = value
	this.q = (this.q+1)%this.length
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
		return false
	}

	this.p = (this.p+1)%this.length
	this.value[this.p] = nil
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
		return false
	}

	this.q = (this.q-1+this.length)%this.length
	this.value[this.q] = nil
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
		return -1
	}
	return this.value[(this.p+1)%this.length].(int)
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
		return -1
	}
	return this.value[(this.q-1+this.length)%this.length].(int)
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return (this.p + 1) % this.length == this.q
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return this.p == this.q
}