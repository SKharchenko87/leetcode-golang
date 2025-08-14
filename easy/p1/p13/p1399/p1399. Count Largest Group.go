package p1399

var results = [10001]int{}

func countLargestGroup(n int) int {
	return results[n]
}

func init() {
	maxCount := 0
	cnt := 0
	m := [37]int{}
	for i := 1; i <= 10000; i++ {
		sum := sumDigits(i)
		m[sum]++
		if m[sum] > maxCount {
			maxCount = m[sum]
			cnt = 1
		} else if m[sum] == maxCount {
			cnt++
		}
		results[i] = cnt
	}
}

func countLargestGroup0(n int) int {
	maxCount := 0
	cnt := 0
	m := make(map[int]int, 40)
	for i := 1; i <= n; i++ {
		sum := sumDigits(i)
		m[sum]++
		if m[sum] > maxCount {
			maxCount = m[sum]
			cnt = 1
		} else if m[sum] == maxCount {
			cnt++
		}
	}
	return cnt
}

func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
