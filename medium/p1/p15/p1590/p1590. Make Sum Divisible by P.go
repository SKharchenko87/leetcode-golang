package p1590

func minSubarray(nums []int, p int) int {
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		sum = (sum + nums[i]) % p
	}

	if sum == 0 {
		return 0
	}
	m := make(map[int]int, l)
	m[0] = -1
	res := l
	cur := 0
	for i := 0; i < l; i++ {
		if sum-nums[i]%p == 0 {
			return 1
		}
		cur = (cur + nums[i]) % p
		if val, ok := m[(p+cur-sum)%p]; ok {
			res = min(res, i-val)
		}
		m[cur] = i
	}

	if res == l {
		return -1
	}
	return res

}

/*TLE*/
func minSubarray2(nums []int, p int) int {
	l := len(nums)
	prefSum := make([]int, l+2)
	sum := 0
	for i, cur := range nums {
		sum = (sum + cur) % p
		prefSum[i+1] = sum
	}
	if sum%p == 0 {
		return 0
	}
	res := l
	for i := 1; i <= l; i++ {
		for j := i + 1; j <= l+1; j++ {
			prev := prefSum[i-1]
			cur := (p + prefSum[j-1] - prev) % p
			if (sum-cur+p)%p == 0 {
				if res > j-i {
					res = j - i
				}
			}
		}
	}
	if res == l {
		return -1
	}
	return res

}

func minSubarray1(nums []int, p int) int {
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
