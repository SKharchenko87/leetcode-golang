package p3272

var pow10 = []int64{1, 10, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 10_000_000_000}
var factorial = []int64{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}
var answers = map[byte]map[byte]int64{
	1:  {1: 9, 2: 4, 3: 3, 4: 2, 5: 1, 6: 1, 7: 1, 8: 1, 9: 1},
	2:  {1: 9, 2: 4, 3: 3, 4: 2, 5: 1, 6: 1, 7: 1, 8: 1, 9: 1},
	3:  {1: 243, 2: 108, 3: 69, 4: 54, 5: 27, 6: 30, 7: 33, 8: 27, 9: 23},
	4:  {1: 252, 2: 172, 3: 84, 4: 98, 5: 52, 6: 58, 7: 76, 8: 52, 9: 28},
	5:  {1: 10935, 2: 7400, 3: 3573, 4: 4208, 5: 2231, 6: 2468, 7: 2665, 8: 2231, 9: 1191},
	6:  {1: 10944, 2: 9064, 3: 3744, 4: 6992, 5: 3256, 6: 3109, 7: 3044, 8: 5221, 9: 1248},
	7:  {1: 617463, 2: 509248, 3: 206217, 4: 393948, 5: 182335, 6: 170176, 7: 377610, 8: 292692, 9: 68739},
	8:  {1: 617472, 2: 563392, 3: 207840, 4: 494818, 5: 237112, 6: 188945, 7: 506388, 8: 460048, 9: 69280},
	9:  {1: 41457015, 2: 37728000, 3: 13726509, 4: 33175696, 5: 15814071, 6: 12476696, 7: 36789447, 8: 30771543, 9: 4623119},
	10: {1: 41457024, 2: 39718144, 3: 13831104, 4: 37326452, 5: 19284856, 6: 13249798, 7: 40242031, 8: 35755906, 9: 4610368},
}

func countGoodIntegers(n int, k int) int64 {
	return answers[byte(n)][byte(k)]
}

func countGoodIntegers0(n int, k int) int64 {
	m := make(map[int64]bool, 2000)
	center := n/2 + n%2 - 1
	var dfs func(x int64, i int) int64
	dfs = func(x int64, i int) int64 {
		res := int64(0)
		if center >= i {
			start := int64(0)
			if i == 0 {
				start = 1
			}
			for j := start; j <= 9; j++ {
				y := x + pow10[i]*j
				if !(n%2 == 1 && i == n/2) {
					y += pow10[n-i-1] * j
				}
				if y%int64(k) == 0 {
					z, count := getDigits(y)
					if _, ok := m[z]; !ok {
						m[z] = true
						res += count
					}
				}
				res += dfs(y, i+1)
			}
		}
		return res
	}
	res := dfs(0, 0)
	return res
}

func getDigits(num int64) (int64, int64) {
	var freq [10]int
	length := 0
	for num > 0 {
		freq[num%10]++
		num /= 10
		length++
	}

	var x int64
	var variants int64
	variants = factorial[length-1] * int64(length-freq[0])
	for i := 9; i >= 0; i-- {
		b := freq[i]
		if b > 0 {
			for j := 0; j < b; j++ {
				x = x*10 + int64(i)
			}
			variants /= factorial[b]
		}
	}

	return x, variants
}
