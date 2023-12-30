package p0216

import "fmt"

type arrSum struct {
	arr []int
	sum int
}

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	if k > n {
		return res
	}
	arr := make([]arrSum, min(n, 9))
	tmp := make([]arrSum, min(n, 9))
	for i := 1; i <= len(arr); i++ {
		arr[i-1].arr = []int{i}
		arr[i-1].sum = i
		tmp[i-1].arr = []int{i}
		tmp[i-1].sum = i
	}
	tmp = tmp[1:]
	for len(tmp) > 0 {
		newArr := []arrSum{}
		for i := 0; i < len(arr); i++ {
			x := arr[i]
			for j := 0; j < len(tmp); j++ {
				v := tmp[j]
				if x.sum+v.sum <= n && v.sum > x.arr[len(x.arr)-1] {
					if len(x.arr)+1 == k && x.sum+v.sum == n {
						newArr = append(newArr, arrSum{append(x.arr, v.sum), x.sum + v.sum})
					} else if len(x.arr)+1 != k && x.sum+v.sum < n {
						newArr = append(newArr, arrSum{append(x.arr, v.sum), x.sum + v.sum})
					}
					fmt.Println(len(x.arr)+1, k, v.sum, x.sum+v.sum, newArr)
				}
			}
			if len(newArr) == 0 {
				break
			}
		}
		tmp = tmp[1:]
		//fmt.Println(newArr)
		if len(newArr) > 0 {
			arr = newArr
		}

	}
	for _, v := range arr {
		if v.sum == n {
			res = append(res, v.arr)
		}
	}

	fmt.Println(arr)
	return res
}
