package p2461

func maximumSubarraySum(nums []int, k int) int64 {

	m := [100001]int{}
	cnt2 := 0
	var sum, res int64
	for i := 0; i < k; i++ {
		sum += int64(nums[i])
		m[nums[i]]++
		if m[nums[i]] == 2 {
			cnt2++
		}
	}

	if cnt2 == 0 && res < sum {
		res = sum
	}
	for i := k; i < len(nums); i++ {
		sum += int64(nums[i] - nums[i-k])
		m[nums[i]]++
		if m[nums[i]] == 2 {
			cnt2++
		}

		m[nums[i-k]]--
		if m[nums[i-k]] == 1 {
			cnt2--
		}

		if cnt2 == 0 && res < sum {
			res = sum
		}
	}
	return res
}

func maximumSubarraySum0(nums []int, k int) int64 {
	m1, m2 := make(map[int]bool), make(map[int]int)
	var sum, res int64
	for i := 0; i < k; i++ {
		sum += int64(nums[i])
		if _, ok := m1[nums[i]]; ok {
			m2[nums[i]]++
		} else {
			m1[nums[i]] = true
		}
	}

	if len(m2) == 0 && res < sum {
		res = sum
	}
	for i := k; i < len(nums); i++ {
		sum += int64(nums[i] - nums[i-k])
		if _, ok := m1[nums[i]]; ok {
			m2[nums[i]]++
		} else {
			m1[nums[i]] = true
		}

		if _, ok := m2[nums[i-k]]; !ok {
			delete(m1, nums[i-k])
		} else {
			m2[nums[i-k]]--
			if m2[nums[i-k]] == 0 {
				delete(m2, nums[i-k])
			}
		}

		if len(m2) == 0 && res < sum {
			res = sum
		}
	}
	return res
}
