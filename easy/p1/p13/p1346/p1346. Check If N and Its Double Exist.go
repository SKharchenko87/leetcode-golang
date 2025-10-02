package p1346

import (
	"slices"
	"sort"
)

func checkIfExist(arr []int) bool {
	l := len(arr)
	m := make(map[int]struct{}, l)
	for _, v := range arr {
		if _, ok := m[v*2]; ok {
			return true
		} else if _, ok := m[v/2]; ok && v&1 == 0 {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}

func checkIfExist3(arr []int) bool {
	cnts := [20001]int{}
	for _, v := range arr {
		index := v + 10000
		cnts[index]++
		if v == 0 {
			if cnts[index] > 1 {
				return true
			}
		} else {
			if v&1 == 0 {
				if cnts[v/2+10000] > 0 {
					return true
				}
			}
			if -10000 <= v*2 && v*2 <= 10000 && cnts[v*2+10000] > 0 {
				return true
			}
		}
	}
	return false
}

func checkIfExist2(arr []int) bool {
	l := len(arr)
	sort.Ints(arr)
	for i := 0; i < l-1; i++ {
		x := arr[i]
		if x >= 0 {
			index, flg := slices.BinarySearch(arr, x*2)
			if flg {
				if x != 0 || x == 0 && index+1 < l && arr[index+1] == 0 {
					return true
				}
			}
		} else if x&1 == 0 {
			_, flg := slices.BinarySearch(arr, x/2)
			if flg {
				return true
			}
		}

	}
	return false
}

func checkIfExist1(arr []int) bool {
	l := len(arr)
	sort.Ints(arr)
	for i := 0; i < l-1; i++ {
		x := arr[i]
		if x >= 0 {
			index, flg := slices.BinarySearch(arr[i:], x*2)
			if flg {
				if x != 0 || x == 0 && i+index+1 < l && arr[i+index+1] == 0 {
					return true
				}
			}
		} else if x&1 == 0 {
			_, flg := slices.BinarySearch(arr[i:], x/2)
			if flg {
				return true
			}
		}

	}
	return false
}

func checkIfExist0(arr []int) bool {
	l := len(arr)
	m := make(map[int]struct{}, l)
	zeroCnt := 0
	for _, v := range arr {
		if v == 0 {
			zeroCnt++
			if zeroCnt == 2 {
				return true
			}
		} else {
			if _, ok := m[v*2]; ok {
				return true
			}
			if _, ok := m[v/2]; ok && v&1 == 0 {
				return true
			}
			m[v] = struct{}{}
		}
	}
	return false
}
