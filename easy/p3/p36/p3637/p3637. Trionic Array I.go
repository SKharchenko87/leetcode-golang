package p3637

const increasing = true
const decreasing = false

func isTrionic(nums []int) bool {
	n := len(nums)
	if n < 4 || nums[0] >= nums[1] {
		return false
	}

	isUp := true
	cnt := 0
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			return false
		}
		if (isUp && nums[i] >= nums[i+1]) || !isUp && nums[i] <= nums[i+1] {
			if cnt == 2 {
				return false
			}
			cnt++
			isUp = !isUp
		}
	}
	return cnt == 2
}

func isTrionic0(nums []int) bool {
	n := len(nums)
	if n < 4 {
		return false
	}

	var p, q int
	for p = 1; p < n && nums[p] > nums[p-1]; p++ {
	}
	if p > n-2 || p == 1 {
		return false
	}
	for q = p; q < n && nums[q] < nums[q-1]; q++ {
	}
	if q > n-1 {
		return false
	}
	//return slices.IsSorted(nums[q:])

	for i := q; i < n; i++ {
		if nums[i] <= nums[i-1] {
			return false
		}
	}

	return true
}
