package p0974

//  4,  5,  0, -2, -3,  1    5: 1     0: 1
//    4,  0,  0,  -2, -3, 1
//    2    3     2  1  1
// 5
//  4,  5,  0, -2, -3,  1    5: 1     0: 1
//  4,  9,  9,  7,  4,  5 sum
//  4,  4,  4,  2,  4,  0    4: 1+1+2 0: 1
//      _
//          _
//      _   _
//                  _
//      _           _
//          _       _
//                     _
//
///////////
// 5,  5,  5
// 5, 10, 15
// 0,  0,  0   0: 1+2     +3

//  3
//  -1,  3,  7
//  -1   0,  1
//  -1,  2,  9
//           _
//       _

//  0,  1,  7, 0,  1
//    2   3   2
//
//
//0, 0, 0+0, 1+7, 7+1, 0+1+7, 1+7+0
//    0,  1,  0, 0, 1

// 8
//  8,  9,  7,  8,  9
//  8, 17, 24, 32, 41
//  0,  1,  0,  0,  1
//  _
//  _   _   _
//  _   _   _   _
//      _   _
//      _   _   _
//              _
//      _   _   _   _

// 7
//  7,  -5,  5,  -8,  -6,  6,  -4,  7,  -8,  -7
//  7,   2,  7,  -1,  -6,  0,  -4,  3,  -5,  -12
//  0,   2,  0,  -1,  -6,  0,  -4,  3,  -5,  -5

// [7]
// [7,  -5,  5]
// [-5,  5]
// [-8,  -6]?
// []
// []
// []
// []
// []
// []
// []

func subarraysDivByKXX(nums []int, k int) int {
	prefixMod := 0
	result := 0

	// There are k mod groups 0...k-1.
	modGroups := make([]int, k)
	modGroups[0] = 1

	for _, num := range nums {
		// Take modulo twice to avoid negative remainders.
		prefixMod = (prefixMod + (num%k+k)%k) % k
		// Add the count of subarrays that have the same remainder as the current
		// one to cancel out the remainders.
		result += modGroups[prefixMod]
		modGroups[prefixMod]++
	}

	return result
}

func subarraysDivByK(nums []int, k int) int {
	l := len(nums)
	arrK := make([]int, k)
	arrK[0] = 1
	nums[0] = (nums[0]%k + k) % k
	cnt := arrK[nums[0]]
	arrK[nums[0]]++
	for i := 1; i < l; i++ {
		nums[i] = (nums[i-1] + (nums[i]%k+k)%k) % k
		cnt += arrK[nums[i]]
		arrK[nums[i]]++
	}
	return cnt
}

func subarraysDivByKX(nums []int, k int) int {
	l := len(nums)
	m := make(map[int]int, l)
	cnt := 0
	if nums[0]%k == 0 {
		cnt++
	}
	m[nums[0]%k]++
	for i := 1; i < l; i++ {
		nums[i] += nums[i-1]
		//if nums[i]%k == 0 || nums[i]%k > 0 && m[nums[i]%k-k] > 0 || nums[i]%k < 0 && m[k+nums[i]%k] > 0 {
		if nums[i]%k == 0 || m[nums[i]%k-k] > 0 || m[k+nums[i]%k] > 0 {
			cnt++
		}
		m[nums[i]%k]++
		cnt += m[nums[i]%k] - 1
	}
	return cnt
}

//
//
//
//

func subarraysDivByK1(nums []int, k int) int {
	l := len(nums)
	sumArr := make([]int, l)
	cnt := 0
	sumArr[0] = nums[0]
	for i := 1; i < l; i++ {
		sumArr[i] += nums[i] + sumArr[i-1]
	}
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if sumArr[j]%k == 0 {
				cnt++
			}
			sumArr[j] -= nums[i]
		}
	}
	return cnt
}

func subarraysDivByK0(nums []int, k int) int {
	l := len(nums)
	cnt := 0
	for i := 1; i < l; i++ {
		nums[i] += nums[i-1]
	}
	for i := 0; i < l; i++ {
		if nums[i]%k == 0 {
			cnt++
		}
		for j := i + 1; j < l; j++ {
			if (nums[j]-nums[i])%k == 0 {
				cnt++
			}
		}
	}
	return cnt
}

//func subarraysDivByK(nums []int, k int) int {
//	l := len(nums)
//	m := make(map[int]int, l)
//	cnt := 0
//	m[nums[0]%k] = 0
//	if nums[0]%k == 0 {
//		cnt++
//	}
//	for i := 1; i < l; i++ {
//		c := nums[i]
//		if c%k == 0 {
//			cnt++
//		}
//		nums[i] += nums[i-1]
//		cur := nums[i] % k
//		if cur == 0 {
//			cnt++
//		}
//		if _, ok := m[cur]; ok {
//			if c%k == 0 {
//				cnt += m[cur]
//				m[cur]++
//			} else {
//				m[cur]++
//				cnt += m[cur]
//			}
//		} else {
//			m[cur] = 0
//		}
//	}
//	return cnt
//}
