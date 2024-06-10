package p0523

// 13
// 23,  2,  6,  4,  7
// 23, 25, 31, 35, 42
// 10, 12,  5,  9,  3 sum%k

// 6
// 23,  2,  6,  4,  7
// 23, 25, 31, 35, 42
//  5,  1,  1,  5,  0 sum%k
//      _   _         subarray<2

// 6
// 23,  2,  4,  6,  7
// 23, 25, 29, 35, 42
//  5,  1,  5,  5,  0 sum%k
//  5,  1,  5,  5,

// 7
// 23,  2,  4,  6,  6
// 23, 25, 29, 35, 41
//  2,  4,  1,  0,  6 sum%k

// 3
// 5, 0, 0, 0
// 5, 5, 5, 5
// 2, 2, 2, 2

func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	m := make(map[int]int, l)
	m[nums[0]%k] = 0
	for i := 1; i < l; i++ {
		nums[i] += nums[i-1]
		if v, ok := m[nums[i]%k]; nums[i]%k == 0 || ok && i > v+1 {
			return true
		} else if !ok {
			m[nums[i]%k] = i
		}
	}
	return false
}

func checkSubarraySum1(nums []int, k int) bool {
	l := len(nums)
	m := make(map[int]int, l)
	sum := 0
	for i, num := range nums {
		sum += num
		if v, ok := m[sum%k]; sum%k == 0 && i > 0 || ok && i > v+1 {
			return true
		} else if !ok {
			m[sum%k] = i
		}
	}
	return false
}

func checkSubarraySum0(nums []int, k int) bool {
	l := len(nums)
	m := make(map[int]int, l-1)
	nums[0] %= k
	prev := nums[0]

	for i := 1; i < l; i++ {
		nums[i] %= k
		prev = (prev + nums[i]) % k
		m[i] = prev
		if m[i]%k == 0 {
			return true
		}
	}
	for i := 2; i < l; i++ {
		for j := i; j < l; j++ {
			m[j] = (m[j] - nums[i-2]) % k
			if m[j]%k == 0 {
				println(j)
				return true
			}
		}
	}
	return false
}
