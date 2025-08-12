package p2787

const Mod = 1e9 + 7

var pow = [5][301]int{}

var dp = [5][301][301]int{}

func init() {
	for j := 0; j < len(pow[0]); j++ {
		pow[0][j] = j + 1
	}
	for i := 1; i < len(pow); i++ {
		for j := 0; j < len(pow[0]); j++ {
			pow[i][j] = pow[i-1][j] * pow[0][j]
		}
	}
	dp[0][0][0] = 1
	dp[1][0][0] = 1
	dp[2][0][0] = 1
	dp[3][0][0] = 1
	dp[4][0][0] = 1

	for p := 0; p < 5; p++ {
		for i := 1; i <= 300; i++ {
			val := pow[p][i-1]
			for j := 0; j <= 300; j++ {
				dp[p][i][j] = dp[p][i-1][j]
				if j >= val {
					dp[p][i][j] = (dp[p][i][j] + dp[p][i-1][j-val]) % Mod
				}
			}
		}
	}
}

func numberOfWays(n int, x int) int {
	return dp[x-1][n][n]
}

func numberOfWays1(n int, x int) int {
	memorize := [301][301]int{}
	var backtracking func(currSum int, prev int) int
	backtracking = func(currSum int, prev int) int {

		v := memorize[currSum][prev]
		if v > 0 {
			return v - 1
		}

		if currSum == n {
			memorize[currSum][prev] = 2
			return 1
		} else if currSum > n {
			memorize[currSum][prev] = 1
			return 0
		} else {
			res := 0
			c := pow[x-1][prev]
			for i := prev + 1; c <= n-currSum; i++ {
				res += backtracking(currSum+c, i)
				res %= Mod
				c = pow[x-1][i]
			}
			memorize[currSum][prev] = res + 1
			return res
		}
	}
	return backtracking(0, 0)
}

func numberOfWays0(n int, x int) int {
	memorize := [301][301]int{}
	var backtracking func(currSum int, prev int) int
	backtracking = func(currSum int, prev int) int {

		v := memorize[currSum][prev]
		if v > 0 {
			return v - 1
		}

		if currSum == n {
			memorize[currSum][prev] = 2
			return 1
		} else if currSum > n {
			memorize[currSum][prev] = 1
			return 0
		} else {
			res := 0
			c := pow[x-1][prev]
			for i := prev + 1; c <= n-currSum; i++ {
				res += backtracking(currSum+c, i)
				res %= Mod
				c = pow[x-1][i]
			}
			memorize[currSum][prev] = res + 1
			return res
		}
	}
	return backtracking(0, 0)
}
