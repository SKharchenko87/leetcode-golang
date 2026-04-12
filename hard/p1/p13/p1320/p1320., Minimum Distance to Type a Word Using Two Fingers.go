package p1320

var matrixCost [27][26]int

func init() {
	for fingerPos := 0; fingerPos < 26; fingerPos++ {
		for targetPos := 0; targetPos < 26; targetPos++ {
			x1, y1 := fingerPos/6, fingerPos%6
			x2, y2 := targetPos/6, targetPos%6
			matrixCost[fingerPos][targetPos] = abs(x1-x2) + abs(y1-y2)
		}
	}
}

func minimumDistance(word string) int {
	n := len(word)
	mem := make([][27]int, n)
	for i := range mem {
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var rec func(index int, otherFinger int) int
	rec = func(index int, otherFinger int) int {
		if index == n {
			return 0
		}

		if mem[index][otherFinger] != -1 {
			return mem[index][otherFinger]
		}

		target := int(word[index] - 'A')
		cost1 := 0
		if index > 0 {
			prevChar := int(word[index-1] - 'A')
			cost1 = matrixCost[prevChar][target]
		}
		res1 := cost1 + rec(index+1, otherFinger)

		cost2 := matrixCost[otherFinger][target]
		newOtherFinger := 26
		if index > 0 {
			newOtherFinger = int(word[index-1] - 'A')
		}
		res2 := cost2 + rec(index+1, newOtherFinger)

		mem[index][otherFinger] = min(res1, res2)
		return mem[index][otherFinger]
	}

	return rec(0, 26)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
