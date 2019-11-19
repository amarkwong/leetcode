/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	int Val
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return s:=make([]MinStack, 0) 
}

func (this *MinStack) Push(x int)  {
	append(MinStack, x)
}


func (this *MinStack) Pop()  {
	l := len(MinStack)
	if l == 0 {
		err:=fmt.Errorf("Current length is 0, can not pop")
		fmt.Println(err.Error)
	}
	MinStack=MinStack[:l-1]
}


func (this *MinStack) Top() int {
	l := len(s)
	if l == 0 {
		err:=fmt.Errorf("Current length is 0, can not top")
		fmt.Println(err.Error)
	}
	MinStack=MinStack[:l-1]
	return MinStack[l-1]
}


func (this *MinStack) GetMin() int {
	return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

