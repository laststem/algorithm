type MinStack struct {
    data []element
}

type element struct {
    val int
    min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    s := MinStack{}
    return s
}


func (this *MinStack) Push(val int)  {
    elem := element{val: val, min: val}
    if len(this.data) > 0 {
        top := this.data[len(this.data)-1]
        if top.min <= val {
            elem.min = top.min
        }
    }
    this.data = append(this.data, elem)
}


func (this *MinStack) Pop()  {
    this.data = this.data[:len(this.data)-1]
}


func (this *MinStack) Top() int {
    return this.data[len(this.data)-1].val
}


func (this *MinStack) GetMin() int {
    return this.data[len(this.data)-1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
