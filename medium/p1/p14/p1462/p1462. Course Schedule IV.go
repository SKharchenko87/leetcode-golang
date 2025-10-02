package p1462

func setBit(arr *[4]int, a int) {
	index := a / 32
	bitPos := a % 32
	arr[index] |= (1 << bitPos)
}

func checkBit(arr *[4]int, a int) bool {
	index := a / 32
	bitPos := a % 32
	return (arr[index] & (1 << bitPos)) != 0
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {

	g := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		g[pre[0]] = append(g[pre[0]], pre[1])
		indegree[pre[1]] += 1
	}

	q := []int{}
	for i, val := range indegree {
		if val == 0 {
			q = append(q, i)
		}
	}

	all := make([][4]int, numCourses)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, next := range g[curr] {

			all[next][0] |= all[curr][0]
			all[next][1] |= all[curr][1]
			all[next][2] |= all[curr][2]
			all[next][3] |= all[curr][3]
			setBit(&all[next], curr)

			indegree[next] -= 1
			if indegree[next] == 0 {
				q = append(q, next)
			}
		}

	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = checkBit(&all[q[1]], q[0])
	}
	return ans
}
