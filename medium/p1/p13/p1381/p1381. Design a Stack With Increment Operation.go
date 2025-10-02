package p1381

type CustomStack struct {
	stack []int
	index int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{stack: make([]int, maxSize), index: -1}
}

func (this *CustomStack) Push(x int) {
	if len(this.stack) == this.index+1 {
		return
	}
	this.index++
	this.stack[this.index] = x
}

func (this *CustomStack) Pop() int {
	if this.index < 0 {
		return -1
	}
	res := this.stack[this.index]
	this.index--
	return res
}

func (this *CustomStack) Increment(k int, val int) {
	for i := min(k-1, this.index); i >= 0; i-- {
		this.stack[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

func run(cmd []string, param [][]int) []interface{} {
	stack := Constructor(param[0][0])
	res := make([]interface{}, len(cmd))
	for i := 1; i < len(cmd); i++ {
		switch cmd[i] {
		case "push":
			stack.Push(param[i][0])
		case "pop":
			res[i] = stack.Pop()
		case "increment":
			stack.Increment(param[i][0], param[i][1])
		}
	}
	return res
}
