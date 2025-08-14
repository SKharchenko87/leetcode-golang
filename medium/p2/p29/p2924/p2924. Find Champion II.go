package p2924

func findChampion(n int, edges [][]int) int {
	looses := make([]bool, n)

	for _, edge := range edges {
		looses[edge[1]] = true
	}

	res := -1
	for i, loos := range looses {
		if !loos {
			if res != -1 {
				return -1
			}
			res = i
		}
	}

	return res
}
