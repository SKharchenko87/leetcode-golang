package p3477

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(fruits)
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if fruits[i] <= baskets[j] {
				count++
				baskets[j] = 0
				break
			}
		}
	}
	return n - count
}
