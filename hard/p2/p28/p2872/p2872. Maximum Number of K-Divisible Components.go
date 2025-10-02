package p2872

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	if n == 1 {
		return 1
	}

	neighbours := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		neighbours[a] = append(neighbours[a], b)
		neighbours[b] = append(neighbours[b], a)
	}

	cnt := 0

	var dfs func(cur, previous int) int
	dfs = func(cur, previous int) int {
		sum := values[cur]
		for _, neighbour := range neighbours[cur] {
			if neighbour != previous {
				sum += dfs(neighbour, cur)
			}
		}
		if sum%k == 0 {
			cnt++
		}
		return sum
	}

	dfs(0, -1)
	return cnt
}

func maxKDivisibleComponentsX(n int, edges [][]int, values []int, k int) int {
	var traverse func(currentNode, parentNode int, adjList [][]int, componentCount *int) int
	traverse = func(currentNode, parentNode int, adjList [][]int, componentCount *int) int {
		// Step 1: Initialize sum for the current subtree
		sum := 0

		// Step 2: Traverse all neighbors
		for _, neighborNode := range adjList[currentNode] {
			if neighborNode != parentNode {
				// Recursive call to process the subtree rooted at the neighbor
				sum += traverse(neighborNode, currentNode, adjList, componentCount)
				sum %= k // Ensure the sum stays within bounds
			}
		}

		// Step 3: Add the value of the current node to the sum
		sum += values[currentNode]
		sum %= k

		// Step 4: Check if the sum is divisible by k
		if sum == 0 {
			*componentCount++
		}

		// Step 5: Return the computed sum for the current subtree
		return sum
	}

	// Step 1: Create adjacency list from edges
	adjList := make([][]int, n)
	for _, edge := range edges {
		node1 := edge[0]
		node2 := edge[1]
		adjList[node1] = append(adjList[node1], node2)
		adjList[node2] = append(adjList[node2], node1)
	}

	// Step 2: Initialize component count
	componentCount := 0 // Use array to pass by reference

	// Step 3: Start DFS traversal from node 0
	traverse(0, -1, adjList, &componentCount)

	// Step 4: Return the total number of components
	return componentCount
}

type Node struct {
	value      int
	neighbours map[int]struct{}
}

func maxKDivisibleComponents0(n int, edges [][]int, values []int, k int) int {
	if n == 1 {
		return 1
	}

	cnt := 0
	candidate := make([]int, 0, n)
	graph := make(map[int]*Node, n)

	for _, edge := range edges {
		a, b := edge[0], edge[1]

		if _, ok := graph[a]; !ok {
			graph[a] = &Node{value: values[a], neighbours: map[int]struct{}{}}
		}
		graph[a].neighbours[b] = struct{}{}

		if _, ok := graph[b]; !ok {
			graph[b] = &Node{value: values[b], neighbours: map[int]struct{}{}}
		}
		graph[b].neighbours[a] = struct{}{}
	}

	for i, node := range graph {
		if len(node.neighbours) == 1 {
			candidate = append(candidate, i)
		}
	}

	for i := 0; i < len(candidate); i++ {
		index := candidate[i]
		if graph[index].value%k == 0 {
			cnt++
			tmp := graph[index]
			tmp.value = 0
		}
		for key := range graph[index].neighbours {
			tmp := graph[key]
			tmp.value += graph[index].value
			delete(tmp.neighbours, index)
			if len(tmp.neighbours) == 1 {
				candidate = append(candidate, key)
			}

		}
	}
	return cnt
}
