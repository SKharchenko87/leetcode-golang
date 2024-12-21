package p1

func countSubarrays(nums []int) int {
	res := 0
	for i := 2; i < len(nums); i++ {
		if 2*(nums[i-2]+nums[i]) == nums[i-1] {
			res++
		}
	}
	return res
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

func min[T Numeric](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func max[T Numeric](x, y T) T {
	if x > y {
		return x
	}
	return y
}
