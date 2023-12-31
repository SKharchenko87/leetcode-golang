package p3

import (
	"fmt"
	"sort"
)

func maximumLength(s string) int {
	m := map[int32][]int{'a': []int{}, 'b': []int{}, 'c': []int{}, 'd': []int{}, 'e': []int{}, 'f': []int{}, 'g': []int{}, 'h': []int{}, 'i': []int{}, 'j': []int{}, 'k': []int{}, 'l': []int{}, 'm': []int{}, 'n': []int{}, 'o': []int{}, 'p': []int{}, 'q': []int{}, 'r': []int{}, 's': []int{}, 't': []int{}, 'u': []int{}, 'v': []int{}, 'w': []int{}, 'x': []int{}, 'y': []int{}, 'z': []int{}}

	for i, c := range s {
		la := len(m[c])
		if la > 0 && s[i-1] == s[i] {
			m[c][la-1] = m[c][la-1] + 1
		} else {
			m[c] = append(m[c], 1)
		}

	}
	fmt.Println(m)
	gmax := -1
	for _, x := range m {

		sort.Ints(x)
		lx := len(x)
		if lx > 0 {
			var max1, max2, max3 int = -1, -1, -1
			//max1
			if lx > 2 && x[lx-1] == x[lx-2] && x[lx-2] == x[lx-3] {
				max1 = x[lx-1]
			} else if lx > 1 && x[lx-1] == x[lx-2] {
				max1 = x[lx-1] - 1
			} else {
				max1 = x[lx-1] - 2
			}
			//max2
			if lx > 1 && x[lx-1]-x[lx-2] > 0 {
				max2 = x[lx-2]
			}
			//max3
			if lx > 2 {
				max3 = x[lx-3]
			}
			mx := max(max1, max2, max3)
			if gmax < mx {
				gmax = mx
			}
		}
	}
	fmt.Println(m)
	if gmax == 0 {
		return -1
	}
	return gmax
}
