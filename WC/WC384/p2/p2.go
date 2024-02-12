package p2

//[0 1 4]
//[1 0 -1 1 1 0 -1]
// 1 0 0  1 1 0 0
//   2 0  0  2 0
//     3  0 0 0 3

func countMatchingSubarrays(nums []int, pattern []int) int {
	l := len(nums)
	hp := 0
	for i := 0; i < len(pattern); i++ {
		hp += pattern[i]
	}
	arr := make([]int, l-1)
	arrh := []int{}
	sum := 0
	for i := 1; i < l; i++ {
		if nums[i-1] < nums[i] {
			arr[i-1] = 1
		} else if nums[i-1] == nums[i] {
			arr[i-1] = 0
		} else if nums[i-1] > nums[i] {
			arr[i-1] = -1
		}
		sum += arr[i-1]
		if i >= len(pattern) {
			if sum == hp && arr[i-1] == pattern[len(pattern)-1] && arr[i-2] == pattern[len(pattern)-2] {
				arrh = append(arrh, i-len(pattern))
			}
			sum -= arr[i-len(pattern)]
		}
	}
	cnt := 0
	for _, i := range arrh {
		c := 1
		for j := 0; j < len(pattern); j++ {
			if arr[i+j] != pattern[j] {
				c = 0
				break
			}
		}
		cnt += c
	}
	return cnt
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
