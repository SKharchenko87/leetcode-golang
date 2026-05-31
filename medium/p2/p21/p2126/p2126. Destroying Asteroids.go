package p2126

import (
	"container/heap"
	"sort"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	l := len(asteroids)
	heavy := make([]int, 0, l)
	for i := 0; i < l; i++ {
		if mass >= asteroids[i] {
			mass += asteroids[i]
		} else {
			heavy = append(heavy, asteroids[i])
		}
	}
	sort.Ints(heavy)
	for i := 0; i < len(heavy); i++ {
		if mass >= heavy[i] {
			mass += heavy[i]
		} else {
			return false
		}
	}
	return true
}

type MH []int

func (mh *MH) Less(i, j int) bool {
	return (*mh)[i] < (*mh)[j]
}

func (mh *MH) Swap(i, j int) {
	p := *mh
	p[i], p[j] = p[j], p[i]
}

func (mh *MH) Push(x any) {
	*mh = append(*mh, x.(int))
}

func (mh *MH) Pop() any {
	p := *mh
	l := p.Len() - 1
	x := p[l]
	*mh = p[:l]
	return x
}

func (mh *MH) Len() int {
	return len(*mh)
}

func asteroidsDestroyed1(mass int, asteroids []int) bool {
	mh := MH(asteroids)
	heap.Init(&mh)
	for mh.Len() > 0 {
		x := heap.Pop(&mh).(int)
		if x <= mass {
			mass += x
		} else {
			return false
		}
	}
	return true
}

func asteroidsDestroyed0(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	i := 0
	l := len(asteroids)
	for ; i < l; i++ {
		if asteroids[i] <= mass {
			mass += asteroids[i]
		} else {
			return false
		}
	}
	return true
}
