package p0455

import (
	"sort"
	"sync"
)

func sortParallel(wg *sync.WaitGroup, s []int) {
	defer wg.Done()
	sort.Ints(s)
}

func findContentChildren(g []int, s []int) int {
	var wg sync.WaitGroup
	wg.Add(2)
	go sortParallel(&wg, g)
	go sortParallel(&wg, s)
	wg.Wait()
	lg, ls := len(g), len(s)
	j := 0
	res := 0
	for i := 0; i < lg; i++ {
		for ; j < ls; j++ {
			if g[i] <= s[j] {
				res++
				j++
				break
			}
		}
		if j == ls {
			break
		}
	}
	return res
}
