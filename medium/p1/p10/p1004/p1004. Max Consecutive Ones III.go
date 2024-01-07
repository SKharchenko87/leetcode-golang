package p1004

import (
	"fmt"
	"math"
)

/*ToDo оптимизировать за один проход с созданием arr массива [k+2]int в котором будет start(arr[0]) и end(arr[len(arr)])
0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1   k=3
sc c        ce
   sc       c  c        e
      s     c  c           c     e
               sc          c        ce
                  s        c        c  ce
                              s     c  c  c           e
*/

func longestOnes(nums []int, k int) int {
	l := len(nums)
	maxLength := math.MinInt
	var length int
	for i := 0; i < l; i++ {
		firstLastOne := -1
		cntZero := 0
		length = 0
		for j := i; j < l; j++ {
			if nums[j] == 0 {
				if firstLastOne == -1 {
					firstLastOne = j - 1
				}
				if cntZero < k {
					cntZero++
					length++
				} else {
					if maxLength < length {
						maxLength = length
						fmt.Println(i, j)
					}
					break
				}
			} else {
				length++
			}
		}
		i = max(firstLastOne-2, i)
		if maxLength < length {
			maxLength = length
			fmt.Println(i, l-1)
		}
	}
	return maxLength
}
