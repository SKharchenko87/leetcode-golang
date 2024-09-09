package p2326

import . "leetcode/stucture"

var direction = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralMatrix3(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = -1
		}
	}
	directionIndex := 0
	l := n
	size := m * n
	m--
	i, j := 0, -1
	k := 0
	var setMatrixValue = func(value int) {
		if l == 0 {
			directionIndex = (directionIndex + 1) % 4
			if directionIndex%2 == 0 {
				m--
				l = n
			} else {
				n--
				l = m
			}
		}
		i += direction[directionIndex][0]
		j += direction[directionIndex][1]
		res[i][j] = value
		l--
	}
	for ; head != nil && k < size; k++ {
		setMatrixValue(head.Val)
		head = head.Next
	}
	return res
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	directionIndex := 0
	l, size, i, j, k, m := n, m*n, 0, -1, 0, m-1
	var setMatrixValue = func(value int) {
		if l == 0 {
			directionIndex = (directionIndex + 1) % 4
			if directionIndex%2 == 0 {
				m, l = m-1, n
			} else {
				n, l = n-1, m
			}
		}
		i += direction[directionIndex][0]
		j += direction[directionIndex][1]
		res[i][j] = value
		l--
	}
	for ; head != nil && k < size; head, k = head.Next, k+1 {
		setMatrixValue(head.Val)
	}
	for ; k < size; k++ {
		setMatrixValue(-1)
	}
	return res
}

func spiralMatrix1(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	direction := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	directionIndex := 0

	size := m * n

	i, j := 0, -1
	k := 0
	res[0] = make([]int, n)
	for ; k < n; k++ {
		i += direction[directionIndex][0]
		j += direction[directionIndex][1]
		if head == nil {
			res[i][j] = -1
		} else {
			res[i][j] = head.Val
			head = head.Next
		}
	}
	m--
	directionIndex = 1
	for ; k < n+m; k++ {
		i += direction[directionIndex][0]
		j += direction[directionIndex][1]
		res[i] = make([]int, n)
		if head == nil {
			res[i][j] = -1
		} else {
			res[i][j] = head.Val
			head = head.Next
		}
	}
	n--
	l := 0
	for ; k < size; k++ {
		if l == 0 {
			directionIndex = (directionIndex + 1) % 4
			if directionIndex%2 == 0 {
				m--
				l = n
			} else {
				n--
				l = m
			}
		}
		i += direction[directionIndex][0]
		j += direction[directionIndex][1]
		if head == nil {
			res[i][j] = -1
		} else {
			res[i][j] = head.Val
			head = head.Next
		}
		l--
	}
	return res
}

func spiralMatrix0(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	//direction := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	directionIndex := 0
	l := n
	size := m * n
	m--
	i, j := 0, -1
	for k := 0; k < size; k++ {
		if l == 0 {
			directionIndex = (directionIndex + 1) % 4
			if directionIndex%2 == 0 {
				m--
				l = n
			} else {
				n--
				l = m
			}
		}
		i += direction[directionIndex][0]
		j += direction[directionIndex][1]
		if head == nil {
			res[i][j] = -1
		} else {
			res[i][j] = head.Val
			head = head.Next
		}
		l--
	}
	return res
}

func run(m int, n int, head []int) [][]int {
	return spiralMatrix(m, n, ArrToList(head))
}
