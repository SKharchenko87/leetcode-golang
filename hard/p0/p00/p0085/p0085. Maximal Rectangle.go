package p0085

/*

 1  0  1  0  0
 2  0  2  1  1
 3  1  3  2  2
 4  0  0  3  0

*/

func maximalRectangle(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])
	matrix[0] = append(matrix[0], 0)
	for j := 0; j < m; j++ {
		matrix[0][j] -= '0'
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j] -= '0'
			matrix[i][j] *= (matrix[i-1][j] + 1)
		}
		matrix[i] = append(matrix[i], 0)
	}
	maxArea := 0
	h := make([]byte, 0, m+1)
	for i := 0; i < n; i++ {
		h = append(h, 0)
		for j := 1; j < m+1; j++ {
			for index := len(h) - 1; index > -1 && matrix[i][j] < matrix[i][h[index]]; index-- {
				lastH := int(matrix[i][h[index]])
				h = h[:index]
				right := j
				left := -1
				if index > 0 {
					left = int(h[index-1])
				}
				maxArea = max(maxArea, lastH*(right-left-1))
			}
			h = append(h, byte(j))
		}
		h = h[:0]
	}
	return maxArea
}

func maximalRectangle0(matrix [][]byte) int {
	heights := make([]int, len(matrix[0])+1)
	heights[len(heights)-1] = -1
	mx := 0
	for _, row := range matrix {
		for i := range row {
			if row[i] == '1' {
				heights[i]++
			} else {
				heights[i] = 0
			}
		}

		stack := []int{}
		for i, currentHeight := range heights {
			for len(stack) > 0 && heights[stack[len(stack)-1]] > currentHeight {
				prev := heights[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]

				width := i
				if len(stack) > 0 {
					width = i - stack[len(stack)-1] - 1
				}
				mx = max(mx, prev*width)
			}
			stack = append(stack, i)
		}
	}
	return mx
}
