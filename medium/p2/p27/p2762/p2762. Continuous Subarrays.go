package p2762

func continuousSubarrays(nums []int) int64 {
	var ans int64
	maxNum := nums[0]
	minNum := nums[0]

	for left, right := 0, 0; right < len(nums); right++ {
		if abs(nums[right]-maxNum) <= 2 && abs(nums[right]-minNum) <= 2 {
			ans += int64(right - left + 1)
			maxNum = max(maxNum, nums[right])
			minNum = min(minNum, nums[right])
			continue
		}

		maxNum = nums[right]
		minNum = nums[right]
		for j := right; j >= left; j-- {
			if abs(nums[right]-nums[j]) > 2 {
				left = j + 1
				break
			}
			maxNum = max(maxNum, nums[j])
			minNum = min(minNum, nums[j])
		}
		ans += int64(right - left + 1)
	}
	return ans
}

type Ints interface {
	int | int8 | int16 | int32 | int64
}
type Floats interface {
	float32 | float64
}
type Numeric interface{ Ints | Floats }

func abs[T Numeric](x T) T {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func max[T Numeric](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func min[T Numeric](x, y T) T {
	if x < y {
		return x
	}
	return y
}

/*

67,67,66,67,68,69,70,71,71
 s             e           =5
    s          e           =4
	   s       e           =3
	      s        e       =3
		     s        e    =3
			   s          e=4
			       s      e=3
				      s   e=2
					     se=1

*/
