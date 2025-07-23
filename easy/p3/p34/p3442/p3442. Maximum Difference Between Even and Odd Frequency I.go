package p3442

func maxDifference(s string) int {
	alphabet := [26]int{}
	for _, c := range s {
		alphabet[c-'a']++
	}
	oddMaxFrequency := 0
	evenMinFrequency := 100

	for _, cnt := range alphabet {
		// even
		if cnt%2 == 0 && cnt > 0 && evenMinFrequency > cnt {
			evenMinFrequency = cnt
		} else if cnt%2 == 1 && cnt > 0 && oddMaxFrequency < cnt {
			oddMaxFrequency = cnt
		}
	}

	return oddMaxFrequency - evenMinFrequency
}
