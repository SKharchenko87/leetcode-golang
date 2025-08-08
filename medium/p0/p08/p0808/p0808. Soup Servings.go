package p0808

import "math"

var result = [192][192]float64{}

func init() {
	for i := 0; i < 192; i++ {
		for j := 0; j < 192; j++ {
			result[i][j] = -1.0
		}
	}

	var rec func(a, b int) (res float64)
	rec = func(a, b int) (res float64) {

		defer func() {
			if a >= 0 && b >= 0 {
				result[a][b] = math.Round(res*1e5) / 1e5
			}
		}()

		if a <= 0 && b > 0 {
			return 1
		} else if a <= 0 && b <= 0 {
			return 0.5
		} else if a > 0 && b <= 0 {
			return 0
		}

		if v := result[a][b]; v != -1 {
			res = v
			return
		}

		res = 0.25 * (rec(a-4, b) + rec(a-3, b-1) + rec(a-2, b-2) + rec(a-1, b-3))
		return
	}
	for i := 0; i < 192; i++ {
		rec(i, i)
	}

}

func soupServings(n int) float64 {
	if n > 4800 {
		return 1.0
	}
	n = (n + 24) / 25
	return result[n][n]
}

func soupServings3(n int) float64 {
	if n > 4800 {
		return 1.0
	}

	n = (n + 24) / 25

	mem := [192][192]float64{}
	for i := 0; i < 192; i++ {
		for j := 0; j < 192; j++ {
			mem[i][j] = -1.0
		}
	}
	var rec func(a, b int) (res float64)
	rec = func(a, b int) (res float64) {

		defer func() {
			if a >= 0 && b >= 0 {
				mem[a][b] = res
			}
		}()

		if a <= 0 && b > 0 {
			return 1
		} else if a <= 0 && b <= 0 {
			return 0.5
		} else if a > 0 && b <= 0 {
			return 0
		}

		if v := mem[a][b]; v != -1 {
			res = v
			return
		}

		res = 0.25 * (rec(a-4, b) + rec(a-3, b-1) + rec(a-2, b-2) + rec(a-1, b-3))
		return
	}
	return math.Round(rec(n, n)*1e5) / 1e5
}

func soupServings2(n int) float64 {
	if n > 4800 {
		return 1.0
	}
	// нормализация в "порции по 25 мл"
	n = (n + 24) / 25

	mem := map[[2]int]float64{}
	var rec func(a, b int) (res float64)
	rec = func(a, b int) (res float64) {
		if val, ok := mem[[2]int{a, b}]; ok {
			return val
		}
		defer func() { mem[[2]int{a, b}] = res }()

		if a <= 0 && b > 0 {
			return 1
		} else if a <= 0 && b <= 0 {
			return 0.5
		} else if a > 0 && b <= 0 {
			return 0
		}

		res = 0.25 * (rec(a-4, b) + rec(a-3, b-1) + rec(a-2, b-2) + rec(a-1, b-3))
		return
	}
	return math.Round(rec(n, n)*1e5) / 1e5
}

func soupServings1(n int) float64 {
	if n > 4800 {
		return 1.0
	}
	mem := map[[2]int]float64{}
	var rec func(a, b int) float64
	rec = func(a, b int) (res float64) {
		if _, ok := mem[[2]int{a, b}]; ok {
			return mem[[2]int{a, b}]
		}
		defer func() {
			mem[[2]int{a, b}] = res
		}()
		if a <= 0 && b > 0 {
			res = 1
			return
		} else if a <= 0 && b <= 0 {
			res = 0.5
			return
		} else if a > 0 && b <= 0 {
			res = 0
			return
		}
		res = 0.25 * (rec(a-100, b-0) + rec(a-75, b-25) + rec(a-50, b-50) + rec(a-25, b-75))
		return
	}
	return math.Round(rec(n, n)*1e5) / 1e5
}

// TLE
func soupServings0(n int) float64 {
	if n > 4800 {
		return 1.0
	}
	var rec func(a, b int) float64
	rec = func(a, b int) float64 {
		if a <= 0 && b > 0 {
			return 1
		} else if a <= 0 && b <= 0 {
			return 0.5
		} else if a > 0 && b <= 0 {
			return 0
		}
		return 0.25 * (rec(a-100, b-0) + rec(a-75, b-25) + rec(a-50, b-50) + rec(a-25, b-75))
	}
	return math.Round(rec(n, n)*1e5) / 1e5
}
