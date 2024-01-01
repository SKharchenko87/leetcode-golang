package p2

import (
	"fmt"
	"sort"
)

func maximumLength(s string) int {
	m := map[int32][]int{'a': []int{}, 'b': []int{}, 'c': []int{}, 'd': []int{}, 'e': []int{}, 'f': []int{}, 'g': []int{}, 'h': []int{}, 'i': []int{}, 'j': []int{}, 'k': []int{}, 'l': []int{}, 'm': []int{}, 'n': []int{}, 'o': []int{}, 'p': []int{}, 'q': []int{}, 'r': []int{}, 's': []int{}, 't': []int{}, 'u': []int{}, 'v': []int{}, 'w': []int{}, 'x': []int{}, 'y': []int{}, 'z': []int{}}
	for i, c := range s {
		m[c] = append(m[c], i)
	}
	fmt.Println(m)
	gmax := -1
	for k, arr := range m {
		l := len(arr)
		if l > 2 {
			sort.Ints(arr)
			fmt.Println(m)
			x := []int{}
			cnt := 0
			i := 0
			for ; i < l; i++ {
				if i > 0 && arr[i]-arr[i-1] != 1 {
					x = append(x, cnt)
					cnt = 0
				}
				cnt++
			}
			if i == l && cnt != 0 {
				x = append(x, cnt)
			}
			sort.Ints(x)
			lx := len(x)
			var max1, max2, max3 int
			if lx > 2 {
				if x[lx-1] == x[lx-2] && x[lx-2] == x[lx-3] {
					max1 = x[lx-1]
				} else {
					max1 = x[lx-1] - 2
					if x[lx-1]-x[lx-2] > 0 {
						max2 = x[lx-2]
					}
					max3 = x[lx-3]
				}
			} else if lx > 1 {
				if x[lx-1] == x[lx-2] {
					max1 = x[lx-1] - 1
				} else {
					max1 = x[lx-1] - 2
				}
				if x[lx-1]-x[lx-2] > 0 {
					max2 = x[lx-2]
				}
			} else if lx > 0 {
				max1 = x[lx-1] - 2
			}
			mx := max(max1, max2, max3)
			if gmax < mx {
				gmax = mx
			}
		} else {
			m[k] = []int{}
		}

	}
	fmt.Println(m)
	return gmax
}
