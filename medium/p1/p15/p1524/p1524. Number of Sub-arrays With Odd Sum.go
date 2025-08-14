package p1524

const mod = 1e9 + 7

func numOfSubarrays(arr []int) int {
	even := 1
	odd := 0
	prefixSum := 0
	result := 0

	for _, num := range arr {
		prefixSum += num
		if prefixSum%2 == 0 {
			result = (result + odd) % mod
			even++
		} else {
			result = (result + even) % mod
			odd++
		}
	}

	return result

}
