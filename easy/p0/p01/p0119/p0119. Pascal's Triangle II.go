package p0119

const MAX_ROW_INDEX = 33

var triangle = make([][]int, MAX_ROW_INDEX+1)

func init() {
	triangle[0] = []int{1}
	triangle[1] = []int{1, 1}
	for i := 2; i <= MAX_ROW_INDEX; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j] + triangle[i-1][j-1]
		}
	}
}

func getRow(rowIndex int) []int {
	return triangle[rowIndex]
}

func getRow0(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex+1; i++ {
		for j := i - 1; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}

func getRow1(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}

	return row
}
