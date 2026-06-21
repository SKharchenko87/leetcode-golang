package p1833

const limit = 100001

func maxIceCream(costs []int, coins int) int {
	countingSort := [limit]int{}
	for i := 0; i < len(costs); i++ {
		countingSort[costs[i]]++
	}
	ans := 0
	for i := 1; i < limit; i++ {
		if coins >= countingSort[i]*i {
			coins -= countingSort[i] * i
			ans += countingSort[i]
		} else {
			ans += coins / i
			return ans
		}
	}
	return ans
}
