package p0781

func numRabbits(answers []int) int {
	m := make(map[int]int)
	for i := 0; i < len(answers); i++ {
		m[answers[i]]++
	}
	result := 0
	for k, v := range m {
		if k == 0 {
			result += v
			continue
		}
		if k == 1 {
			result += v/(k+1)*(k+1) + v%(k+1)*(k+1)
			continue
		}
		result += k + 1
		if v > k+1 {
			result += v / (k + 1) * (k + 1)
		}
	}
	return result
}
