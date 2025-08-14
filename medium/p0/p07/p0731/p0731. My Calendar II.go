package p0731

import (
	"slices"
)

type MyCalendarTwo struct {
	books [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{books: [][2]int{}}
}

var myCompare = func(ints1 [2]int, ints2 [2]int) int {
	if ints1[0] == ints2[0] {
		return ints1[1] - ints2[1]
	}
	return ints1[0] - ints2[0]
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	c := [2]int{start, end}
	l := len(this.books)
	if l <= 1 {
		if l == 0 {
			this.books = append(this.books, [2]int{start, end})
			return true
		}

		if myCompare(this.books[0], c) < 0 {
			this.books = [][2]int{this.books[0], c}
		} else {
			this.books = [][2]int{c, this.books[0]}
		}
		return true
	}
	index := -1
	indexes := make([]int, 0, l)
	for i := 0; i < l; i++ {
		if end < this.books[i][0] {
			break
		}
		if start > this.books[i][0] || start == this.books[i][0] && end <= this.books[i][1] {
			index++
		}
		if intersection(this.books[i], c) {
			for j := range indexes {
				if intersection(this.books[indexes[j]], this.books[i]) {
					return false
				}
			}
			indexes = append(indexes, i)
		}
	}

	this.books = slices.Insert(this.books, index+1, c)
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

func run(cmd []string, books [][2]int) []bool {
	l := len(cmd)
	res := make([]bool, l-1)
	calendar := Constructor()
	for i := 0; i < l-1; i++ {
		book := books[i+1]
		res[i] = calendar.Book(book[0], book[1])
	}
	return res
}

func intersection(a, b [2]int) bool {
	return max(a[0], b[0]) < min(a[1], b[1])
	//return (a[0]-b[1])*(b[0]-a[1]) > 0
}
