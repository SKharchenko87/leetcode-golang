package p1590

func minSubarray(nums []int, p int) int {
	sum := 0
	mp := map[int]int{}
	for _, v := range nums {
		sum += v
	}
	rem := sum % p

	if rem == 0 {
		return 0
	}
	minCount := len(nums)
	sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		tempRem := sum % p
		if tempRem == rem {
			minCount = min(minCount, i+1)
		}
		k := tempRem - rem
		if _, ok := mp[k]; ok {
			minCount = min(minCount, i-mp[k])
		}

		if _, ok := mp[k+p]; ok {
			minCount = min(minCount, i-mp[k+p])
		}
		mp[tempRem] = i
	}

	if minCount >= len(nums) {
		return -1
	}

	return minCount
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func reverse(nums *[]int) {
	for i, j := 0, len(*nums)-1; i < j; i, j = i+1, j-1 {
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
	}
}

// LTE
func minSubarray0(nums []int, p int) int {
	l := len(nums)
	nums = append(nums, nums[l-1])
	prev := 0
	for i := 1; i <= l; i++ {
		tmp := nums[i]
		nums[i] = prev + nums[i-1]
		prev = tmp
	}
	nums[0] = 0
	l++

	sum := nums[l-1]
	if sum%p == 0 {
		return 0
	}
	if sum < p {
		return -1
	}
	lenSubarray := l - 1
	for i := 0; i < l; i++ {
		for j := 0; j+i < l && j < lenSubarray; j++ {
			if (sum-nums[j+i]+nums[i])%p == 0 {
				lenSubarray = j
				break
			}
		}
	}
	if lenSubarray == l-1 {
		return -1
	}
	return lenSubarray
}
