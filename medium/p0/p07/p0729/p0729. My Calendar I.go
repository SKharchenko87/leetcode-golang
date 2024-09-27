package p0729

import "slices"

type MyCalendar struct {
	arr [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{arr: [][2]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	book := [2]int{start, end}
	index, _ := slices.BinarySearchFunc(this.arr, book, func(a [2]int, b [2]int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})
	if index > 0 && this.arr[index-1][1] > start {
		return false
	}
	if index < len(this.arr) && end > this.arr[index][0] {
		return false
	}
	this.arr = slices.Insert(this.arr, index, book)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

func run(cmd []string, books [][2]int) []bool {
	obj := Constructor()
	res := make([]bool, len(cmd)-1)
	if len(cmd) > 1 {
		obj.arr = append(obj.arr, books[1])
		res[0] = true
	}
	for i := 2; i < len(cmd); i++ {
		book := books[i]
		res[i-1] = obj.Book(book[0], book[1])
	}
	return res
}
