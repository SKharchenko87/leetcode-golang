package p2467

import "math"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)

	// Create adjacency list representation of the tree
	tree := make([][]int, n)
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	// Initialize distanceFromBob array with a large value
	distanceFromBob := make([]int, n)
	for i := range distanceFromBob {
		distanceFromBob[i] = n
	}

	return findPaths(0, -1, 0, bob, amount, tree, distanceFromBob)
}

// findPaths performs DFS to calculate the maximum profit
func findPaths(sourceNode, parentNode, time, bob int, amount []int, tree [][]int, distanceFromBob []int) int {
	maxIncome := 0
	maxChild := math.MinInt32

	// If Bob is at the current node, set distance to 0, otherwise set to max
	if sourceNode == bob {
		distanceFromBob[sourceNode] = 0
	} else {
		distanceFromBob[sourceNode] = len(amount)
	}

	// Traverse adjacent nodes
	for _, adjacentNode := range tree[sourceNode] {
		if adjacentNode != parentNode {
			maxChild = max(maxChild, findPaths(adjacentNode, sourceNode, time+1, bob, amount, tree, distanceFromBob))
			distanceFromBob[sourceNode] = min(distanceFromBob[sourceNode], distanceFromBob[adjacentNode]+1)
		}
	}

	// Alice reaches the node first
	if distanceFromBob[sourceNode] > time {
		maxIncome += amount[sourceNode]
	} else if distanceFromBob[sourceNode] == time { // Alice and Bob reach simultaneously
		maxIncome += amount[sourceNode] / 2
	}

	// If leaf node, return the max income
	if maxChild == math.MinInt32 {
		return maxIncome
	} else {
		return maxIncome + maxChild
	}
}
