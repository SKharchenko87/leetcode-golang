package p4

func canMakePalindromeQueries(s string, queries [][]int) []bool {
	m := map[int32][]int{'a': []int{}, 'b': []int{}, 'c': []int{}, 'd': []int{}, 'e': []int{}, 'f': []int{}, 'g': []int{}, 'h': []int{}, 'i': []int{}, 'j': []int{}, 'k': []int{}, 'l': []int{}, 'm': []int{}, 'n': []int{}, 'o': []int{}, 'p': []int{}, 'q': []int{}, 'r': []int{}, 's': []int{}, 't': []int{}, 'u': []int{}, 'v': []int{}, 'w': []int{}, 'x': []int{}, 'y': []int{}, 'z': []int{}}
	for i, c := range s {
		m[c] = append(m[c], i)
	}
	cnt := 0
	for _, v := range m {
		if len(v)%2 == 1 {
			cnt++
			if cnt > 1 {
				return []bool{false}
			}
		}
	}
	return []bool{false}
}
