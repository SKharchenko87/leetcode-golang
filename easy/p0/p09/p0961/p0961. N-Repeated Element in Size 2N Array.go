package p0961

func repeatedNTimes(nums []int) int {
	n := len(nums) / 2
	m := [10001]bool{}
	for i := 0; i < n+2; i++ {
		if m[nums[i]] {
			return nums[i]
		}
		m[nums[i]] = true
	}
	return 0
}

var mGlobal = make(map[int]struct{}, 10001)

func repeatedNTimes2(nums []int) int {
	clear(mGlobal)
	n := len(nums) / 2
	for i := 0; i < n+2; i++ {
		if _, ok := mGlobal[nums[i]]; ok {
			return nums[i]
		}
		mGlobal[nums[i]] = struct{}{}
	}
	return 0
}

var sliceGlobal = make([]bool, 10001)

func repeatedNTimes1(nums []int) int {
	clear(sliceGlobal)
	n := len(nums) / 2
	for i := 0; i < n+2; i++ {
		if sliceGlobal[nums[i]] {
			return nums[i]
		}
		sliceGlobal[nums[i]] = true
	}
	return 0
}

func repeatedNTimes0(nums []int) int {
	n := len(nums) / 2
	m := make(map[int]struct{}, n)
	for i := 0; i < n+2; i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = struct{}{}
		} else {
			return nums[i]
		}
	}
	return 0
}
