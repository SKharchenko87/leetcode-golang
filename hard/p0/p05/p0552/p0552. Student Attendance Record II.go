package p0552

const MOD = 1000_000_007

type ccc struct {
	g   int
	r   int
	y   int
	ry  int
	ryy int
	yy  int
}

func checkRecord(n int) int {
	dp := make([]ccc, n+1)
	dp[1] = ccc{g: 1, r: 1, y: 1, ry: 0, ryy: 0, yy: 0}

	for i := 2; i <= n; i++ {
		dp[i].g = dp[i-1].g + dp[i-1].y + dp[i-1].yy
		dp[i].r = dp[i-1].g*2 + dp[i-1].y + dp[i-1].yy
		dp[i].y = dp[i-1].g
		dp[i].ry = dp[i-1].r
		dp[i].ryy = dp[i-1].ry
		dp[i].yy = dp[i-1].y
		println(dp[n].g, dp[n].r, dp[n].y, dp[n].ry, dp[n].ryy, dp[n].yy)
	}
	return dp[n].g + dp[n].r + dp[n].y + dp[n].ry + dp[n].ryy + dp[n].yy
}

func checkRecord2(n int) int {
	if n == 0 {
		return 0
	}

	// dp[a][l] represents the number of valid sequences with 'a' absents and 'l' consecutive lates.
	dp := [2][3]int{}

	dp[0][0] = 1 // Initial state: "P"
	dp[0][1] = 1 // Initial state: "L"
	dp[0][2] = 0 // No initial state: "LL"

	dp[1][0] = 1 // Initial state: "A"
	dp[1][1] = 0 // No initial state: "AL"
	dp[1][2] = 0 // No initial state: "ALL"

	// Iterate over the length of the record
	for i := 1; i < n; i++ {
		newDp := [2][3]int{}

		// Update states without absences
		newDp[0][0] = (dp[0][0] + dp[0][1] + dp[0][2]) % MOD // Ending with "P"
		newDp[0][1] = dp[0][0]                               // Ending with "L"
		newDp[0][2] = dp[0][1]                               // Ending with "LL"

		// Update states with one absence
		newDp[1][0] = (dp[1][0] + dp[1][1] + dp[1][2] + dp[0][0] + dp[0][1] + dp[0][2]) % MOD // Ending with "P"
		newDp[1][1] = dp[1][0]                                                                // Ending with "L"
		newDp[1][2] = dp[1][1]                                                                // Ending with "LL"

		dp = newDp // Move to the next state
	}

	// Sum all valid states
	return (dp[0][0] + dp[0][1] + dp[0][2] + dp[1][0] + dp[1][1] + dp[1][2]) % MOD
}

type variant struct {
	green        int
	greenRed     int
	yellow       int
	yellowYellow int
	yellowRed    int
	red          int
	z            int
	total        int
}

func checkRecord1(n int) int {
	dp := make([]variant, n+1)
	dp[1] = variant{1, 0, 1, 0, 0, 1, 0, 3}
	if n > 1 {
		dp[2] = variant{3, 1, 3, 1, 1, 2, 1, 8}
	}

	for i := 3; i <= n; i++ {
		x := dp[i-1]
		println(dp[i-1].green, dp[i-1].greenRed, dp[i-1].yellow, dp[i-1].yellowYellow, dp[i-1].yellowRed, x.z, dp[i-1].red, x.total)

		dp[i].green = x.green + x.yellow + x.red
		dp[i].greenRed = x.greenRed + x.red + x.yellowRed

		dp[i].yellow = x.green + x.yellow - x.yellowYellow + x.red
		dp[i].yellowYellow = x.yellow - x.yellowYellow
		dp[i].yellowRed = x.yellowRed + x.greenRed + x.red - dp[i-2].z
		dp[i].z = x.red + dp[i-2].z
		dp[i].red = x.green - x.greenRed + x.yellow - x.yellowRed
		dp[i].total = (dp[i].green + dp[i].yellow + dp[i].red) % MOD
	}
	println(dp[n].green, dp[n].greenRed, dp[n].yellow, dp[n].yellowYellow, dp[n].yellowRed, dp[n].z, dp[n].red, dp[n].total)
	return dp[n].total
}

func checkRecord0(n int) int {
	m := map[[3]int]int{}
	var recursive func(i, a, l int) int
	recursive = func(i, a, l int) int {
		if v, ok := m[[3]int{i, a, l}]; ok {
			return v
		}
		if a > 1 || l > 2 {
			return 0
		}
		if i == n {
			return 1
		}
		i++
		p := (recursive(i, a+1, 0) + recursive(i, a, l+1) + recursive(i, a, 0)) % MOD
		m[[3]int{i - 1, a, l}] = p
		return p
	}
	return recursive(0, 0, 0)
}
