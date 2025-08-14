package p0874

import (
	"slices"
	"sort"
)

const (
	TURN_LEFT  = -2
	TURN_RIGHT = -1
)

type position struct {
	y, x int
}

func robotSim(commands []int, obstacles [][]int) int {
	sort.Slice(obstacles, func(i, j int) bool {
		if obstacles[i][0] == obstacles[j][0] {
			return obstacles[i][1] < obstacles[j][1]
		}
		return obstacles[i][0] < obstacles[j][0]
	})
	maxEuclideanDistance := 0
	currentPosition := position{0, 0}
	directions := position{0, 1}
	for _, command := range commands {
		if command == TURN_LEFT {
			directions.y, directions.x = -1*directions.x, directions.y
		} else if command == TURN_RIGHT {
			directions.y, directions.x = directions.x, -1*directions.y
		} else {
			for i := 0; i < command; i++ {
				y := currentPosition.y + directions.y
				x := currentPosition.x + directions.x
				_, flg := slices.BinarySearchFunc(obstacles, []int{y, x}, func(a []int, b []int) int {
					if a[0] == b[0] {
						return a[1] - b[1]
					}
					return a[0] - b[0]
				})
				if flg {
					break
				}
				currentPosition.y, currentPosition.x = y, x
			}
			maxEuclideanDistance = max(maxEuclideanDistance, currentPosition.y*currentPosition.y+currentPosition.x*currentPosition.x)
		}
	}
	return maxEuclideanDistance
}

func robotSim1(commands []int, obstacles [][]int) int {
	obstaclesMap := make(map[position]bool, len(obstacles))
	for _, obstacle := range obstacles {
		y, x := obstacle[0], obstacle[1]
		obstaclesMap[position{y, x}] = true
	}
	maxEuclideanDistance := 0
	currentPosition := position{0, 0}
	directions := position{0, 1}
	for _, command := range commands {
		if command == TURN_LEFT {
			directions.y, directions.x = -1*directions.x, directions.y
		} else if command == TURN_RIGHT {
			directions.y, directions.x = directions.x, -1*directions.y
		} else {
			for i := 0; i < command; i++ {
				y := currentPosition.y + directions.y
				x := currentPosition.x + directions.x
				if obstaclesMap[position{y, x}] {
					break
				}
				currentPosition.y, currentPosition.x = y, x
				maxEuclideanDistance = max(maxEuclideanDistance, y*y+x*x)
			}
		}
	}
	return maxEuclideanDistance
}

func robotSim0(commands []int, obstacles [][]int) int {
	obstaclesX := make(map[int]map[int]bool, len(obstacles))
	obstaclesY := make(map[int]map[int]bool, len(obstacles))
	for _, obstacle := range obstacles {
		if _, ok := obstaclesX[obstacle[0]]; !ok {
			obstaclesX[obstacle[0]] = make(map[int]bool)
		}
		obstaclesX[obstacle[0]][obstacle[1]] = true
		if _, ok := obstaclesY[obstacle[1]]; !ok {
			obstaclesY[obstacle[1]] = make(map[int]bool)
		}
		obstaclesY[obstacle[1]][obstacle[0]] = true
	}
	position := []int{0, 0}
	maxEuclideanDistance := 0
	directions := []int{0, 1}
	for _, command := range commands {
		if command < 0 {
			changeDirection(command, directions)
		} else {
			for i := 0; i < command; i++ {
				position[0] += directions[0]
				position[1] += directions[1]
				if _, ok := obstaclesX[position[0]][position[1]]; ok {
					position[0] -= directions[0]
					position[1] -= directions[1]
					break
				}
				maxEuclideanDistance = max(maxEuclideanDistance, position[0]*position[0]+position[1]*position[1])
			}
		}
	}
	return maxEuclideanDistance
}

func changeDirection(command int, directions []int) {
	if command == TURN_LEFT {
		directions[0], directions[1] = -1*directions[1], directions[0]
	} else if command == TURN_RIGHT {
		directions[0], directions[1] = directions[1], -1*directions[0]
	}
}
