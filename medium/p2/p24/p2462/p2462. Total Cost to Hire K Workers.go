package p2462

type heap struct {
	arr    []int
	length int
}

type heapChild int

const (
	left heapChild = iota + 1
	right
)

func Constructor(capacity int) heap {
	return heap{make([]int, capacity), 0}
}

func (this *heap) addElement(x int) {
	i := this.length
	this.arr[i] = x
	this.length++
	parent := (i - 1) / 2
	for i > 0 && this.arr[i] < this.arr[parent] {
		this.swap(i, parent)
		i = parent
		parent = (i - 1) / 2
	}
}

func (this *heap) deleteMinElement() int {
	res := this.arr[0]
	this.length--
	this.swap(0, this.length)
	i := -1
	j := 0
	for i != j {
		i = j
		j = this.checkMove(i, j, left)
		j = this.checkMove(i, j, right)
		this.swap(i, j)

	}
	return res
}

func (this *heap) checkMove(i, j int, child heapChild) int {
	c := 2*i + int(child)
	if c < this.length && this.arr[j] > this.arr[c] {
		j = c
	}
	return j
}

func (this *heap) swap(i, j int) {
	this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
}

func totalCost(costs []int, k int, candidates int) int64 {
	l := len(costs)
	if l == 1 {
		return int64(costs[0])
	}
	var res int64
	heapLeft := Constructor(candidates)
	heapRight := Constructor(candidates)
	if 2*candidates >= l {
		candidates = l / 2
	}
	j := 0
	for ; j < l && j < candidates; j++ {
		heapLeft.addElement(costs[j])
		heapRight.addElement(costs[l-j-1])
	}
	indexLeft := j
	indexRight := l - j - 1
	for i := 0; i < k; i++ {
		if heapLeft.length == 0 {
			x := heapRight.deleteMinElement()
			res += int64(x)
			if indexLeft <= indexRight {
				heapRight.addElement(costs[indexRight])
				indexRight--
			}
		} else if heapRight.length == 0 {
			x := heapLeft.deleteMinElement()
			res += int64(x)
			if indexLeft <= indexRight {
				heapLeft.addElement(costs[indexLeft])
				indexLeft++
			}
		} else if heapLeft.arr[0] < heapRight.arr[0] {
			x := heapLeft.deleteMinElement()
			res += int64(x)
			if indexLeft <= indexRight {
				heapLeft.addElement(costs[indexLeft])
				indexLeft++
			}
		} else if heapLeft.arr[0] > heapRight.arr[0] {
			x := heapRight.deleteMinElement()
			res += int64(x)
			if indexLeft <= indexRight {
				heapRight.addElement(costs[indexRight])
				indexRight--
			}
		} else if heapLeft.arr[0] == heapRight.arr[0] {
			if costs[indexRight] >= costs[indexLeft] {
				x := heapLeft.deleteMinElement()
				res += int64(x)
				if indexLeft <= indexRight {
					heapLeft.addElement(costs[indexLeft])
					indexLeft++
				}
			} else {
				x := heapRight.deleteMinElement()
				res += int64(x)
				if indexLeft <= indexRight {
					heapRight.addElement(costs[indexRight])
					indexRight--
				}
			}

		}

	}
	return res
}

//func totalCost(costs []int, k int, candidates int) int64 {
//	var count int64
//	var indexCurrent int
//	for i := 0; i < k; i++ {
//		l := len(costs)
//		indexMinLeft := 0
//		minLeft := costs[indexMinLeft]
//		indexMinRight := l - 1
//		minRight := costs[indexMinRight]
//		for j := 1; j < l-1 && j < candidates; j++ {
//			if minLeft > costs[j] {
//				indexMinLeft = j
//				minLeft = costs[indexMinLeft]
//			}
//			if minRight > costs[l-j-1] {
//				indexMinRight = l - j - 1
//				minRight = costs[indexMinRight]
//			}
//		}
//
//		if minLeft < minRight {
//			indexCurrent = indexMinLeft
//		} else {
//			indexCurrent = indexMinRight
//		}
//		count += int64(costs[indexCurrent])
//		costs = append(costs[:indexCurrent], costs[indexCurrent+1:]...)
//	}
//	return count
//}
