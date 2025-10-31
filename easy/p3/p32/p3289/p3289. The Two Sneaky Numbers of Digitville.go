package p3289

import "math"

func getSneakyNumbers(nums []int) (res []int) {
	n := len(nums) - 2
	referenceAmount := n * (n - 1) / 2
	referenceAmountSq := n * (n - 1) * (2*n - 1) / 6
	var sum, sumSq int
	for _, num := range nums {
		sum += num
		sumSq += num * num
	}
	sumAB := sum - referenceAmount                  /*A+B*/
	sumABSq := sumSq - referenceAmountSq            /*A^2+B^2*/
	sumABSqMinus := 2*sumABSq - sumAB*sumAB         /*2*A^2+2*B^2-(A+B)*(A+B) = 2*A^2+2*B^2-A^2-2*A*B-*B^2 = A^2 - 2*A*B + B^2 = (A-B)^2 */
	diffAB := int(math.Sqrt(float64(sumABSqMinus))) /*A-B*/
	A := (sumAB + diffAB) / 2
	B := (sumAB - diffAB) / 2
	return []int{A, B}
}

func getSneakyNumbers1(nums []int) (res []int) {
	var h, l int64
	res = make([]int, 0, 2)
	f := func(x int, diff int, register *int64) bool {
		if 1<<x&*register > 0 {
			res = append(res, x+diff)
			if len(res) == 2 {
				return true
			}
		}
		*register |= 1 << x
		return false
	}
	for _, num := range nums {
		if num < 50 {
			if f(num, 0, &l) {
				return
			}
		} else {
			num -= 50
			if f(num, 50, &h) {
				return
			}
		}
	}
	return
}

func getSneakyNumbers0(nums []int) (res []int) {
	var h, l int64
	res = make([]int, 0, 2)
	for _, num := range nums {
		if num < 50 {
			if 1<<num&l > 0 {
				res = append(res, num)
				if len(res) == 2 {
					return
				}
			}
			l |= 1 << num
		} else {
			num -= 50
			if 1<<num&h > 0 {
				res = append(res, num+50)
				if len(res) == 2 {
					return
				}
			}
			h |= 1 << num
		}
	}
	return
}
