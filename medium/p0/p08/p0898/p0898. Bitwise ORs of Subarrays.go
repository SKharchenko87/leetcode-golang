package p0898

func subarrayBitwiseORs(arr []int) int {
	l := len(arr)

	m := make(map[int]struct{}, l/2)
	prev := make(map[int]struct{})
	curr := make(map[int]struct{})
	m[arr[0]] = struct{}{}
	curr[arr[0]] = struct{}{}
	for i := 1; i < l; i++ {
		prev, curr = curr, prev
		clear(curr)
		curr[arr[i]] = struct{}{}
		m[arr[i]] = struct{}{}
		for old := range prev {
			v := old | arr[i]
			if _, ok := curr[v]; !ok {
				curr[v] = struct{}{}
				m[v] = struct{}{}
			}
		}
	}
	return len(m)
}

func subarrayBitwiseORs1(arr []int) int {
	l := len(arr)
	prev := map[int]struct{}{}
	m := map[int]struct{}{arr[0]: {}}
	curr := map[int]struct{}{arr[0]: {}}
	for i := 1; i < l; i++ {
		prev, curr = curr, prev
		clear(curr)
		curr[arr[i]] = struct{}{}
		m[arr[i]] = struct{}{}
		for old := range prev {
			v := old | arr[i]
			curr[v] = struct{}{}
			m[v] = struct{}{}
		}
	}
	return len(m)
}

func subarrayBitwiseORs0(arr []int) int {
	l := len(arr)
	prev := map[int]struct{}{}
	m := map[int]struct{}{arr[0]: {}}
	curr := map[int]struct{}{arr[0]: {}}
	for i := 1; i < l; i++ {
		prev, curr = curr, prev
		clear(curr)
		for old := range prev {
			v := old | arr[i]
			curr[v] = struct{}{}
			m[v] = struct{}{}
		}
		curr[arr[i]] = struct{}{}
		m[arr[i]] = struct{}{}
	}
	return len(m)
}

// TLE
func subarrayBitwiseORsTLE(arr []int) int {
	l := len(arr)
	m := make(map[int]struct{})
	for i := 0; i < l; i++ {
		curr := arr[i]
		m[curr] = struct{}{}
		for j := i + 1; j < l; j++ {
			curr |= arr[j]
			m[curr] = struct{}{}
		}
	}
	return len(m)
}
