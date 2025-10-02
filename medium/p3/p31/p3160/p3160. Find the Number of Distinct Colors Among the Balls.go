package p3160

func queryResults(limit int, queries [][]int) []int {
	res := make([]int, len(queries))
	colorCount := make(map[int]int)
	ballColors := make(map[int]int)
	for i, query := range queries {
		index := query[0]
		color := query[1]
		if prevColor, ok := ballColors[index]; ok {
			colorCount[prevColor]--
			if colorCount[prevColor] == 0 {
				delete(colorCount, prevColor)
			}
		}
		ballColors[index] = color
		colorCount[color]++
		res[i] = len(colorCount)
	}
	return res
}

func queryResults0(limit int, queries [][]int) []int {
	balls := make([]int, limit+1)
	res := make([]int, len(queries))
	freq := make(map[int]int)
	cnt := 0
	for i, query := range queries {
		index := query[0]
		color := query[1]
		if balls[index] != color {
			freq[balls[index]]--
			if freq[balls[index]] == 0 {
				cnt--
			}
			freq[color]++
			if freq[color] == 1 {
				cnt++
			}
		}

		res[i] = cnt
		balls[index] = color
	}
	return res
}
