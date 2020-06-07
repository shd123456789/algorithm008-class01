// 2遍
type MyCircularDeque struct {
    head int
    tail int
    capacity int
    data []int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque {
        head : 0,
        tail : 0,
        capacity : k + 1,
        data : make([]int, k + 1),
    }
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    this.head = (this.head - 1 + this.capacity) % this.capacity
    this.data[this.head] = value
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    this.data[this.tail] = value
    this.tail = (this.tail + 1 + this.capacity) % this.capacity
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.head = (this.head + 1 + this.capacity) % this.capacity
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    this.tail = (this.tail - 1 + this.capacity) % this.capacity
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.head]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[(this.tail - 1 + this.capacity) % this.capacity]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return this.tail == this.head
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return (this.tail + 1) % this.capacity == this.head
}
