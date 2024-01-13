package p2336

type SmallestInfiniteSet struct {
	arr    [1000]int
	length int
	max    int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.length > 0 {
		return this.deleteMinFromHeap()
	} else {
		this.max++
		return this.max
	}
}

// Добавление в Heap
func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.max >= num && !this.isIn(num) {
		this.addToHeap(num)
	}

}

// Удаление минимума в Heap
func (this *SmallestInfiniteSet) deleteMinFromHeap() int {
	res := this.arr[0]
	this.swap(0, this.length-1)
	this.length--
	i, j, c := -1, 0, 0 // i-иднекс локального минимума, j - кандидат локального минимума, c - временная переменная
	for i != j {
		i = j
		c = 2*i + 1
		if c < this.length && this.arr[i] > this.arr[c] {
			j = c
			this.swap(i, j)
		}
		c = 2*i + 2
		if c < this.length && this.arr[i] > this.arr[c] {
			j = c
			this.swap(i, j)
		}
	}
	return res
}

func (this *SmallestInfiniteSet) swap(i, j int) {
	this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
}

func (this *SmallestInfiniteSet) addToHeap(m int) {
	i := this.length
	this.arr[this.length] = m
	this.length++
	var parentId int
	for {
		parentId = (i - 1) / 2
		if parentId >= 0 && this.arr[parentId] > this.arr[i] {
			this.swap(parentId, i)
			i = parentId
		} else {
			return
		}
	}
}
func (this *SmallestInfiniteSet) isIn(x int) bool {
	for i := this.length - 1; i >= 0; i-- {
		if this.arr[i] == x {
			return true
		}
	}
	return false
}

const null = 0

func run(commands []string, args [][]int) []int {
	res := make([]int, len(commands))
	var cur SmallestInfiniteSet
	for i, v := range commands {
		switch v {
		case "SmallestInfiniteSet":
			cur = Constructor()
		case "addBack":
			cur.AddBack(args[i][0])
		case "popSmallest":
			res[i] = cur.PopSmallest()
		}
	}
	return res
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
