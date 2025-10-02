package p3021

func flowerGame(n int, m int) int64 {
	n64 := int64(n)
	m64 := int64(m)
	return n64/2*((m64+1)/2) + m64/2*((n64+1)/2)
}
