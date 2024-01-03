package p0452

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	l := len(points)
	if l == 1 {
		return l
	}
	mpoints := map[int]int{}
	sort.Slice(points, func(i, j int) bool {
		mpoints[points[i][0]]++
		mpoints[points[i][1]]++
		mpoints[points[j][0]]++
		mpoints[points[j][1]]++
		return points[i][1]-points[i][0] < points[j][1]-points[j][0]
	})
	//fmt.Println(points)
	keys := make([]int, len(mpoints))
	index := 0
	for k := range mpoints {
		keys[index] = k
		index++
	}
	sort.Ints(keys)
	count := 0
	for l > 0 {
		minArr := [][]int{}
		for i := 0; i < len(keys); i++ {
			flgDelere := false
			k := keys[i]
			if points[0][0] <= k && k <= points[0][1] {
				tmp := make([][]int, l)
				copy(tmp, points)
				for j := l - 1; j >= 0; j-- {
					if tmp[j][0] <= k && k <= tmp[j][1] {
						tmp = append(tmp[:j], tmp[j+1:]...)
						//fmt.Println(tmp)
						flgDelere = true
					}
				}
				if len(tmp) <= len(minArr) || flgDelere && len(minArr) == 0 {
					minArr = tmp
				}
				if len(minArr) == 0 {
					return count + 1
				}
			} else if points[0][0] > k {
				continue
			} else {
				break
			}
		}
		//fmt.Println()
		fmt.Println(len(minArr))
		count++
		points = minArr
		l = len(points)
	}
	return count
}
