package p0901

type StockSpanner struct {
	price int
	prev  *StockSpanner
	last  *StockSpanner
}

func Constructor() StockSpanner {
	res := StockSpanner{}
	return res
}

func (this *StockSpanner) Next(price int) int {
	cnt := 0
	next := StockSpanner{price, this.last, nil}
	next.last = &next
	this.last = &next
	x := this.last
	for x != nil {
		if x.price <= price {
			cnt++
		} else {
			break
		}
		x = x.prev
	}
	return cnt
}

func p0901(comands []string, values [][]int) []int {
	obj := Constructor()
	l := len(values)
	res := make([]int, l-1)
	for i := 1; i < l; i++ {
		res[i-1] = obj.Next(values[i][0])
	}
	return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
