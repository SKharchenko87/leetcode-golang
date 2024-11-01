package p2463

import (
	"sort"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	// Sort robot and factory by position
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	// Flatten factory positions according to their capacities
	factoryPositions := []int{}
	for _, factory := range factory {
		for i := 0; i < factory[1]; i++ {
			factoryPositions = append(factoryPositions, factory[0])
		}
	}

	robotCount := len(robot)
	factoryCount := len(factoryPositions)
	nextDist := make([]int, factoryCount+1)
	current := make([]int, factoryCount+1)

	// Fill DP table using two rows for optimization
	for i := robotCount - 1; i >= 0; i-- {
		// No factory left case
		if i != robotCount-1 {
			nextDist[factoryCount] = int(1e12)
		}

		// Initialize current row
		current[factoryCount] = int(1e12)

		for j := factoryCount - 1; j >= 0; j-- {
			// Assign current robot to current factory
			assign := abs(robot[i]-factoryPositions[j]) + nextDist[j+1]

			// Skip current factory for this robot
			skip := current[j+1]
			// Take the minimum option
			current[j] = min(assign, skip)
		}

		// Move to next robot
		copy(nextDist, current)
	}

	// Return minimum distance starting from the first robot
	return int64(current[0])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
