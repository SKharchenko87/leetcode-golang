package p0514

func findRotateSteps(ring string, key string) int {
	m := map[[2]int]int{}
	length := len(ring)
	var bfs func(index int, indexKey int) int
	var findLetter = func(direction string, index int, indexKey int) int {
		res := 1
		for ring[index] != key[indexKey] {
			switch direction {
			case "left":
				index--
				if index == -1 {
					index = length - 1
				}
			case "right":
				index++
				index %= length
			}
			res++
		}
		if _, ok := m[[2]int{index, indexKey + 1}]; !ok {
			m[[2]int{index, indexKey + 1}] = bfs(index, indexKey+1)
		}
		return res + m[[2]int{index, indexKey + 1}]
	}
	bfs = func(index int, indexKey int) int {
		if len(key) == indexKey {
			return 0
		}
		left := findLetter("left", index, indexKey)
		right := findLetter("right", index, indexKey)
		return min(left, right)
	}
	return bfs(0, 0)
}
