package p2179

type FenwickTree struct {
	size int
	tree []int
}

func NewFenwickTree(size int) *FenwickTree {
	return &FenwickTree{
		size: size,
		tree: make([]int, size+1),
	}
}

func (ft *FenwickTree) Update(index int, delta int) {
	for index <= ft.size {
		ft.tree[index] += delta
		index += index & -index
	}
}

func (ft *FenwickTree) Query(index int) int {
	res := 0
	for index > 0 {
		res += ft.tree[index]
		index -= index & -index
	}
	return res
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	pos2 := make([]int, n)
	for idx, num := range nums2 {
		pos2[num] = idx + 1 // 1-based indexing
	}

	a := make([]int, n)
	for i, num := range nums1 {
		a[i] = pos2[num]
	}

	left := make([]int, n)
	fenwick := NewFenwickTree(n)
	for i := 0; i < n; i++ {
		left[i] = fenwick.Query(a[i] - 1)
		fenwick.Update(a[i], 1)
	}

	right := make([]int, n)
	fenwick = NewFenwickTree(n)
	for i := n - 1; i >= 0; i-- {
		right[i] = (n - 1 - i) - fenwick.Query(a[i])
		fenwick.Update(a[i], 1)
	}

	var res int64 = 0
	for i := 0; i < n; i++ {
		res += int64(left[i] * right[i])
	}

	return res
}
