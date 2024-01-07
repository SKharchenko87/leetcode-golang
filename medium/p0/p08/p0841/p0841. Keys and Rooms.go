package p0841

func canVisitAllRooms(rooms [][]int) bool {
	m := map[int]int{}
	m[0]++
	flg := true
	for flg {
		flg = false
		for k, v := range m {
			if v == 1 {
				flg = true
				for i := 0; i < len(rooms[k]); i++ {
					if _, ok := m[rooms[k][i]]; !ok {
						m[rooms[k][i]]++
					}
				}
				m[k]++
			}
		}
	}
	return len(m) == len(rooms)
}
