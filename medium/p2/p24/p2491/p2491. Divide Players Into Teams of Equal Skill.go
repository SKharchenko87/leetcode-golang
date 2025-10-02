package p2491

func dividePlayers(skill []int) int64 {
	l := len(skill)
	m := make(map[int]int, l)
	sum := 0
	for i := 0; i < l; i++ {
		sum += skill[i]
		m[skill[i]]++
	}
	if sum%(l/2) != 0 {
		return -1
	}
	x := sum / (l / 2)
	var res int64
	for k, _ := range m {
		for m[k] > 0 {
			if _, ok := m[x-k]; !ok {
				return -1
			}
			res += int64(x-k) * int64(k)
			m[x-k]--
			if m[x-k] == -1 {
				return -1
			}
			m[k]--
			if m[k] == -1 {
				return -1
			}
		}
	}
	return res
}
