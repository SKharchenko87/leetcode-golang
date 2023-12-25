package main

import (
	"fmt"
)

func main() {
	recentCounter := Constructor()
	fmt.Println(recentCounter.ping(1))    // requests = [1], range is [-2999,1], return 1
	fmt.Println(recentCounter.ping(100))  // requests = [1, 100], range is [-2900,100], return 2
	fmt.Println(recentCounter.ping(3001)) // requests = [1, 100, 3001], range is [1,3001], return 3
	fmt.Println(recentCounter.ping(3002)) // requests = [1, 100, 3001, 3002], range is [2,3002], return 3

}

type RecentCounter struct {
	arr []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) ping(t int) int {
	this.arr = append(this.arr, t)
	l := len(this.arr)
	t3000 := t - 3000
	for i := l - 1; i >= 0; i-- {
		if t3000 > this.arr[i] {
			return l - 1 - i
		}
	}
	return l
}

func (this *RecentCounter) Ping(t int) int {
	return this.ping(t)
}
