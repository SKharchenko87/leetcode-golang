package p1074

// ToDo Разобраться
func numSubmatrixSumTarget(matrix [][]int, target int) int {

	// height and width of matrix
	h, w := len(matrix), len(matrix[0])

	// update prefix sum on each row
	for y := 0; y < h; y++ {
		for x := 1; x < w; x++ {
			matrix[y][x] = matrix[y][x] + matrix[y][x-1]
		}
	}

	// number of submatrices that sum to target
	counter := 0

	// sliding windows on x-axis, in range[left, right]
	for left := 0; left < w; left++ {
		for right := left; right < w; right++ {

			// accumulation for area so far
			accumulation := map[int]int{0: 1}

			// area of current submatrices, bounded by [left, right] with height y
			area := 0

			// scan each possible height on y-axis
			for y := 0; y < h; y++ {

				if left > 0 {
					area += matrix[y][right] - matrix[y][left-1]
				} else {
					area += matrix[y][right]
				}

				// if ( area - target ) exists, then target must exist in submatrices
				counter += accumulation[(area - target)]

				// update dictionary with current accumulation area
				accumulation[area] += 1

			}

		} //end of right loop

	} //end of left loop

	return counter

}
