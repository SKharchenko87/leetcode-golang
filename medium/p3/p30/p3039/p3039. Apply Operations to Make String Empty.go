package p3039

// ToDo benchmark
func lastNonEmptyString(s string) string {
	letterFrequency := [26]int{}
	maxLetterFrequency := 0
	for _, c := range s {
		index := c - 'a'
		letterFrequency[index]++
		if letterFrequency[index] > maxLetterFrequency {
			maxLetterFrequency = letterFrequency[index]
		}
	}
	res := make([]byte, 26)
	j := 25
	for i := len(s) - 1; i >= 0; i-- {
		if letterFrequency[s[i]-'a'] == maxLetterFrequency {
			res[j] = s[i]
			letterFrequency[s[i]-'a'] = 0
			j--
		}
	}
	return string(res[j+1:])
}

func lastNonEmptyString1(s string) string {
	letterFrequency := [26]int{}
	maxLetterFrequency := 0
	for _, c := range s {
		index := c - 'a'
		letterFrequency[index]++
		if letterFrequency[index] > maxLetterFrequency {
			maxLetterFrequency = letterFrequency[index]
		}
	}
	var res []byte
	for i := len(s) - 1; i >= 0; i-- {
		if letterFrequency[s[i]-'a'] == maxLetterFrequency {
			res = append([]byte{s[i]}, res...)
			letterFrequency[s[i]-'a'] = 0
		}
	}
	return string(res[:])
}
