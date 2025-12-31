package p1970

type DSU []int

func NewDSU(n int) (dsu DSU) {
	dsu = make(DSU, n)
	for i := range dsu {
		dsu[i] = i
	}
	return
}

func (dsu *DSU) Find(x int) int {
	if (*dsu)[x] == x {
		return x
	}
	(*dsu)[x] = dsu.Find((*dsu)[x]) // сжатие путей
	return (*dsu)[x]
}

func (dsu *DSU) Union(i, j int) {
	rootI := dsu.Find(i)
	rootJ := dsu.Find(j)
	if rootI != rootJ {
		if rootI < rootJ {
			(*dsu)[rootJ] = rootI
		} else {
			(*dsu)[rootI] = rootJ
		}
	}
}

func latestDayToCross(row int, col int, cells [][]int) int {

	dsu := NewDSU(row*col + 2)
	isNotDry := make([]bool, row*col+2)
	l := len(cells)
	for i := 0; i < l; i++ {
		ri, ci := cells[i][0], cells[i][1]
		isNotDry[(ri-1)*col+ci] = true
	}

	f := func(ri, j int) {
		curIndex := (ri-1)*col + j
		if !isNotDry[curIndex] {
			if ri == 1 {
				dsu.Union(0, curIndex)
			}

			// Верх
			upIndex := (ri-2)*col + j
			if ri > 1 && !isNotDry[upIndex] {
				dsu.Union(upIndex, curIndex)
			}
			// Лево
			leftIndex := (ri-1)*col + j - 1
			if j > 1 && !isNotDry[leftIndex] {
				dsu.Union(leftIndex, curIndex)
			}
			// Право
			rightIndex := (ri-1)*col + j + 1
			if j < col && !isNotDry[rightIndex] {
				dsu.Union(rightIndex, curIndex)
			}
			// Низ
			downIndex := (ri)*col + j
			if ri < row && !isNotDry[(ri)*col+j] {
				dsu.Union(downIndex, curIndex)
			}

			if ri == row {
				dsu.Union(row*col+1, curIndex)
			}
		}
	}

	for ri := 2; ri < row; ri++ {
		for j := 1; j <= col; j++ {
			f(ri, j)
		}
	}

	for i := l - 1; i >= 0; i-- {
		ri, ci := cells[i][0], cells[i][1]
		isNotDry[(ri-1)*col+ci] = false
		f(ri, ci)
		if dsu.Find(0) == dsu.Find(row*col+1) {
			return i
		}
	}

	return 0
}
