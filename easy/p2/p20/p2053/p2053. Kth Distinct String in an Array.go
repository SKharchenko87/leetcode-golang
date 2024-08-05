package p2053

func kthDistinct(arr []string, k int) string {
	m := make(map[string]int)
	l := len(arr)
	for i := 0; i < l; i++ {
		m[arr[i]]++
	}

	for i := 0; k > 0 && i < l; i++ {
		if m[arr[i]] == 1 {
			k--
		}
		if k == 0 {
			return arr[i]
		}
	}
	return ""
}
