package p2780

func minimumIndex(nums []int) int {
	m := len(nums)
	dominantValue := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			dominantValue = num
			count++
		} else if dominantValue == num {
			count++
		} else {
			count--
		}
	}

	dominantCount := 0
	for _, num := range nums {
		if dominantValue == num {
			dominantCount++
		}
	}

	cntLeft := 0
	for i := 0; i < m; i++ {
		if nums[i] == dominantValue {
			cntLeft++
		}
		if 2*cntLeft > (i+1) && cntLeft > 0 && (dominantCount-cntLeft) > (m-i-1)/2 {
			return i
		}
	}
	return -1
}

func minimumIndex0(nums []int) int {
	m := len(nums)
	dominantValue := -1
	dominantCount := 0
	dominantsCandidate := map[int]int{}
	for i := 0; i < m; i++ {
		dominantsCandidate[nums[i]]++
		if dominantsCandidate[nums[i]] > dominantCount {
			dominantCount = dominantsCandidate[nums[i]]
			dominantValue = nums[i]
		}
	}
	cntLeft := 0
	for i := 0; i < m; i++ {
		if nums[i] == dominantValue {
			cntLeft++
		}
		if 2*cntLeft > (i+1) && cntLeft > 0 && (dominantCount-cntLeft) > (m-i-1)/2 {
			return i
		}
	}
	return -1
}
