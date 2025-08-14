package p0241

import "strconv"

func diffWaysToCompute(expression string) []int {

	var result []int

	for i := 0; i < len(expression); i++ {
		if expression[i] < '0' || '9' < expression[i] {

			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])

			for _, l := range left {
				for _, r := range right {
					switch expression[i] {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}

	if len(result) == 0 {
		num, _ := strconv.Atoi(expression)
		result = append(result, num)
	}

	return result
}
