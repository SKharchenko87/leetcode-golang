package p2515

func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	for i := 0; i < n; i++ {
		if words[(startIndex+i)%n] == target {
			return i
		}
		if words[(startIndex-i+n)%n] == target {
			return i
		}
	}
	return -1
}
