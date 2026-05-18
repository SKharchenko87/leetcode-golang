package p1306

func canReach(arr []int, start int) bool {
	n := len(arr)
	visited := make([]bool, n)
	tmp := make([]int, 1, n)
	tmp[0] = start
	i := 0
	for {
		prev := len(tmp)
		for ; i < len(tmp); i++ {
			index := tmp[i]
			l := index - arr[index]
			r := index + arr[index]
			if l >= 0 && !visited[l] {
				if arr[l] == 0 {
					return true
				}
				visited[l] = true
				tmp = append(tmp, l)
			}
			if n > r && !visited[r] {
				if arr[r] == 0 {
					return true
				}
				visited[r] = true
				tmp = append(tmp, r)
			}
		}
		if len(tmp) == prev {
			return false
		}
	}
}
