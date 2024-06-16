package p1122

import "sort"

// ToDo banchmark

func relativeSortArray(arr1 []int, arr2 []int) []int {
	cntArr := [1001]int{}
	for _, e := range arr1 {
		cntArr[e]++
	}
	index := -1
	for _, e := range arr2 {
		cnt := cntArr[e]
		for j := 0; j < cnt; j++ {
			index++
			arr1[index] = e
			cntArr[e]--
		}
	}
	for i := 0; i <= 1000; i++ {
		e := cntArr[i]
		if e > 0 {
			for j := 0; j < e; j++ {
				index++
				arr1[index] = i
			}
		}
	}
	return arr1
}

func relativeSortArray5(arr1 []int, arr2 []int) []int {
	cntArr2 := map[int]int{}
	for i, e := range arr2 {
		cntArr2[e] = i + 1
	}
	sort.Slice(arr1, func(i, j int) bool {
		eI, okI := cntArr2[arr1[i]]
		eJ, okJ := cntArr2[arr1[j]]

		if !okI && !okJ {
			return arr1[i] < arr1[j]
		}
		if okI && !okJ {
			return true
		}
		if !okI && okJ {
			return false
		}
		if eI < eJ {
			return true
		}
		return false
	})
	return arr1
}

func relativeSortArray4(arr1 []int, arr2 []int) []int {
	cntArr2 := [1001]int{}
	for i, e := range arr2 {
		cntArr2[e] = i + 1
	}
	sort.Slice(arr1, func(i, j int) bool {
		if cntArr2[arr1[i]] == 0 && cntArr2[arr1[j]] == 0 {
			return arr1[i] < arr1[j]
		}
		if cntArr2[arr1[i]] != 0 && cntArr2[arr1[j]] == 0 {
			return true
		}
		if cntArr2[arr1[i]] == 0 && cntArr2[arr1[j]] != 0 {
			return false
		}
		if cntArr2[arr1[i]] < cntArr2[arr1[j]] {
			return true
		}
		return false
	})
	return arr1
}

func relativeSortArray3(arr1 []int, arr2 []int) []int {
	cntArr := make([]int, 333, 1001)
	minArr1, maxArr1 := 1001, -1
	for _, e := range arr1 {
		if e >= len(cntArr) {
			tmpCntArr := make([]int, e-len(cntArr)+1)
			cntArr = append(cntArr, tmpCntArr...)
		}
		cntArr[e]++
		if e > maxArr1 {
			maxArr1 = e
		}
		if e < minArr1 {
			minArr1 = e
		}
	}
	index := -1
	for _, e := range arr2 {
		cnt := cntArr[e]
		for j := 0; j < cnt; j++ {
			index++
			arr1[index] = e
			cntArr[e]--
		}
	}
	for i := minArr1; i <= maxArr1; i++ {
		e := cntArr[i]
		if e > 0 {
			for j := 0; j < e; j++ {
				index++
				arr1[index] = i
			}
		}
	}
	return arr1
}

func relativeSortArray2(arr1 []int, arr2 []int) []int {
	cntArr := map[int]int{}
	minArr1, maxArr1 := 1001, -1
	for _, e := range arr1 {
		cntArr[e]++
		if e > maxArr1 {
			maxArr1 = e
		}
		if e < minArr1 {
			minArr1 = e
		}
	}
	index := -1
	for _, e := range arr2 {
		cnt := cntArr[e]
		for j := 0; j < cnt; j++ {
			index++
			arr1[index] = e
			cntArr[e]--
		}
	}
	for i := minArr1; i <= maxArr1; i++ {
		e := cntArr[i]
		if e > 0 {
			for j := 0; j < e; j++ {
				index++
				arr1[index] = i
			}
		}
	}
	return arr1
}

func relativeSortArray1(arr1 []int, arr2 []int) []int {
	cntArr := make(map[int]int, 1001)
	minArr1, maxArr1 := 1001, -1
	for _, e := range arr1 {
		cntArr[e]++
		if e > maxArr1 {
			maxArr1 = e
		}
		if e < minArr1 {
			minArr1 = e
		}
	}
	index := -1
	for _, e := range arr2 {
		cnt := cntArr[e]
		for j := 0; j < cnt; j++ {
			index++
			arr1[index] = e
			cntArr[e]--
		}
	}
	for i := minArr1; i <= maxArr1; i++ {
		e := cntArr[i]
		if e > 0 {
			for j := 0; j < e; j++ {
				index++
				arr1[index] = i
			}
		}
	}
	return arr1
}

func relativeSortArray0(arr1 []int, arr2 []int) []int {
	cntArr := [1001]int{}
	minArr1, maxArr1 := 1001, -1
	for _, e := range arr1 {
		cntArr[e]++
		if e > maxArr1 {
			maxArr1 = e
		}
		if e < minArr1 {
			minArr1 = e
		}
	}
	index := -1
	for _, e := range arr2 {
		cnt := cntArr[e]
		for j := 0; j < cnt; j++ {
			index++
			arr1[index] = e
			cntArr[e]--
		}
	}
	for i := minArr1; i <= maxArr1; i++ {
		e := cntArr[i]
		if e > 0 {
			for j := 0; j < e; j++ {
				index++
				arr1[index] = i
			}
		}
	}
	return arr1
}
