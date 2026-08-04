package p0486

/*ToDo*/
func predictTheWinner(nums []int) bool {
	return false
}

/* Не верно из-за You may assume that both players are playing optimally.
func predictTheWinner(nums []int) bool {
	n := len(nums)
	ones := (n + 1) / 2
	v := (1 << ones) - 1
	limit := 1 << n
	for v < limit {
		c := v & -v
		r := v + c
		v = (((r ^ v) >> 2) / c) | r
		x := v
		player1, player2 := 0, 0
		for i := 0; i < n; i++ {
			if x&1 == 1 {
				player1 += nums[i]
			} else {
				player2 += nums[i]
			}
			x >>= 1
		}
		if player1 >= player2 {
			return true
		}
	}
	return false
}
*/
