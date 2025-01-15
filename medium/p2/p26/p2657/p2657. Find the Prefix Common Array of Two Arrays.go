package p2657

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	freq, res := make([]int, n), make([]int, n)
	cnt := 0
	for i := 0; i < n; i++ {
		freq[A[i]-1]++
		if freq[A[i]-1] > 1 {
			cnt++
		}
		freq[B[i]-1]++
		if freq[B[i]-1] > 1 {
			cnt++
		}
		res[i] = cnt
	}
	return res
}

func findThePrefixCommonArray1(A []int, B []int) []int {
	n := len(A)
	a, b, res := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[A[i]-1] = i
		b[B[i]-1] = i
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if A[i] == B[i] {
			cnt++
		} else {
			if b[A[i]-1] < a[A[i]-1] {
				cnt++
			}
			if a[B[i]-1] < b[B[i]-1] {
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}

func findThePrefixCommonArray0(A []int, B []int) []int {
	n := len(A)
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[A[i]-1] = i
	}
	for i := 0; i < n; i++ {
		b[B[i]-1] = i
	}
	res := make([]int, n)
	cnt := 0
	for i := 0; i < n; i++ {
		if A[i] == B[i] {
			cnt++
		} else {
			if b[A[i]-1] < a[A[i]-1] {
				cnt++
			}
			if a[B[i]-1] < b[B[i]-1] {
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}
