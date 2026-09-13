package p0835

import "math/bits"

func largestOverlap(img1 [][]int, img2 [][]int) int {
	cnt := 0
	n := len(img1)
	lines1, lines2 := getLines(img1), getLines(img2)
	for dj := -n + 1; dj < n; dj++ {
		for di := -n + 1; di < n; di++ {
			c := 0
			for i := 0; i < n; i++ {

				if index2 := i + dj; 0 <= index2 && index2 < n {
					var line1Shifted uint
					if di > 0 {
						line1Shifted = lines1[i] >> di
					} else {
						line1Shifted = lines1[i] << (-di)
					}
					c += bits.OnesCount(line1Shifted & lines2[i+dj])
				}
			}
			cnt = max(cnt, c)
		}
	}
	return cnt
}

func getLines(img [][]int) []uint {
	n := len(img)
	lines := make([]uint, n)
	for i := 0; i < len(img); i++ {
		for j := 0; j < n; j++ {
			lines[i] = lines[i]<<1 | uint(img[i][j])
		}
	}
	return lines
}
