package p0216

import "fmt"

type arrSum struct {
	arr []int
	sum int
}

func combinationSum3(k int, n int) [][]int {
	var minArr, maxArr []int
	var minSum, maxSum int
	switch k {
	case 1:
		if 1 <= n && n <= 9 {
			return [][]int{{n}}
		}
		minArr, minSum = []int{1}, 1
		maxArr, maxSum = []int{9}, 9
	case 2:
		minArr, minSum = []int{1, 2}, 3
		maxArr, maxSum = []int{8, 9}, 17
	case 3:
		minArr, minSum = []int{1, 2, 3}, 1
		maxArr, maxSum = []int{7, 8, 9}, 24
	case 4:
		minArr, minSum = []int{1, 2, 3, 4}, 10
		maxArr, maxSum = []int{6, 7, 8, 9}, 30
	case 5:
		minArr, minSum = []int{1, 2, 3, 4, 5}, 15
		maxArr, maxSum = []int{5, 6, 7, 8, 9}, 35
	case 6:
		minArr, minSum = []int{1, 2, 3, 4, 5, 6}, 21
		maxArr, maxSum = []int{4, 5, 6, 7, 8, 9}, 39
	case 7:
		minArr, minSum = []int{1, 2, 3, 4, 5, 6, 7}, 28
		maxArr, maxSum = []int{3, 4, 5, 6, 7, 8, 9}, 42
	case 8:
		minArr, minSum = []int{1, 2, 3, 4, 5, 6, 7, 8}, 36
		maxArr, maxSum = []int{2, 3, 4, 5, 6, 7, 8, 9}, 44
	case 9:
		if n == 45 {
			return [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}
		} else {
			return [][]int{}
		}
	}
	if !(minSum <= n && n <= maxSum) {
		return [][]int{}
	}
	fmt.Println(minArr, maxArr)

out:
	for i := 1; i <= 9; i++ {
		index := i - 1
		for j := i; j < maxArr[index]; j++ {
			if maxSum-j == n {
				maxSum = n
				maxArr[index] = maxArr[index] - j
				break out
			}
		}
		maxSum = maxSum - maxArr[index] + i
		maxArr[index] = i
	}
	fmt.Println(minArr, maxArr)
	res := [][]int{maxArr}
	couter := 0
	for i := 1; i < k; i++ {
		couter = couter + maxArr[i] - maxArr[i-1] - 1
	}
	fmt.Println(couter)
	return res
	//res := [][]int{}
	//if k > n {
	//	return res
	//}
	//arr := make([]arrSum, min(n, 9))
	//tmp := make([]arrSum, min(n, 9))
	//for i := 1; i <= len(arr); i++ {
	//	arr[i-1].arr = []int{i}
	//	arr[i-1].sum = i
	//	tmp[i-1].arr = []int{i}
	//	tmp[i-1].sum = i
	//}
	//tmp = tmp[1:]
	//for len(tmp) > 0 {
	//	newArr := []arrSum{}
	//	for i := 0; i < len(arr); i++ {
	//		x := arr[i]
	//		for j := 0; j < len(tmp); j++ {
	//			v := tmp[j]
	//			if x.sum+v.sum <= n && v.sum > x.arr[len(x.arr)-1] {
	//				if len(x.arr)+1 == k && x.sum+v.sum == n {
	//					newArr = append(newArr, arrSum{append(x.arr, v.sum), x.sum + v.sum})
	//				} else if len(x.arr)+1 != k && x.sum+v.sum < n {
	//					newArr = append(newArr, arrSum{append(x.arr, v.sum), x.sum + v.sum})
	//				}
	//				fmt.Println(len(x.arr)+1, k, v.sum, x.sum+v.sum, newArr)
	//			}
	//		}
	//		if len(newArr) == 0 {
	//			break
	//		}
	//	}
	//	tmp = tmp[1:]
	//	//fmt.Println(newArr)
	//	if len(newArr) > 0 {
	//		arr = newArr
	//	}
	//
	//}
	//for _, v := range arr {
	//	if v.sum == n {
	//		res = append(res, v.arr)
	//	}
	//}
	//
	//fmt.Println(arr)
	//return res
}
