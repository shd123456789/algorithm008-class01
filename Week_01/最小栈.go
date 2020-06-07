// 最小栈与原栈同步
type MinStack struct {
    stack []int
    minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stack : []int{},
        minStack : []int{math.MaxInt64},
    }
}


func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    min := this.minStack[len(this.minStack) - 1]
    if min > x {
        min = x
    }
    this.minStack = append(this.minStack, min)
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack) - 1]
    this.minStack = this.minStack[:len(this.minStack) - 1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack) - 1]
}

// 最小栈与原栈不同步
type MinStack struct {
    stack []int
    minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stack : []int{},
        minStack : []int{math.MaxInt64},
    }
}


func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    min := this.minStack[len(this.minStack) - 1]
    if min >= x {
        this.minStack = append(this.minStack, x)
    }
}


func (this *MinStack) Pop()  {
    if this.stack[len(this.stack) - 1] == this.minStack[len(this.minStack) - 1] {
        this.minStack = this.minStack[:len(this.minStack) - 1]
    }
    this.stack = this.stack[:len(this.stack) - 1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack) - 1]
}


