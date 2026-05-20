package p1559

// ToDo
func containsCycle(grid [][]byte) bool {
	//m, n := len(grid), len(grid[0])
	//ref := make([][]int, 26)
	//for i := 0; i < 26; i++ {
	//	ref[i] = make([]int, 0, n/2)
	//}
	//if m < 2 || n < 2 {
	//	return false
	//}
	//
	//_ = m
	//
	//prevGroup := make([]int, n)
	//ref[grid[0][0]-'a'] = append(ref[grid[0][0]-'a'], 1)
	//prevGroup[0] = 1
	//for j := 1; j < n; j++ {
	//	if grid[0][j] == grid[0][j-1] {
	//		prevGroup[j] = prevGroup[j-1]
	//	} else {
	//		ch:=
	//		prevGroup[j] = len(ref[]) + 1
	//		ref[grid[0][j]-'a'] = append(ref[grid[0][j]-'a'], prevGroup[j])
	//	}
	//
	//}
	//for i := 1; i < m; i++ {
	//	if grid[i][0] != grid[i-1][0] {
	//		prevGroup[0] = len(ref[grid[i][0]-'a']) + 1
	//		ref[grid[i][0]-'a'] = append(ref[grid[i][0]-'a'], prevGroup[0])
	//	}
	//	for j := 1; j < n; j++ {
	//
	//			up, left := grid[i][j] == grid[i-1][j], grid[i][j] == grid[i][j-1]
	//			if up && left {
	//				prevGroup[j] = min(prevGroup[j], prevGroup[j-1])
	//			} else if !up && left {
	//				prevGroup[j] = prevGroup[j-1]
	//			} else if up && !left {
	//			} else {
	//				prevGroup[j] = len(ref[grid[i][j]-'a']) + 1
	//			}
	//
	//	}
	//}
	return false
}
