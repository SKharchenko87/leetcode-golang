package p2579

func coloredCells(n int) int64 {
	return 1 + 2*int64((n-1)*n)
}

func coloredCells0(n int) int64 {
	n--
	return 1 + 4*int64((n+1)*n/2)
}
