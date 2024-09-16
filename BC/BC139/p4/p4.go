package p4

import (
	"slices"
	"sort"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func maxPathLength(coordinates [][]int, k int) int {

	x, y := coordinates[k][0], coordinates[k][1]
	l := len(coordinates)
	var (
		orderX = make([][2]int, l)
		orderY = make([][2]int, l)
	)
	for i, coordinate := range coordinates {
		if coordinate[0] < x && coordinate[1] < y || coordinate[0] > x && coordinate[1] > y {
			orderX[i] = [2]int{coordinate[0], coordinate[1]}
			orderY[i] = [2]int{coordinate[0], coordinate[1]}
		}
	}
	orderX = append(orderX, [2]int{x, y})
	sort.Slice(orderX, func(i, j int) bool {
		if orderX[i][0] == orderX[j][0] {
			return orderX[i][1] < orderX[j][1]
		}
		return orderX[i][0] < orderX[j][0]
	})
	orderY = append(orderY, [2]int{x, y})
	sort.Slice(orderY, func(i, j int) bool {
		if orderY[i][1] == orderY[j][1] {
			return orderY[i][0] < orderY[j][0]
		}
		return orderY[i][1] < orderY[j][1]
	})

	candidates := map[[2]int]int{[2]int{-1, -1}: 0}
	newCandidates := map[[2]int]int{}
	visited := make(map[[2]int]int, len(orderX))

	for ints, level := range candidates {
		index, _ := slices.BinarySearchFunc(orderX, ints, func(a [2]int, b [2]int) int {
			if a[0] == b[0] {
				return a[1] - b[1]
			}
			return a[0] - b[0]
		})
		for i := index + 1; i < len(orderX); i++ {
			if orderX[i][0] > ints[0] && orderX[i][1] > ints[1] {
				if newCandidates[orderX[i]] < level+1 {
					newCandidates[orderX[i]] = level + 1
				}
			}
		}
	}

	for ; i < len(orderX); i++ {
		if orderX[i][0] > prevX && orderX[i][0] <= currX && orderX[i][1] > prevY && orderX[i][1] <= currY {
			if visited[orderX[i]] < level {
				visited[orderX[i]] = level
				candidate = append(candidate, orderX[i])
			}
		}
	}

	prevX, prevY := -1, -1
	currX, currY := orderX[0][0], orderY[0][1]
	visited := make(map[[2]int]int, len(orderX))
	candidate := make([][2]int, 0, len(orderX))
	level := 0
	i := 0
	for ; i < len(orderX); i++ {
		if orderX[i][0] > prevX && orderX[i][0] <= currX && orderX[i][1] > prevY && orderX[i][1] <= currY {
			if visited[orderX[i]] < level {
				visited[orderX[i]] = level
				candidate = append(candidate, orderX[i])
			}
		}
		if orderX[i][0] > currX {
			break
		}
	}
	prevX, prevY := currX, currY
	currX, currY := orderX[i][0], orderY[i][1]

	sort.Slice(coordinates, func(i, j int) bool {
		if coordinates[i][0] == coordinates[j][0] {
			return coordinates[i][1] < coordinates[j][1]
		}
		return coordinates[i][0] < coordinates[j][0]
	})
	index, _ := slices.BinarySearchFunc(
		coordinates, []int{x, y},
		func(ints []int, ints2 []int) int {
			if ints[0] == ints2[0] {
				return ints[1] - ints2[1]
			}
			return ints[0] - ints2[0]
		})
	cnt := 1
	curX, curY := x, y
	for i := index - 1; i >= 0; i-- {
		tmp := 0

		if coordinates[i][0] < curX && coordinates[i][1] < curY {
			cnt++
			curX, curY = coordinates[i][0], coordinates[i][1]
		}
	}
	curX, curY = x, y
	for i := index + 1; i < l; i++ {
		if coordinates[i][0] > curX && coordinates[i][1] > curY {
			cnt++
			curX, curY = coordinates[i][0], coordinates[i][1]
		}
	}
	return cnt
	//newCoordinates := make([][2]int, l)
	//for i := 0; i < l; i++ {
	//	if (x < coordinates[i][0] && y < coordinates[i][1]) || (
	//		x > coordinates[i][0] && y > coordinates[i][1]) {
	//		newCoordinates[index][0], newCoordinates[index][1] = coordinates[i][0], coordinates[i][1]
	//		index++
	//	}
	//}
	////return index + 1
	//newCoordinates=newCoordinates[:index]
	//sort.Slice(newCoordinates, func(i, j int) bool {
	//	if newCoordinates[i][0] == newCoordinates[j][0] {
	//		return newCoordinates[i][1] < newCoordinates[j][1]
	//	}
	//	return newCoordinates[i][0] < newCoordinates[j][0]
	//})
	//cnt:=0
	//for _, coordinate := range newCoordinates{
	//
	//}
}
