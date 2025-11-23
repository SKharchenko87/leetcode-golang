package p3541

func maxFreqSum(s string) int {
	maxVowel := 0
	maxConsonant := 0
	freqMap := make(map[int32]int)
	for _, c := range s {
		freqMap[c]++
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			if maxVowel < freqMap[c] {
				maxVowel = freqMap[c]
			}
		} else {
			if maxConsonant < freqMap[c] {
				maxConsonant = freqMap[c]
			}
		}
	}
	return maxVowel + maxConsonant
}
