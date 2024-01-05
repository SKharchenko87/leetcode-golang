package p0300

func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	arrDP := make([]int, l)
	maxLength := 1
	for j := 0; j < l; j++ {
		arrDP[j] = 1
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] && arrDP[j] < arrDP[i]+1 {
				arrDP[j] = arrDP[i] + 1
			}
		}
		if arrDP[j] > maxLength {
			maxLength = arrDP[j]
		}
	}
	return maxLength
}

//func lengthOfLIS(nums []int) int {
//	arr := []int{}
//	for _, item := range nums {
//		idx := bisect_left(&arr, item)
//		if idx == len(arr) {
//			arr = append(arr, item)
//		} else {
//			arr[idx] = item
//		}
//	}
//	return len(arr)
//}
//
//func bisect_left(arr *[]int, target int) int {
//	left := 0
//	right := len(*arr) - 1
//
//	for left <= right {
//		mid := (left + right) / 2
//		if (*arr)[mid] >= target {
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//	}
//	return left
//}

// 10, 9, 2, 5, 3, 7, 101, 18
//func lengthOfLIS(nums []int) int {
//	listPath := [][][]int{}
//	totalLength := 0
//	l := len(nums)
//	indexes := make([]int, l)
//	for i := range indexes {
//		indexes[i] = i
//	}
//	for totalLength <= l {
//		curIndex := 0
//		cur := nums[indexes[curIndex]]
//		tmp := []int{indexes[curIndex], indexes[curIndex] + 1}
//		curArr := [][]int{}
//		flg := true
//		for i := curIndex + 1; i < len(indexes); i++ {
//			if cur < nums[indexes[i]] && flg {
//				cur = nums[indexes[i]]
//				tmp[1] = indexes[i+1]
//			} else if cur < nums[indexes[i]] && !flg {
//				cur = nums[indexes[i]]
//				curIndex = i
//				tmp = []int{indexes[i], indexes[i] + 1}
//				flg = true
//			} else {
//				if flg {
//					curArr = append(curArr, tmp)
//					length := tmp[1] - tmp[0]
//					totalLength += length
//					indexes = append(indexes[:curIndex], indexes[i:]...)
//					i = i - (length)
//				}
//				flg = false
//			}
//		}
//		if flg {
//			curArr = append(curArr, tmp)
//			totalLength += tmp[1] - tmp[0]
//			indexes = append(indexes[:curIndex], indexes[len(indexes):]...)
//		}
//		fmt.Println(curArr)
//		fmt.Println(indexes)
//		fmt.Println(totalLength)
//		listPath = append(listPath, curArr)
//		if len(indexes) == 0 {
//			break
//		}
//	}
//	fmt.Println(listPath)
//
//	path := [][][]int{}
//	for i := len(listPath) - 1; i >= 1; i-- {
//		tmp := make([][]int, len(listPath[i]))
//		for j:=len(listPath[i-1]) -1
//	}
//
//	//totalLength := 0
//
//	return 0
//}
