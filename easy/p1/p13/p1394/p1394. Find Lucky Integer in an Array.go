package p1394

func findLucky(arr []int) int {
	freq := make([]int, 501)
	for i := 0; i < len(arr); i++ {
		freq[arr[i]]++
	}
	for i := 500; i > 0; i-- {
		if freq[i] == i {
			return i
		}
	}
	return -1
}
