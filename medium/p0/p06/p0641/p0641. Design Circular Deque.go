package p0641

type MyCircularDeque struct {
	queue      []int
	start, end int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{queue: make([]int, k+1), start: -1, end: -1}
}

func (this *MyCircularDeque) plus(value int) int {
	return (value + 1) % len(this.queue)
}

func (this *MyCircularDeque) minus(value int) int {
	return (len(this.queue) + value - 1) % len(this.queue)
}

func (this *MyCircularDeque) insert(value int) (bool, bool) {
	if this.IsFull() {
		return true, false
	}
	if this.IsEmpty() {
		this.queue[0] = value
		this.start = 0
		this.end = 0
		return true, true
	}
	return false, true
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	flg, res := this.insert(value)
	if flg {
		return res
	}
	this.start = this.minus(this.start)
	this.queue[this.start] = value
	return res
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	flg, res := this.insert(value)
	if flg {
		return res
	}
	this.end = this.plus(this.end)
	this.queue[this.end] = value
	return res
}

func (this *MyCircularDeque) delete() (bool, bool) {
	if this.IsEmpty() {
		return true, false
	}
	if this.start == this.end {
		this.start = -1
		this.end = -1
		return true, true
	}
	return false, true
}

func (this *MyCircularDeque) DeleteFront() bool {
	flg, res := this.delete()
	if !flg {
		this.start = this.plus(this.start)
	}
	return res
}

func (this *MyCircularDeque) DeleteLast() bool {
	flg, res := this.delete()
	if !flg {
		this.end = this.minus(this.end)
	}
	return res
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	res := this.queue[this.start]
	return res
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	res := this.queue[this.end]
	return res
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.start == -1 && this.end == -1
}

func (this *MyCircularDeque) IsFull() bool {
	l := 0
	if this.start > this.end {
		l = len(this.queue) - this.start + 1 + this.end
	} else {
		l = this.end - this.start + 1
	}
	return l == len(this.queue)-1
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func run(cmd []string, values [][]int) []interface{} {
	deque := Constructor(values[0][0])
	res := make([]interface{}, len(cmd))
	for i := 1; i < len(cmd); i++ {
		switch cmd[i] {
		case "insertFront":
			res[i] = deque.InsertFront(values[i][0])
		case "insertLast":
			res[i] = deque.InsertLast(values[i][0])
		case "deleteFront":
			res[i] = deque.DeleteFront()
		case "deleteLast":
			res[i] = deque.DeleteLast()
		case "getFront":
			res[i] = deque.GetFront()
		case "getRear":
			res[i] = deque.GetRear()
		case "isEmpty":
			res[i] = deque.IsEmpty()
		case "isFull":
			res[i] = deque.IsFull()
		}
	}
	return res
}
