package p2392

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	n := len(rowConditions)
	m := len(colConditions)
	matrix := make([][]int, k)

	sortRow := sortPosition(rowConditions, k, n)
	if sortRow == nil || len(sortRow) < k {
		return [][]int{}
	}

	sortCol := sortPosition(colConditions, k, m)
	if sortCol == nil || len(sortCol) < k {
		return [][]int{}
	}

	for i := 0; i < k; i++ {
		matrix[sortRow[i+1]] = make([]int, k)
		matrix[sortRow[i+1]][sortCol[i+1]] = i + 1
	}

	return matrix
}

func parseConditions(conditions [][]int, parentsCount *[]int, parentToChildren *map[int][]int) {
	for _, condition := range conditions {
		(*parentsCount)[condition[1]-1]++
		if (*parentToChildren)[condition[0]] == nil {
			(*parentToChildren)[condition[0]] = []int{}
		}
		(*parentToChildren)[condition[0]] = append((*parentToChildren)[condition[0]], condition[1])
	}
}

// {{1, 4}, {4, 3}, {1, 2}, {1, 4}}
func sortPosition(conditions [][]int, k, l int) map[int]int {
	positions := map[int]int{}
	index := 0
	parentsCount := make([]int, k)
	parentToChildren := make(map[int][]int, l)

	//childParent := make(map[int]map[int]bool, l)

	parseConditions(conditions, &parentsCount, &parentToChildren)
	parents := map[int]bool{}

	for i := 0; i < k; i++ {
		if parentsCount[i] == 0 && !parents[i+1] {
			//parents[i+1] = true
			positions[i+1] = index
			index++
			for _, children := range parentToChildren[i+1] {
				parentsCount[children-1]--
				if parentsCount[children-1] == 0 {
					parents[children] = true
				}
			}
		}
	}
	//, {4, 3}, {5, 8}, {8, 2}, {5, 8}, {3, 2}, {4, 3}, {4, 8}, },
	for len(parents) > 0 {
		newParent := map[int]bool{}
		for parent, _ := range parents {
			positions[parent] = index
			index++
			for _, children := range parentToChildren[parent] {
				parentsCount[children-1]--
				if parentsCount[children-1] == 0 {
					newParent[children] = true
				}
			}
		}
		parents = newParent
	}

	return positions
}

func buildMatrix0(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	n := len(rowConditions)
	m := len(colConditions)
	matrix := make([][]int, k)

	sortRow := sortPosition0(rowConditions, k, n)
	if sortRow == nil || len(sortRow) < k {
		return [][]int{}
	}

	sortCol := sortPosition0(colConditions, k, m)
	if sortCol == nil || len(sortCol) < k {
		return [][]int{}
	}

	for i := 0; i < k; i++ {
		matrix[sortRow[i+1]] = make([]int, k)
		matrix[sortRow[i+1]][sortCol[i+1]] = i + 1
	}

	return matrix
}

func parseConditions0(conditions [][]int, parentChildren *map[int][]int, childParent *map[int]map[int]bool) {
	for _, condition := range conditions {
		if (*parentChildren)[condition[0]] == nil {
			(*parentChildren)[condition[0]] = []int{}
		}
		(*parentChildren)[condition[0]] = append((*parentChildren)[condition[0]], condition[1])
		if (*childParent)[condition[1]] == nil {
			(*childParent)[condition[1]] = map[int]bool{}
		}
		(*childParent)[condition[1]][condition[0]] = true
	}
}

func sortPosition0(conditions [][]int, k, l int) map[int]int {
	positions := map[int]int{}
	index := 0
	parentChildren := make(map[int][]int, l)
	childParent := make(map[int]map[int]bool, l)

	parseConditions0(conditions, &parentChildren, &childParent)
	parents := map[int]bool{}
	for i := 0; i < k; i++ {
		if _, ok := childParent[i+1]; !ok {
			parents[i+1] = true
		}
	}

	for len(parents) > 0 {
		newParent := map[int]bool{}
		for parent, _ := range parents {
			if _, ok := childParent[parent]; !ok {
				positions[parent] = index
				index++
				for _, child := range parentChildren[parent] {
					delete(childParent[child], parent)
					if len(childParent[child]) == 0 {
						newParent[child] = true
						delete(childParent, child)
					}
				}
			}
		}
		parents = newParent
	}
	return positions
}
