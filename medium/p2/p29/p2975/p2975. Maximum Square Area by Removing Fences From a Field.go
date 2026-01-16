package p2975

import (
	"sort"
)

const MOD = 1e9 + 7

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	hFences = append(hFences, m, 1)
	vFences = append(vFences, n, 1)
	sort.Ints(hFences)
	sort.Ints(vFences)
	hmFences := make(map[int]bool, len(hFences)*10)
	var x int
	for i := 0; i < len(hFences); i++ {
		for j := i + 1; j < len(hFences); j++ {
			x = hFences[j] - hFences[i]
			hmFences[x] = true
		}
	}

	var best int
	for i := 0; i < len(vFences); i++ {
		for j := i + 1; j < len(vFences); j++ {
			x = vFences[j] - vFences[i]
			if x > best && hmFences[x] {
				best = x
			}
		}
	}

	if best == 0 {
		return -1
	}

	return int(int64(best*best) % MOD)
}

func maximizeSquareArea7(m int, n int, hFences []int, vFences []int) int {
	hmFences := map[int]bool{}
	sort.Ints(hFences)
	sort.Ints(vFences)
	hmFences[m-1] = true
	var x int
	for i := 0; i < len(hFences); i++ {
		hmFences[hFences[i]-1] = true
		for j := i + 1; j < len(hFences); j++ {
			x = hFences[j] - hFences[i]
			hmFences[x] = true
		}
		hmFences[m-hFences[i]] = true
	}

	var best int
	x = n - 1
	if hmFences[x] {
		best = x
	}
	for i := 0; i < len(vFences); i++ {
		x = vFences[i] - 1
		if x > best {
			if hmFences[x] {
				best = x
			}
		}

		for j := i + 1; j < len(vFences); j++ {
			x = vFences[j] - vFences[i]
			if x > best {
				if hmFences[x] {
					best = x
				}
			}
		}
		x = n - vFences[i]
		if x > best {
			if hmFences[x] {
				best = x
			}
		}
	}

	if best == 0 {
		return -1
	}

	return int(int64(best*best) % MOD)
}

func maximizeSquareArea6(m int, n int, hFences []int, vFences []int) int {
	hmFences := map[int]struct{}{}
	sort.Ints(hFences)
	sort.Ints(vFences)
	hmFences[m-1] = struct{}{}
	var x int
	for i := 0; i < len(hFences); i++ {
		hmFences[hFences[i]-1] = struct{}{}
		for j := i + 1; j < len(hFences); j++ {
			x = hFences[j] - hFences[i]
			hmFences[x] = struct{}{}
		}
		hmFences[m-hFences[i]] = struct{}{}
	}

	var best int
	x = n - 1
	if _, ok := hmFences[x]; ok {
		best = x
	}
	for i := 0; i < len(vFences); i++ {
		x = vFences[i] - 1
		if x > best {
			if _, ok := hmFences[x]; ok {
				best = x
			}
		}

		for j := i + 1; j < len(vFences); j++ {
			x = vFences[j] - vFences[i]
			if x > best {
				if _, ok := hmFences[x]; ok {
					best = x
				}
			}
		}
		x = n - vFences[i]
		if x > best {
			if _, ok := hmFences[x]; ok {
				best = x
			}
		}
	}

	if best == 0 {
		return -1
	}

	return int(int64(best*best) % MOD)
}

func maximizeSquareArea5(m int, n int, hFences []int, vFences []int) int {
	hmFences := map[int]bool{}
	sort.Ints(hFences)
	sort.Ints(vFences)
	hmFences[m-1] = true
	var x int
	for i := 0; i < len(hFences); i++ {
		hmFences[hFences[i]-1] = true
		for j := i + 1; j < len(hFences); j++ {
			x = hFences[j] - hFences[i]
			hmFences[x] = true
		}
		hmFences[m-hFences[i]] = true
	}

	var best int
	x = n - 1
	if hmFences[x] {
		best = x
	}
	for i := 0; i < len(vFences); i++ {
		x = vFences[i] - 1
		if x > best && hmFences[x] {
			best = x
		}

		for j := i + 1; j < len(vFences); j++ {
			x = vFences[j] - vFences[i]
			if x > best && hmFences[x] {
				best = x
			}

		}
		x = n - vFences[i]
		if x > best && hmFences[x] {
			best = x
		}
	}

	if best == 0 {
		return -1
	}

	return int(int64(best*best) % MOD)
}

func maximizeSquareArea4(m int, n int, hFences []int, vFences []int) int {
	hFences = append(hFences, m, 1)
	vFences = append(vFences, n, 1)
	hmFences := map[int]bool{}
	var x int
	for i := 0; i < len(hFences); i++ {
		for j := i + 1; j < len(hFences); j++ {
			if hFences[j] > hFences[i] {
				x = hFences[j] - hFences[i]
			} else {
				x = hFences[i] - hFences[j]
			}
			hmFences[x] = true
		}
	}

	var best int
	for i := 0; i < len(vFences); i++ {
		for j := i + 1; j < len(vFences); j++ {
			if vFences[j] > vFences[i] {
				x = vFences[j] - vFences[i]
			} else {
				x = vFences[i] - vFences[j]
			}
			if x > best && hmFences[x] {
				best = x
			}
		}
	}

	if best == 0 {
		return -1
	}

	return int(int64(best*best) % MOD)
}

func maximizeSquareArea3(m int, n int, hFences []int, vFences []int) int {
	hmFences := map[int]struct{}{}
	//sort.Ints(hFences)
	hmFences[m-1] = struct{}{}
	var x int
	for i := 0; i < len(hFences); i++ {
		hmFences[hFences[i]-1] = struct{}{}
		for j := i + 1; j < len(hFences); j++ {
			if hFences[j] > hFences[i] {
				x = hFences[j] - hFences[i]
			} else {
				x = hFences[i] - hFences[j]
			}
			hmFences[x] = struct{}{}
		}
		hmFences[m-hFences[i]] = struct{}{}
	}

	var res int64
	res = -1
	//sort.Ints(vFences)
	x = n - 1
	if _, ok := hmFences[x]; ok {
		res = max(res, int64(x*x))
	}
	for i := 0; i < len(vFences); i++ {
		x = vFences[i] - 1
		if _, ok := hmFences[x]; ok {
			res = max(res, int64(x*x))
		}
		for j := i + 1; j < len(vFences); j++ {
			if vFences[j] > vFences[i] {
				x = vFences[j] - vFences[i]
			} else {
				x = vFences[i] - vFences[j]
			}
			//x = vFences[j] - vFences[i]
			if _, ok := hmFences[x]; ok {
				res = max(res, int64(x*x))
			}
		}
		x = n - vFences[i]
		if _, ok := hmFences[x]; ok {
			res = max(res, int64(x*x))
		}
	}

	return int(res % MOD)
}

func maximizeSquareArea2(m int, n int, hFences []int, vFences []int) int {
	hmFences := map[int]struct{}{}
	sort.Ints(hFences)
	hmFences[m-1] = struct{}{}
	for i := 0; i < len(hFences); i++ {
		hmFences[hFences[i]-1] = struct{}{}
		for j := i + 1; j < len(hFences); j++ {
			hmFences[hFences[j]-hFences[i]] = struct{}{}
		}
		hmFences[m-hFences[i]] = struct{}{}
	}

	var res int64
	res = -1
	sort.Ints(vFences)
	x := n - 1
	if _, ok := hmFences[x]; ok {
		res = max(res, int64(x*x))
	}
	for i := 0; i < len(vFences); i++ {
		x = vFences[i] - 1
		if _, ok := hmFences[x]; ok {
			res = max(res, int64(x*x))
		}
		for j := i + 1; j < len(vFences); j++ {
			x = vFences[j] - vFences[i]
			if _, ok := hmFences[x]; ok {
				res = max(res, int64(x*x))
			}
		}
		x = n - vFences[i]
		if _, ok := hmFences[x]; ok {
			res = max(res, int64(x*x))
		}
	}

	return int(res % MOD)
}

// TLE
func maximizeSquareArea1(m int, n int, hFences []int, vFences []int) int {
	hmFences := map[int]struct{}{}
	sort.Ints(hFences)
	hmFences[m-1] = struct{}{}
	for i := 0; i < len(hFences); i++ {
		hmFences[hFences[i]-1] = struct{}{}
		for j := i + 1; j < len(hFences); j++ {
			hmFences[hFences[j]-hFences[i]] = struct{}{}
		}
		hmFences[m-hFences[i]] = struct{}{}
	}

	vmFences := map[int]struct{}{}
	sort.Ints(vFences)
	vmFences[n-1] = struct{}{}
	for i := 0; i < len(vFences); i++ {
		vmFences[vFences[i]-1] = struct{}{}
		for j := i + 1; j < len(vFences); j++ {
			vmFences[vFences[j]-vFences[i]] = struct{}{}
		}
		vmFences[n-vFences[i]] = struct{}{}
	}
	var res int64
	res = -1
	for hk := range hmFences {
		for vk := range vmFences {
			if hk == vk {
				res = max(res, int64(hk*vk))
			}
		}
	}

	return int(res % MOD)
}

func maximizeSquareArea0(m int, n int, hFences []int, vFences []int) int {
	hmFences := map[int]struct{}{}
	sort.Ints(hFences)
	hmFences[m-1] = struct{}{}
	for i := 0; i < len(hFences); i++ {
		hmFences[hFences[i]-1] = struct{}{}
		for j := i + 1; j < len(hFences); j++ {
			hmFences[hFences[j]-hFences[i]] = struct{}{}
		}
		hmFences[m-hFences[i]] = struct{}{}
	}
	hLen := make([]int, 0, len(hmFences))
	for k := range hmFences {
		hLen = append(hLen, k)
	}
	sort.Ints(hLen)

	vmFences := map[int]struct{}{}
	sort.Ints(vFences)
	vmFences[n-1] = struct{}{}
	for i := 0; i < len(vFences); i++ {
		vmFences[vFences[i]-1] = struct{}{}
		for j := i + 1; j < len(vFences); j++ {
			vmFences[vFences[j]-vFences[i]] = struct{}{}
		}
		vmFences[n-vFences[i]] = struct{}{}
	}
	vLen := make([]int, 0, len(vmFences))
	for k := range vmFences {
		vLen = append(vLen, k)
	}
	sort.Ints(vLen)
	i := len(hLen) - 1
	j := len(vLen) - 1
	for ; i >= 0; i-- {

		for ; j > 0 && vLen[j] > hLen[i]; j-- {
		}
		if vLen[j] == hLen[i] {
			return int(int64(vLen[j]) * int64(hLen[i]) % MOD)
		}
	}
	return -1
}
