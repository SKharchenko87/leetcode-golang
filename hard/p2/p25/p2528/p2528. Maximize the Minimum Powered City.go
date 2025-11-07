package p2528

import "math"

var arr = make([]int, 100001)

func check(stations []int, r int, k int, x int) bool {
	l := len(stations)
	clear(arr)
	var leftSum, rightSum int
	for i := 0; i < r && i < l; i++ {
		rightSum += stations[i]
	}
	for i := 0; i < l; i++ {
		if i-r-1 >= 0 {
			leftSum -= arr[i-r-1]
			leftSum -= stations[i-r-1]
		}
		if i+r < l {
			rightSum += arr[i+r]
			rightSum += stations[i+r]
		}

		sum := leftSum + rightSum
		if sum < x {
			diff := x - sum
			k -= diff
			if k < 0 {
				return false
			}
			index := min(l-1, i+r)
			arr[index] += diff
			rightSum += diff
		}

		rightSum -= arr[i]
		rightSum -= stations[i]

		leftSum += arr[i]
		leftSum += stations[i]

	}
	return true
}

func maxPower(stations []int, r int, k int) int64 {
	var low, high int
	low, high = 0, math.MaxInt

	var result int64
	for low <= high {
		mid := low + (high-low)/2
		if check(stations, r, k, mid) {
			result = int64(mid)
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}
