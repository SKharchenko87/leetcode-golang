package p0947

import "slices"

func removeStones(stones [][]int) int {
	res := 0
	for len(stones) > 0 {
		candidateRow, candidateColumn := [10_001]bool{}, [10_001]bool{}
		candidateRow[stones[0][0]], candidateColumn[stones[0][1]] = true, true
		stones = stones[1:]
		tmp := res - 1
		for tmp != res {
			tmp = res
			for j := 0; j < len(stones); j++ {
				if candidateRow[stones[j][0]] || candidateColumn[stones[j][1]] {
					if candidateRow[stones[j][0]] {
						candidateColumn[stones[j][1]] = true
					}
					if candidateColumn[stones[j][1]] {
						candidateRow[stones[j][0]] = true
					}
					res++
					stones = slices.Delete(stones, j, j+1)
					j--
				}
			}
		}
	}
	return res
}

func removeStones2(stones [][]int) int {
	res := 0
	for i := 0; i < len(stones); i++ {
		candidateRow, candidateColumn := [10_001]bool{}, [10_001]bool{}
		candidateRow[stones[i][0]] = true
		candidateColumn[stones[i][1]] = true
		stones = stones[1:]
		tmp := res - 1
		for tmp != res {
			tmp = res
			for j := 0; j < len(stones); j++ {
				if candidateRow[stones[j][0]] || candidateColumn[stones[j][1]] {
					if candidateRow[stones[j][0]] {
						candidateColumn[stones[j][1]] = true
					}
					if candidateColumn[stones[j][1]] {
						candidateRow[stones[j][0]] = true
					}
					res++
					stones = append(stones[:j], stones[j+1:]...)
					j--
				}
			}
		}
		i--
	}
	return res
}

func removeStones1(stones [][]int) int {
	res := 0
	for i := 0; i < len(stones); i++ {
		candidateRow, candidateColumn := map[int]bool{stones[i][0]: true}, map[int]bool{stones[i][1]: true}
		stones = stones[1:]
		tmp := res - 1
		for tmp != res {
			tmp = res
			for j := 0; j < len(stones); j++ {
				if candidateRow[stones[j][0]] || candidateColumn[stones[j][1]] {
					if candidateRow[stones[j][0]] {
						candidateColumn[stones[j][1]] = true
					}
					if candidateColumn[stones[j][1]] {
						candidateRow[stones[j][0]] = true
					}
					res++
					stones = append(stones[:j], stones[j+1:]...)
					j--
				}
			}
		}
		i--
	}
	return res
}

func removeStones0(stones [][]int) int {
	res := 0
	for i := 0; i < len(stones); i++ {
		candidateRow, candidateColumn := map[int]bool{stones[i][0]: true}, map[int]bool{stones[i][1]: true}
		stones = stones[1:]
		for j := 0; j < len(stones); j++ {
			if _, ok := candidateRow[stones[j][0]]; ok {
				candidateColumn[stones[j][1]] = true
				res++
				stones = append(stones[:j], stones[j+1:]...)
				j--
			} else if _, ok := candidateColumn[stones[j][1]]; ok {
				candidateRow[stones[j][0]] = true
				res++
				stones = append(stones[:j], stones[j+1:]...)
				j--
			}
		}
	}
	return res
}
