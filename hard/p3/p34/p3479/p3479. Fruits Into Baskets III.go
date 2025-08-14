package p3479

import (
	"cmp"
	"math"
	"slices"
	"sort"
)

type SegmentTree struct {
	maxVal      int
	left, right *SegmentTree
}

func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	//tmpArr := make([]*SegmentTree, n)
	tmpArr := make([]*SegmentTree, 0, n)
	for i := 0; i < n; i++ {
		tmpArr = append(tmpArr, &SegmentTree{arr[i], nil, nil})
	}

	for len(tmpArr)/2 >= 1 {
		for i := 0; i < len(tmpArr)/2+len(tmpArr)%2; i++ {
			left := tmpArr[i*2]
			leftMaxVal := left.maxVal
			var right *SegmentTree
			rightMaxVal := 0
			if i*2+1 < len(tmpArr) {
				right = tmpArr[i*2+1]
				rightMaxVal = right.maxVal
			}
			segment := &SegmentTree{max(leftMaxVal, rightMaxVal), left, right}
			tmpArr[i] = segment
		}
		tmpArr = tmpArr[:len(tmpArr)/2+len(tmpArr)%2]
	}

	return tmpArr[0]
}

func (st *SegmentTree) getBasket(val int) bool {

	if st == nil || st.maxVal < val {
		return false
	}

	tmpArr := make([]*SegmentTree, 0, 12)
	for {
		tmpArr = append(tmpArr, st)

		if st.left != nil && st.left.maxVal >= val {
			st = st.left
			continue
		} else if st.right != nil && st.right.maxVal >= val {
			st = st.right
			continue
		} else {
			break
		}
	}
	tmpArr[len(tmpArr)-1].maxVal = 0
	for i := len(tmpArr) - 2; i >= 0; i-- {
		current := tmpArr[i]
		leftVal := 0
		if current.left != nil {
			if current.left.maxVal == 0 {
				current.left = nil
			} else {
				leftVal = current.left.maxVal
			}
		}
		rightVal := 0
		if current.right != nil {
			if current.right.maxVal == 0 {
				current.right = nil
			} else {
				rightVal = current.right.maxVal
			}
		}
		current.maxVal = max(leftVal, rightVal)
	}
	return true
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	st := NewSegmentTree(baskets)
	res := len(fruits)
	for _, val := range fruits {
		if st.getBasket(val) {
			res--
		}
	}
	return res
}

// TLE
type Pair struct {
	index    int
	capacity int
}

func numOfUnplacedFruits3(fruits []int, baskets []int) int {
	n := len(fruits)
	pairs := make([]Pair, 0, n)
	for i, basket := range baskets {
		pairs = append(pairs, Pair{i, basket})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].capacity == pairs[j].capacity {
			return pairs[i].index < pairs[j].index // при равной вместимости — более левый
		}
		return pairs[i].capacity < pairs[j].capacity
	})

	used := make([]bool, n)
	res := n
	for _, fruit := range fruits {
		index, _ := slices.BinarySearchFunc(pairs, Pair{0, fruit}, func(a, b Pair) int {
			return cmp.Compare(a.capacity, b.capacity)
		})

		for j := index; j < n; j++ {
			if !used[pairs[j].index] {
				used[pairs[j].index] = true
				res--
				break
			}
		}

	}
	return res
}

// Ok
func numOfUnplacedFruits2(fruits []int, baskets []int) int {
	n := len(baskets)
	lengthPart := 2 * (int(math.Log(float64(n))) + 1)
	l := n / lengthPart
	if n%lengthPart != 0 {
		l++
	}
	parts := make([][]int, l)
	index := 0
	for i := 0; i < l; i++ {
		parts[i] = make([]int, 0, lengthPart)
		for ; index < n && index < (i+1)*lengthPart; index++ {
			parts[i] = append(parts[i], baskets[index])
		}
		sort.Ints(parts[i])
	}
	res := n
LOOP:
	for _, fruit := range fruits {
		for i, part := range parts {
			lengthCurrentPart := len(part)
			if lengthCurrentPart == 0 || part[lengthCurrentPart-1] < fruit {
				continue
			}
			left := i * lengthPart
			end := min(n, left+lengthPart)
			for ; left < end; left++ {
				if baskets[left] >= fruit {
					index, _ := slices.BinarySearch(part, baskets[left])
					parts[i] = slices.Delete(part, index, index+1)
					baskets[left] = 0
					res--
					continue LOOP
				}
			}
		}

	}
	return res
}

func getIndexFreeBasket(freeIndexes *[]int, baskets []int, val int) int {
	for i := 0; i < len(*freeIndexes); i++ {
		idx := (*freeIndexes)[i]
		if baskets[idx] >= val {
			*freeIndexes = slices.Delete(*freeIndexes, i, i+1)
			return idx
		}
	}
	return -1
}

// TLE
func numOfUnplacedFruits1(fruits []int, baskets []int) int {

	freeIndexes := make([]int, len(baskets))
	for i := 0; i < len(baskets); i++ {
		freeIndexes[i] = i
	}
	res := 0
	for i := 0; i < len(fruits); i++ {
		if getIndexFreeBasket(&freeIndexes, baskets, fruits[i]) == -1 {
			res++
		}
	}
	return res
}

type Node struct {
	val  int
	next *Node
}

// TLE
func numOfUnplacedFruits0(fruits []int, baskets []int) int {

	n := len(fruits)
	var root, next *Node
	next = &Node{val: baskets[n-1], next: nil}
	for i := n - 2; i >= 0; i-- {
		root = &Node{val: baskets[i], next: next}
		next = root
	}
	root = &Node{val: 0, next: next}
	notExists := math.MaxInt

	for _, fruit := range fruits {

		if notExists <= fruit {
			continue
		}
		current := root
		for current.next != nil && current.next.val < fruit {
			current = current.next
		}
		if current.next != nil {
			current.next = current.next.next
			n--
			continue
		}
		notExists = fruit
	}

	return n
}
