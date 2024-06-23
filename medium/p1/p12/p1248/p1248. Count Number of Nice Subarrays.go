package p1248

func numberOfSubarrays(nums []int, k int) int {
	oddCnt := 0
	result := 0
	for i, num := range nums {
		nums[i] = 0
		nums[oddCnt]++
		oddCnt += num & 1
		if j := oddCnt - k; j >= 0 {
			result += nums[j]
		}
	}
	return result
}

func numberOfSubarrays0(nums []int, k int) int {
	//l:=len(nums)
	cntArr := []int{}
	cnt := 0
	res := 0
	for _, num := range nums {
		cnt++
		if num%2 == 1 {
			cntArr = append(cntArr, cnt)
			cnt = 0
			lca := len(cntArr)
			if lca > k {
				res += cntArr[lca-k-1] * cntArr[lca-1]
			}
		}
	}
	lca := len(cntArr)
	if lca >= k {
		res += cntArr[lca-k] * (cnt + 1)
	}
	return res
}

func numberOfSubarrays3(nums []int, k int) int {
	res0 := atMost(nums, k)
	res1 := atMost(nums, k-1)
	return res0 - res1
}

func atMost(nums []int, k int) int {
	cntOdd := 0
	start := 0
	res := 0
	for end, num := range nums {
		cntOdd += num & 1
		for cntOdd > k {
			cntOdd -= nums[start] & 1
			start++
		}
		res += end - start + 1
	}

	return res
}

func numberOfSubarrays2(nums []int, k int) int {
	qsize := 0
	start := 0
	gap := 0
	res := 0
	for _, num := range nums {
		qsize += num & 1
		if qsize == k {
			gap = 0
			for qsize == k {
				qsize -= nums[start] & 1
				gap++
				start++
			}
		}
		res += gap
	}
	return res
}

func numberOfSubarrays1(nums []int, k int) int {
	curSum := 0
	cntMap := make(map[int]int)
	cntMap[curSum] = 1
	res := 0
	for _, num := range nums {
		curSum += num & 1
		if v, ok := cntMap[curSum-k]; ok {
			res += v
		}
		cntMap[curSum]++
	}
	return res
}

func numberOfSubarrays1_1(nums []int, k int) int {
	curSum := 0
	cntMap := make(map[int]int)
	cntMap[curSum] = 1
	res := 0
	for _, num := range nums {
		curSum += num & 1
		res += cntMap[curSum-k]
		cntMap[curSum]++
	}
	return res
}
