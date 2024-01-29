package p0232

type MyQueue struct {
	_arr []int
}

func Constructor() MyQueue {
	return MyQueue{_arr: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this._arr = append(this._arr, x)
}

func (this *MyQueue) Pop() (x int) {
	x = this._arr[0]
	this._arr = this._arr[1:]
	return
}

func (this *MyQueue) Peek() int {
	return this._arr[0]
}

func (this *MyQueue) Empty() bool {
	return len(this._arr) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func run(command []string, param [][]int) {

}
