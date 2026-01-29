package p2976

const limit = 1e9 + 1

var costMatrix = [26][26]int{}

func clearMatrix() {
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			costMatrix[i][j] = limit
		}
	}
	for i := 0; i < 26; i++ {
		costMatrix[i][i] = 0
	}
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	clearMatrix()

	for i := 0; i < len(original); i++ {
		original[i] -= 'a'
		changed[i] -= 'a'
		o, t := original[i], changed[i]
		costMatrix[o][t] = min(costMatrix[o][t], cost[i])
	}
	newOriginal := make([]byte, 0, 26*26)
	newChanged := make([]byte, 0, 26*26)

	for len(original) > 0 {
		for i := 0; i < len(original); i++ {
			o, t := original[i], changed[i]
			for newT := byte(0); newT < 26; newT++ {
				candidateCost := costMatrix[o][t] + costMatrix[t][newT]
				if candidateCost < costMatrix[o][newT] {
					newOriginal, newChanged = append(newOriginal, o), append(newChanged, newT)
					costMatrix[o][newT] = candidateCost
				}
			}
		}
		original, newOriginal = newOriginal, original[:0]
		changed, newChanged = newChanged, changed[:0]
	}

	var res, v int64
	for i := range source {
		if v = int64(costMatrix[source[i]-'a'][target[i]-'a']); v == limit {
			return -1
		}
		res += v
	}
	return res
}

//'a', 'b', 't', 't', 'e', 'd'}, changed: []byte{'b', 't', 'b', 'e', 'b', 'e'
// a->b->t->e	0->1->2->4
// b->t->e 		1->2->4
// t->e t->b 	2->4 2->1
// e->b->t		4->1->2
// d->e			3->4

func minimumCost1(source string, target string, original []byte, changed []byte, cost []int) int64 {
	var res int64
	matrix := findMinPath(original, changed, cost)
	for i := 0; i < len(source); i++ {
		if source[i] != target[i] {
			if matrix[source[i]-'a'][target[i]-'a'] == 0 {
				return -1
			}
			res += int64(matrix[source[i]-'a'][target[i]-'a'])
		}
	}
	return res
}

func findMinPath(original []byte, changed []byte, cost []int) (res [26][26]int) {

	graph := map[byte]map[byte]int{}
	for i, b := range original {
		if res[b-'a'][changed[i]-'a'] > cost[i] || res[b-'a'][changed[i]-'a'] == 0 {
			res[b-'a'][changed[i]-'a'] = cost[i]
			if _, ok := graph[b-'a']; !ok {
				graph[b-'a'] = map[byte]int{}
			}
			graph[b-'a'][changed[i]-'a'] = cost[i]
		}
	}

	for b, m := range graph {
		for len(m) > 0 {
			newM := map[byte]int{}
			for b2, i := range m {
				for b3, i2 := range graph[b2] {
					if res[b][b3] > i+i2 || res[b][b3] == 0 {
						res[b][b3] = i + i2
						newM[b3] = i + i2
					}
				}
			}
			m = newM
		}
	}
	return res
}
