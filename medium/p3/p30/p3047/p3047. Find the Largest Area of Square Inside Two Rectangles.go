package p3047

import "sort"

type W struct {
	a, b [][]int
}

func (w W) Len() int {
	return len(w.a)
}

func (w W) Less(i, j int) bool {
	if w.a[i][0] == w.a[j][0] {
		return w.b[i][0] < w.b[j][0]
	}
	return w.a[i][0] < w.a[j][0]
}

func (w W) Swap(i, j int) {
	w.a[i], w.a[j] = w.a[j], w.a[i]
	w.b[i], w.b[j] = w.b[j], w.b[i]
}

func largestSquareArea(bottomLeft, topRight [][]int) int64 {
	w := W{bottomLeft, topRight}
	sort.Sort(w)
	var res int64
	for i := 0; i < w.Len()-1; i++ {
		for j := i + 1; j < w.Len(); j++ {
			xA0, yA0, xA1, yA1 := bottomLeft[i][0], bottomLeft[i][1], topRight[i][0], topRight[i][1]
			xB0, yB0, xB1, yB1 := bottomLeft[j][0], bottomLeft[j][1], topRight[j][0], topRight[j][1]
			if xA0 <= xB0 && xB0 < xA1 {
				res = max(res, Intersection(xA0, yA0, xA1, yA1, xB0, yB0, xB1, yB1))
			} else {
				break
			}
		}
	}
	return res
}

func Intersection(xA0, yA0, xA1, yA1 int, xB0, yB0, xB1, yB1 int) int64 {
	var dx, dy int
	dx = min(xA1, xB1) - xB0
	if yA0 <= yB0 && yB0 <= yA1 {
		dy = min(yA1, yB1) - yB0
	} else if yA0 <= yB1 && yB1 <= yA1 {
		dy = min(yA1, yB1) - yA0
	} else if yB0 <= yA0 && yA1 <= yB1 {
		dy = yA1 - yA0
	}
	d := min(dx, dy)
	return int64(d * d)
}
