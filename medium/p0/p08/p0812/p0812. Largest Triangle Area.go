package p0812

import "math"

const actual = 100_000

type point struct {
	x, y int
}

func round(x float64) float64 {
	return math.Round(x*actual) / actual
}

func largestTriangleArea(points [][]int) float64 {
	l := len(points)
	res := 0.0
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				a := point{points[i][0], points[i][1]}
				b := point{points[j][0], points[j][1]}
				c := point{points[k][0], points[k][1]}
				res = max(res, getArea(a, b, c))
			}
		}
	}
	return math.Round(res*actual) / actual
}

func getSection(a, b point) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2.0) + math.Pow(float64(a.y-b.y), 2.0))
}

func getArea(a, b, c point) float64 {
	ab := getSection(a, b)
	bc := getSection(b, c)
	ca := getSection(c, a)
	p := (ab + bc + ca) / 2.0
	pab := round(p - ab)
	pbc := round(p - bc)
	pca := round(p - ca)
	if pab == 0 || pbc == 0 || pca == 0 {
		return 0.0
	}
	area := math.Sqrt(p * (p - ab) * (p - bc) * (p - ca))
	return area
}
