package p3625

import (
	"sort"
)

type section struct {
	dx, dy int
	bu, bd int
	l2     int
	cx     int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// TLE
func countTrapezoids(points [][]int) int {
	l := len(points)
	sections := make([]section, 0, l*l)

	var dx, dy, bu, bd, g, l2, cx int
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			cx = -1001
			dx = points[j][0] - points[i][0]
			dy = points[j][1] - points[i][1]
			if dx == 0 {
				cx = points[i][0]
			}

			l2 = dx*dx + dy*dy
			if dx == 0 && dy < 0 {
				dy = -dy
			}
			if dx < 0 {
				dx = -dx
				dy = -dy
			}
			g = gcd(abs(dx), abs(dy))
			dx /= g
			dy /= g

			bu = points[j][1]*points[i][0] - points[i][1]*points[j][0]
			bd = points[j][0] - points[i][0]

			if !(bu == 0 && bd == 0) {
				g = gcd(abs(bu), abs(bd))
				bu /= g
				bd /= g
			}

			sections = append(sections, section{dx, dy, bu, bd, l2, cx})

		}
	}
	sort.Slice(sections, func(i, j int) bool {
		if sections[i].dx*sections[j].dy == sections[j].dx*sections[i].dy {
			if sections[i].bu*sections[j].bd == sections[j].bu*sections[i].bd {
				return sections[i].l2 < sections[j].l2
			}
			return sections[i].bu*sections[j].bd < sections[j].bu*sections[i].bd
		}
		return sections[i].dx*sections[j].dy < sections[j].dx*sections[i].dy
	})
	res := 0
	romb := 0
	for i := 0; i < len(sections)-1; i++ {
		for j := i + 1; j < len(sections) && sections[i].dx*sections[j].dy == sections[j].dx*sections[i].dy; j++ {
			if sections[j].bd != 0 && sections[i].bd != 0 && sections[i].bu*sections[j].bd == sections[j].bu*sections[i].bd || sections[j].bd == 0 && sections[i].bd == 0 && sections[i].cx == sections[j].cx {
				continue
			}
			if sections[i].l2 == sections[j].l2 {
				romb++
			}
			//fmt.Println(i, j)
			res++
		}
	}

	return res - romb/2
}

func countTrapezoids0(points [][]int) int {
	l := len(points)
	lineCoefficient := map[int]map[int]int{}
	sums := map[int]int{}
	var a, b, dx, dy int
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			dx = points[j][0] - points[i][0]
			if dx == 0 {
				b = points[j][0]
				if _, ok := lineCoefficient[10000001]; ok {
					lineCoefficient[10000001][b]++
				} else {
					lineCoefficient[10000001] = map[int]int{b: 1}
				}
				sums[10000001]++
			} else {
				dy = points[j][1] - points[i][1]
				a = dy * 100000 / dx
				b = (points[j][1]*points[i][0] - points[i][1]*points[j][0]) * 100000 / (points[i][0] - points[j][0])
				if _, ok := lineCoefficient[a]; ok {
					lineCoefficient[a][b]++
				} else {
					lineCoefficient[a] = map[int]int{b: 1}
				}
				sums[a]++
			}

		}
	}
	res := 0
	for k, mb := range lineCoefficient {
		for _, v := range mb {
			res += (sums[k] - v) * v
		}

	}
	return res / 2
}
