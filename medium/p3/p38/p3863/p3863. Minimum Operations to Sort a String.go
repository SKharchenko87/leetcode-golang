package p3863

func minOperations(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	} else if n == 2 {
		if s[0] <= s[1] {
			return 0
		}
		return -1
	}

	last, first := s[0], s[0]
	cntLast, cntFirst := 1, 1
	isSorted := true
	for i := 1; i < n; i++ {
		isSorted = isSorted && s[i-1] <= s[i]

		if first == s[i] {
			cntFirst++
		} else if first > s[i] {
			first = s[i]
			cntFirst = 1
		}

		if last == s[i] {
			cntLast++
		} else if last < s[i] {
			last = s[i]
			cntLast = 1
		}
	}

	if isSorted {
		return 0
	}

	if s[0] == last && s[n-1] == first {
		if cntFirst > 1 || cntLast > 1 {
			return 2
		}
		return 3
	}

	if s[0] == first || s[n-1] == last {
		return 1
	}

	return 2
}
